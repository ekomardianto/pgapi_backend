package models

type Meta struct {
	PerPage   int `json:"per_page"`
	Page      int `json:"page"`
	CountData int `json:"count_data"`
	CountPage int `json:"count_page"`
}
type ResponseSuccess struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}
type ResponseSuccessIpay struct {
	StatusCode int         `json:"Status"`
	Success    bool        `json:"Success"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data,omitempty"`
}

type ResponseFailed struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type PaginateResponseSuccess struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Meta       Meta        `json:"meta"`
}
type PaginateResponseFailed struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Meta       Meta   `json:"meta"`
}
