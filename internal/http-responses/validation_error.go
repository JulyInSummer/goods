package http_responses

import (
	"net/http"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func ValidationError(w http.ResponseWriter, err error) error {
	return utils.WriteJSON(w, http.StatusBadRequest, types.ApiError{
		Code:    constants.ApiBadRequest,
		Message: constants.MessageValidationError,
		Details: utils.ParseValidationErrors(err),
	})
}
