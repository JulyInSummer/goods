package api

import (
	"net/http"

	"goods_project/internal/constants"
	"goods_project/internal/http-responses"
	"goods_project/internal/utils"
)

func (s *Server) ListGoods(w http.ResponseWriter, r *http.Request) error {
	log := s.Log.With("op", "api.ListGoods")

	err, limit := utils.QueryParamToInt(r, "limit", constants.DefaultLimit)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	err, offset := utils.QueryParamToInt(r, "offset", constants.DefaultOffset)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	err, resp := s.Service.ListGood(s.Ctx, *limit, *offset)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.InternalError(w)
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}
