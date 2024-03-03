package service

import (
	"context"
	"encoding/json"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Service) UpdateGood(ctx context.Context, id, projectId int, payload *types.UpdateRequest) (error, *types.GoodResponse) {
	log := s.log.With("op", "service.UpdateGood")

	err, res := s.data.GoodUpdate(ctx, id, projectId, payload)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	s.cache.Invalidate(ctx, id)

	data, err := json.Marshal(res)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	if err = s.pub.Publish(constants.PublisherLogsSubject, data); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	return nil, res
}
