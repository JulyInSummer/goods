package pg_storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (p *pgStorage) GetGood(ctx context.Context, id int, projectId int) (error, *types.GoodResponse) {
	log := p.log.With("op", "pg-storage.GetGood")

	var resp types.GoodResponse

	if err := p.db.QueryRowContext(ctx, "SELECT * FROM goods WHERE id=$1 AND project_id=$2 AND removed=false", id, projectId).Scan(
		&resp.Id,
		&resp.ProjectId,
		&resp.Name,
		&resp.Description,
		&resp.Priority,
		&resp.Removed,
		&resp.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			return constants.NotFoundError, nil
		}
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	return nil, &resp
}

func (p *pgStorage) InternalGet(ctx context.Context, id int) (error, *types.GoodResponse) {
	log := p.log.With("op", "pg-storage.InternalGetGood")

	var resp types.GoodResponse

	if err := p.db.QueryRowContext(ctx, "SELECT id, project_id, name, description, priority, removed FROM goods WHERE id=$1", id).Scan(
		&resp.Id,
		&resp.ProjectId,
		&resp.Name,
		&resp.Description,
		&resp.Priority,
		&resp.Removed,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			return constants.NotFoundError, nil
		}
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	return nil, &resp
}
