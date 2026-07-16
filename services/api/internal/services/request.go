package services

type ListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
