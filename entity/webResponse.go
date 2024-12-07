package entity

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"status"`
	Data    interface{} `json:"data"`
}
