package pg_storage

import (
	"context"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (p *pgStorage) GoodList(ctx context.Context, limit, offset int) (error, *types.ListResponse) {
	log := p.log.With("op", "pg-storage.GoodList")

	resp := types.ListResponse{
		Meta: &types.Meta{
			Limit:  limit,
			Offset: offset,
		},
		Goods: []*types.GoodResponse{},
	}

	if err := p.db.QueryRowContext(ctx, "SELECT COUNT(*) AS total, COUNT(CASE WHEN removed=true THEN 1 END) AS removed FROM goods").Scan(
		&resp.Meta.Total,
		&resp.Meta.Removed,
	); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	rows, err := p.db.QueryContext(ctx, "SELECT * FROM goods WHERE removed=false LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	for rows.Next() {
		var good types.GoodResponse
		if err = rows.Scan(&good.Id, &good.ProjectId, &good.Name, &good.Description, &good.Priority, &good.Removed, &good.CreatedAt); err != nil {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			return err, nil
		}

		resp.Goods = append(resp.Goods, &good)
	}

	return nil, &resp
}
