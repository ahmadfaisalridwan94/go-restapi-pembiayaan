package definitions

type (
	Meta struct {
		Page      int `json:"page"`
		PerPage   int `json:"perPage"`
		TotalPage int `json:"totalPage"`
		Total     int `json:"total"`
	}
	SuccessResponse struct {
		ResponseCode string      `json:"ResponseCode"`
		ResponseDesc string      `json:"ResponseDesc"`
		Meta         interface{} `json:"meta,omitempty"`
		ResponseData interface{} `json:"ResponseData"`
	}
)
