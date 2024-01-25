package openapi

type ResponseStatus string

// List of ResponseStatus
const (
	RESPONSESTATUS_SUCCESS ResponseStatus = "success"
	RESPONSESTATUS_FAIL    ResponseStatus = "fail"
	RESPONSESTATUS_ERROR   ResponseStatus = "error"
)
