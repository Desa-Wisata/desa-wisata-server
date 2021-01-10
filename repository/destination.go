package repository

import (
	"context"

	"github.com/desa-wisata/desa-wisata-server/domain"
	"github.com/desa-wisata/library/logger"
	"github.com/desa-wisata/library/utils"
)

func (r *repo) GetDestination(ctx context.Context) (context.Context, []domain.Destination, error) {
	var dest []domain.Destination

	start := utils.TimeNow()

	res := r.db.Table("destination").Find(&dest)

	execTime := utils.TimeDuration(start)
	ctx = logger.RecordDatabase(ctx, execTime, "getDestination", utils.ErrorString(res.Error))

	return ctx, dest, nil
}
