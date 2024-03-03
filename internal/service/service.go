package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"
	"goods_project/internal/cache"
	"goods_project/internal/constants"
	"goods_project/internal/storage"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

type Service struct {
	log   *slog.Logger
	data  storage.Storage
	cache *cache.RedisCache
	pub   *nats.Conn
}

func NewService(log *slog.Logger, data storage.Storage, cache *cache.RedisCache, pub *nats.Conn) *Service {
	return &Service{
		log:   log,
		data:  data,
		cache: cache,
		pub:   pub,
	}
}

func (s *Service) GetGood(ctx context.Context, id, projectId int) (error, *types.GoodResponse) {
	log := s.log.With("op", "service.GetGood")

	res, err := s.cache.Get(ctx, id)
	if err == nil {
		return nil, res
	}

	err, res = s.data.GetGood(ctx, id, projectId)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	if err = s.cache.Set(ctx, res); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
	}

	return nil, res
}
