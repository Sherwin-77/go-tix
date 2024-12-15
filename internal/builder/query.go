package builder

import (
	"github.com/sherwin-77/go-tix/pkg/constants"
	"github.com/sherwin-77/go-tix/pkg/query"
	"gorm.io/gorm"
)

func NewUserQueryBuilder() query.Builder {
	return query.NewBuilder(
		nil,
		[]query.FilterParam{
			{DisplayName: "Email", FieldName: "email", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
			{DisplayName: "Name", FieldName: "name", InternalName: "username", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
		},
		[]query.SortParam{
			{DisplayName: "Email", FieldName: "email"},
			{DisplayName: "Name", FieldName: "username"},
			{DisplayName: "Created At", FieldName: "created_at"},
		},
		query.SortParam{DisplayName: "Name", FieldName: "username", Direction: query.SortDirectionAscending},
	)
}

func NewEventQueryBuilder() query.Builder {
	return query.NewBuilder(
		func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN tickets ON events.id = tickets.event_id").
				Group("events.id").
				Select("events.*", "MIN(tickets.price) AS min_price", "MAX(tickets.price) AS max_price")
		},
		[]query.FilterParam{
			{DisplayName: "Title", FieldName: "title", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
			{DisplayName: "Organizer", FieldName: "organizer", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
			{
				DisplayName:       "Start At Before",
				FieldName:         "start_at_before",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("start_at < ?", value)
				},
			},
			{
				DisplayName:       "Start At After",
				FieldName:         "start_at_after",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("start_at > ?", value)
				},
			},
			{
				DisplayName:       "End At Before",
				FieldName:         "end_at_before",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("end_at < ?", value)
				},
			},
			{
				DisplayName:       "End At After",
				FieldName:         "end_at_after",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("end_at > ?", value)
				},
			},
		},
		[]query.SortParam{
			{DisplayName: "Title", FieldName: "title"},
			{DisplayName: "Organizer", FieldName: "organizer"},
			{DisplayName: "Start At", FieldName: "start_at"},
			{DisplayName: "End At", FieldName: "end_at"},
			{DisplayName: "Price", FieldName: "price", Callback: func(db *gorm.DB, isDescending bool) *gorm.DB {
				if isDescending {
					db = db.Order("max_price DESC")
				} else {
					db = db.Order("min_price ASC")
				}
				return db
			}},
		},
		query.SortParam{DisplayName: "Start At", FieldName: "start_at", Direction: query.SortDirectionAscending},
	)
}

func NewEventApprovalQueryBuilder() query.Builder {
	return query.NewBuilder(
		nil,
		[]query.FilterParam{
			{DisplayName: "Event Title", FieldName: "title", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
			{
				DisplayName:       "Start At Before",
				FieldName:         "start_at_before",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("start_at < ?", value)
				},
			},
			{
				DisplayName:       "Start At After",
				FieldName:         "start_at_after",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("start_at > ?", value)
				},
			},
			{
				DisplayName:       "End At Before",
				FieldName:         "end_at_before",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("end_at < ?", value)
				},
			},
			{
				DisplayName:       "End At After",
				FieldName:         "end_at_after",
				DisplayFilterType: constants.FilterResponseDateRange,
				FilterType:        query.FilterTypeCustom,
				Callback: func(db *gorm.DB, value string) *gorm.DB {
					return db.Where("end_at > ?", value)
				},
			},
		},
		[]query.SortParam{
			{DisplayName: "Event Title", FieldName: "title"},
			{DisplayName: "Start At", FieldName: "start_at"},
			{DisplayName: "End At", FieldName: "end_at"},
			{DisplayName: "Created At", FieldName: "created_at"},
		},
		query.SortParam{DisplayName: "Created At", FieldName: "created_at", Direction: query.SortDirectionAscending},
	)
}
