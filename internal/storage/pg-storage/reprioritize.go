package pg_storage

import (
	"context"

	"goods_project/internal/constants"
	"goods_project/internal/types"
	"goods_project/internal/utils"
)

func (p *pgStorage) Reprioritize(ctx context.Context, id int, projectId int, payload *types.ReprioritizeRequest) (error, *types.PrioritiesResponse) {
	log := p.log.With("op", "pg-storage.Reprioritize")

	priorities := types.PrioritiesResponse{
		Priorities: []*types.Priority{},
	}

	rows, err := p.db.QueryContext(ctx, "SELECT updated_id AS id, updated_priority AS priority FROM update_priority($1, $2, $3)", id, projectId, payload.NewPriority)
	if err != nil {
		log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
		return constants.NotFoundError, nil
	}

	for rows.Next() {
		var data types.Priority

		if err = rows.Scan(&data.Id, &data.PriorityNum); err != nil {
			log.Error(constants.ActionFailedErrorPrompt, utils.SlErr(err))
			return err, nil
		}

		priorities.Priorities = append(priorities.Priorities, &data)
	}

	return nil, &priorities

}
