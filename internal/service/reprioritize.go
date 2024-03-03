package service

import (
	"context"
	"encoding/json"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Service) Reprioritize(ctx context.Context, id, projectId int, payload *types.ReprioritizeRequest) (error, *types.PrioritiesResponse) {
	log := s.log.With("op", "service.Reprioritize")

	err, resp := s.data.Reprioritize(ctx, id, projectId, payload)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	for _, p := range resp.Priorities {
		s.cache.Invalidate(ctx, p.Id)

		dbErr, fromDb := s.data.InternalGet(ctx, p.Id)
		if dbErr != nil {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		}

		data, mErr := json.Marshal(fromDb)
		if mErr != nil {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		}

		if err = s.pub.Publish(constants.PublisherLogsSubject, data); err != nil {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		}
	}

	return nil, resp
}
