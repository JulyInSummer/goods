package pg_storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (p *pgStorage) GoodUpdate(ctx context.Context, id int, projectId int, payload *types.UpdateRequest) (error, *types.GoodResponse) {
	log := p.log.With("op", "pg-storage.GoodUpdate")

	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	var resp types.GoodResponse

	if payload.Description != nil {
		if err = tx.QueryRowContext(
			ctx,
			"UPDATE goods SET name=$1, description=$2 WHERE id=$3 AND project_id=$4 AND removed=false RETURNING *",
			payload.Name,
			payload.Description,
			id,
			projectId,
		).Scan(
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
				tx.Rollback()
				return constants.NotFoundError, nil
			}
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			tx.Rollback()
			return err, nil
		}
	} else {
		if err = tx.QueryRowContext(
			ctx,
			"UPDATE goods SET name=$1 WHERE id=$2 AND project_id=$3 AND removed=false RETURNING *",
			payload.Name,
			id,
			projectId,
		).Scan(
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
				tx.Rollback()
				return constants.NotFoundError, nil
			}
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			tx.Rollback()
			return err, nil
		}
	}

	tx.Commit()
	return nil, &resp
}
