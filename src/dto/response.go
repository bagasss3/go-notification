package dto

type ResponseSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
