package entity

type WebResponse struct {
	Code    int         `json:"statu_code"`
	Message string      `json:"status"`
	Data    interface{} `json:"data"`
}
