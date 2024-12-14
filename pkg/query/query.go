package query

import (
	"fmt"
	"github.com/sherwin-77/go-tix/pkg/constants"
	"github.com/sherwin-77/go-tix/pkg/response"
	"gorm.io/gorm"
	"math"
	"net/url"
	"strconv"
	"strings"
)

type FilterType string
type SortDirection string

const (
	FilterTypeExact   FilterType = "exact"
	FilterTypePartial FilterType = "partial"
	FilterTypeCustom  FilterType = "custom"

	SortDirectionAscending  SortDirection = "ASC"
	SortDirectionDescending SortDirection = "DESC"
)

type FilterParam struct {
	DisplayName       string
	FieldName         string
	InternalName      string
	DisplayFilterType string
	FilterType        FilterType
	Callback          func(db *gorm.DB, value string) *gorm.DB
}

type SortParam struct {
	DisplayName  string
	FieldName    string
	InternalName string
	Direction    SortDirection
	Callback     func(db *gorm.DB, isDescending bool) *gorm.DB
}

type Builder interface {
	ApplyBuilder(db *gorm.DB, queryParams url.Values, model interface{}) (*gorm.DB, *response.Meta)
}

type builder struct {
	AllowedFilters []FilterParam
	AllowedSorts   []SortParam
	DefaultSort    SortParam
}

func NewBuilder(allowedFilters []FilterParam, allowedSorts []SortParam, defaultSort SortParam) Builder {
	return &builder{
		AllowedFilters: allowedFilters,
		AllowedSorts:   allowedSorts,
		DefaultSort:    defaultSort,
	}
}

func (b *builder) getFilterField(param FilterParam) string {
	if param.InternalName == "" {
		return param.FieldName
	}

	return param.InternalName
}

func (b *builder) getSortField(param SortParam) string {
	if param.InternalName == "" {
		return param.FieldName
	}

	return param.InternalName
}

// extractFilters parses query params and matches them to allowed filters
func (b *builder) extractFilters(queryParams url.Values, allowedFilters []FilterParam) map[string]string {
	filters := make(map[string]string)
	for _, filter := range allowedFilters {
		if value := queryParams.Get(fmt.Sprintf("filter[%s]", filter.FieldName)); value != "" {
			filters[filter.FieldName] = value
		}
	}
	return filters
}

// extractSorting parses query params for sorting (field and direction)
func (b *builder) extractSorting(queryParams url.Values, allowedSorts []SortParam, defaultSort SortParam) (string, SortDirection, func(db *gorm.DB, isDescending bool) *gorm.DB) {
	sortField := queryParams.Get("sort")
	sortDirection := SortDirectionAscending
	var sortCallback func(db *gorm.DB, isDescending bool) *gorm.DB
	if strings.HasPrefix(sortField, "-") {
		sortDirection = SortDirectionDescending
		sortField = strings.TrimPrefix(sortField, "-")
	}

	isValidSort := false
	for _, sort := range allowedSorts {
		if sort.FieldName == sortField {
			sortField = b.getSortField(sort)
			isValidSort = true
			sortCallback = sort.Callback
			break
		}
	}

	if !isValidSort {
		sortField = b.getSortField(defaultSort)
		sortDirection = defaultSort.Direction
	}

	return sortField, sortDirection, sortCallback
}

// extractPagination parses query params for pagination (limit and page)
func (b *builder) extractPagination(queryParams url.Values) (int, int) {
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil || limit < 1 {
		limit = constants.DefaultPerPage
	}
	limit = min(limit, constants.MaxPerPage)

	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	return limit, page
}

func (b *builder) getFilterMeta(queryParams url.Values) []response.FilterMeta {
	var filters []response.FilterMeta
	for _, filter := range b.AllowedFilters {
		value := queryParams.Get(fmt.Sprintf("filter[%s]", filter.FieldName))
		filters = append(filters, response.FilterMeta{
			Name:       filter.FieldName,
			Label:      filter.DisplayName,
			FilterType: filter.DisplayFilterType,
			Value:      value,
		})
	}
	return filters
}

func (b *builder) getSortMeta() []response.SortMeta {
	var sorts []response.SortMeta
	for _, sort := range b.AllowedSorts {
		sorts = append(sorts, response.SortMeta{
			Name:  sort.FieldName,
			Label: sort.DisplayName,
		})
	}
	return sorts
}

func (b *builder) ApplyBuilder(db *gorm.DB, queryParams url.Values, model interface{}) (*gorm.DB, *response.Meta) {
	if b.InitBuilder != nil {
		db = b.InitBuilder(db)
	}

	filters := b.extractFilters(queryParams, b.AllowedFilters)

	for _, filter := range b.AllowedFilters {
		if value, ok := filters[filter.FieldName]; ok {
			switch filter.FilterType {
			case FilterTypeExact:
				field := b.getFilterField(filter)
				db = db.Where(field+" = ?", value)
			case FilterTypePartial:
				field := b.getFilterField(filter)
				db = db.Where(field+" ILIKE ?", "%"+value+"%")
			case FilterTypeCustom:
				if filter.Callback != nil {
					db = filter.Callback(db, value)
				}
			default: // Fallback if filter type not defined
				field := b.getFilterField(filter)
				db = db.Where(field+" = ?", value)
			}
		}
	}

	var count int64
	db.Model(model).Count(&count)

	sortField, sortDirection, sortCallback := b.extractSorting(queryParams, b.AllowedSorts, b.DefaultSort)
	if sortCallback != nil {
		db = sortCallback(db, sortDirection == SortDirectionDescending)
	} else {
		db = db.Order(sortField + " " + string(sortDirection))
	}

	limit, page := b.extractPagination(queryParams)
	db = db.Offset((page - 1) * limit).Limit(limit)

	defaultSort := b.DefaultSort.FieldName
	if b.DefaultSort.Direction == SortDirectionDescending {
		defaultSort = "-" + defaultSort
	}

	selectedSort := sortField
	if sortDirection == SortDirectionDescending {
		selectedSort = "-" + selectedSort
	}

	meta := &response.Meta{
		Page:         page,
		PerPage:      limit,
		LastPage:     int(math.Ceil(float64(count) / float64(limit))),
		Total:        count,
		Filters:      b.getFilterMeta(queryParams),
		Sorts:        b.getSortMeta(),
		SelectedSort: selectedSort,
		DefaultSort:  defaultSort,
	}

	return db, meta
}
