package request

type PaginationRequest struct {
	Page     int32 `json:"page" validate:"required,gte=1"`
	PageSize int32 `json:"pageSize" validate:"required,gte=10,lte=100"`
}
