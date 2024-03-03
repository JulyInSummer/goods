package http_responses

import (
	"net/http"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func BadRequest(w http.ResponseWriter, msg string) error {
	return utils.WriteJSON(w, http.StatusBadRequest, types.ApiError{
		Code:    constants.ApiBadRequest,
		Message: msg,
		Details: nil,
	})
}
