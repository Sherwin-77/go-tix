package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    *Meta       `json:"meta,omitempty" swaggerignore:"true"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Page         int          `json:"page"`
	PerPage      int          `json:"per_page"`
	LastPage     int          `json:"last_page"`
	Total        int64        `json:"total"`
	Filters      []FilterMeta `json:"filters"`
	Sorts        []SortMeta   `json:"sorts"`
	SelectedSort string       `json:"selected_sort"`
	DefaultSort  string       `json:"default_sort"`
}

type FilterMeta struct {
	Name       string `json:"name"`
	FilterType string `json:"filter_type"`
	Label      string `json:"label"`
	Value      string `json:"value"`
}

type SortMeta struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

func NewResponse(code int, message string, data interface{}, meta *Meta) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func NewErrorResponse(code int, message string, data interface{}) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
