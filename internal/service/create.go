package service

import (
	"context"
	"encoding/json"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Service) CreateGood(ctx context.Context, projectId int, payload *types.CreateRequest) (error, *types.GoodResponse) {
	log := s.log.With("op", "service.CreateGood")

	err, res := s.data.GoodCreate(ctx, projectId, payload)
	if err != nil {
		log.Error(err.Error())
		return err, nil
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	if err = s.pub.Publish(constants.PublisherLogsSubject, data); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	return nil, res
}
