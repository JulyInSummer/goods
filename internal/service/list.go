package service

import (
	"context"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (s *Service) ListGood(ctx context.Context, limit, offset int) (error, *types.ListResponse) {
	log := s.log.With("op", "service.ListGood")

	err, resp := s.data.GoodList(ctx, limit, offset)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	return nil, resp
}
