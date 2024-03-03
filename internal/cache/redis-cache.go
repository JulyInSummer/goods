package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"goods_project/internal/config"
	"goods_project/internal/types"
)

type RedisCache struct {
	client *redis.Client
}

func NewCache(cfg *config.Config) *RedisCache {
	address := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   cfg.Redis.DB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	return &RedisCache{client: client}
}

func (r *RedisCache) Invalidate(ctx context.Context, id int) {
	key := strconv.Itoa(id)

	r.client.Del(ctx, key)
}

func (r *RedisCache) Get(ctx context.Context, id int) (*types.GoodResponse, error) {
	key := strconv.Itoa(id)

	res := r.client.Get(ctx, key)

	resBytes, err := res.Bytes()
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(resBytes)

	var response types.GoodResponse

	if err = gob.NewDecoder(b).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *RedisCache) Set(ctx context.Context, record *types.GoodResponse) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(record); err != nil {
		return err
	}

	id := record.Id
	recordKey := strconv.Itoa(id)

	return r.client.Set(ctx, recordKey, b.Bytes(), 60*time.Second).Err()
}
