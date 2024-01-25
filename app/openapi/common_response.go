package openapi

type CommonResponse struct {
	Message string         `json:"message"`
	Status  ResponseStatus `json:"status"`
}
