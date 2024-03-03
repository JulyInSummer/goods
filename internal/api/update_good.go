package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"goods_project/internal/constants"
	"goods_project/internal/http-responses"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Server) UpdateGood(w http.ResponseWriter, r *http.Request) error {
	log := s.Log.With("op", "api.UpdateGood")

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	projectId := r.URL.Query().Get("projectId")
	projectIdInt, err := strconv.Atoi(projectId)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	var payload types.UpdateRequest
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	if err = s.Validate.Struct(payload); err != nil {
		return http_responses.ValidationError(w, err)
	}

	err, resp := s.Service.UpdateGood(s.Ctx, idInt, projectIdInt, &payload)
	if err != nil {
		if errors.Is(err, constants.NotFoundError) {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			return http_responses.NotFound(w)
		}
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.InternalError(w)
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}
