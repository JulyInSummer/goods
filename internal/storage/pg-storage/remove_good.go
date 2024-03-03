package pg_storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (p *pgStorage) GoodRemove(ctx context.Context, id int, projectId int) (error, *types.RemoveResponse) {
	log := p.log.With("op", "pg-storage.GoodRemove")

	stmt, err := p.db.Prepare("UPDATE goods SET removed=true WHERE id=$1 AND project_id=$2 AND removed=false RETURNING id, project_id, removed")
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	var resp types.RemoveResponse
	if err = stmt.QueryRowContext(ctx, id, projectId).Scan(
		&resp.Id,
		&resp.ProjectId,
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
