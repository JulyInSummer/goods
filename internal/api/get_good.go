package api

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"goods_project/internal/constants"
	"goods_project/internal/http-responses"
	"goods_project/internal/utils"
)

func (s *Server) GetGood(w http.ResponseWriter, r *http.Request) error {
	log := s.Log.With("op", "api.GetGood")

	id := r.URL.Query().Get("id")
	projectId := r.URL.Query().Get("projectId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	projectIdInt, err := strconv.Atoi(projectId)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	err, res := s.Service.GetGood(s.Ctx, idInt, projectIdInt)
	if err != nil {
		if errors.Is(err, constants.NotFoundError) {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			return http_responses.NotFound(w)
		}
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.InternalError(w)
	}

	return utils.WriteJSON(w, http.StatusOK, res)
}
