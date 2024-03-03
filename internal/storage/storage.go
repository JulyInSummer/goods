package storage

import (
	"context"

	"goods_project/internal/types"
)

type Storage interface {
	GoodCreate(ctx context.Context, projectId int, payload *types.CreateRequest) (error, *types.GoodResponse)
	GoodUpdate(ctx context.Context, id int, projectId int, payload *types.UpdateRequest) (error, *types.GoodResponse)
	GoodRemove(ctx context.Context, id int, projectId int) (error, *types.RemoveResponse)
	GetGood(ctx context.Context, id int, projectId int) (error, *types.GoodResponse)
	GoodList(ctx context.Context, limit int, offset int) (error, *types.ListResponse)
	Reprioritize(ctx context.Context, id int, projectId int, payload *types.ReprioritizeRequest) (error, *types.PrioritiesResponse)
	InternalGet(ctx context.Context, id int) (error, *types.GoodResponse)
}
