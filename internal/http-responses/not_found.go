package http_responses

import (
	"net/http"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func NotFound(w http.ResponseWriter) error {
	return utils.WriteJSON(w, http.StatusNotFound, types.ApiError{
		Code:    constants.ApiNotFound,
		Message: constants.MessageNotFound,
		Details: nil,
	})
}
