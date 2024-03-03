package pg_storage

import (
	"context"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (p *pgStorage) GoodCreate(ctx context.Context, projectId int, payload *types.CreateRequest) (error, *types.GoodResponse) {
	log := p.log.With("op", "pg_store.GoodCreate")

	stmt, err := p.db.Prepare("INSERT INTO goods(project_id, name, description) VALUES($1, $2, $3) RETURNING *")
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	defer stmt.Close()

	var response types.GoodResponse
	if err = stmt.QueryRowContext(ctx, projectId, payload.Name, payload.Description).Scan(
		&response.Id,
		&response.ProjectId,
		&response.Name,
		&response.Description,
		&response.Priority,
		&response.Removed,
		&response.CreatedAt,
	); err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return err, nil
	}

	return nil, &response
}
