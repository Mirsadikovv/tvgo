package utils

type ID HttpResponseID[int64]                // @name ID
type ErrorResponse HttpErrorResponse[string] // @name ErrorResponse

type numberOrStr interface {
	int | int32 | int64 | int16 | int8 |
		uint | uint32 | uint64 | uint8 |
		string
}

type HttpResponseID[T numberOrStr] struct {
	Id T `json:"id"`
}

type HttpErrorResponse[T any] struct {
	Message T `json:"message"`
}
