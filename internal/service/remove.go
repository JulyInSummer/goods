package service

import (
	"context"
	"encoding/json"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Service) RemoveGood(ctx context.Context, id, projectId int) (error, *types.RemoveResponse) {
	log := s.log.With("op", "service.RemoveGood")

	err, resp := s.data.GoodRemove(ctx, id, projectId)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	s.cache.Invalidate(ctx, id)

	err, fromDb := s.data.InternalGet(ctx, id)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	data, err := json.Marshal(fromDb)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	if err = s.pub.Publish(constants.PublisherLogsSubject, data); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	return nil, resp
}
