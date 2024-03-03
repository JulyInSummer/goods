package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"goods_project/internal/constants"
	"goods_project/internal/http-responses"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Server) CreateGood(w http.ResponseWriter, r *http.Request) error {
	log := s.Log.With("op", "api.CreateGood")

	projId := r.URL.Query().Get("projectId")
	projIdInt, err := strconv.Atoi(projId)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return http_responses.BadRequest(w, constants.MessageInvalidQueryParams)
	}

	var payload types.CreateRequest
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return http_responses.BadRequest(w, constants.MessageInvalidRequestBody)
	}

	if err = s.Validate.Struct(payload); err != nil {
		return http_responses.ValidationError(w, err)
	}

	err, resp := s.Service.CreateGood(s.Ctx, projIdInt, &payload)
	if err != nil {
		return http_responses.InternalError(w)
	}

	return utils.WriteJSON(w, http.StatusCreated, resp)
}
