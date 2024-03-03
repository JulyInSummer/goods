package http_responses

import (
	"net/http"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func InternalError(w http.ResponseWriter) error {
	return utils.WriteJSON(w, http.StatusInternalServerError, types.ApiError{
		Code:    constants.ApiInternalError,
		Message: constants.MessageInternalError,
		Details: nil,
	})
}
