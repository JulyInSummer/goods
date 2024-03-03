package constants

import "github.com/pkg/errors"

const (
	ApiInternalError = iota
	ApiBadRequest
	ApiValidationError
	ApiNotFound
)

const (
	ActionFailedErrorPrompt = "failed to perform an action"

	MessageInvalidQueryParams = "invalid query params"
	MessageInvalidRequestBody = "invalid request body"
	MessageInternalError      = "internal server error"
	MessageNotFound           = "not found"
	MessageValidationError    = "validation error"

	DefaultLimit  = 10
	DefaultOffset = 1

	PublisherLogsSubject = "logs"
)

var (
	NotFoundError = errors.New("Record not found")
)
