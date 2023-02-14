package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type SatuanRepository interface {
	InsertSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan)
	UpdateSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan)
	DeleteSatuan()
	FindAllSatuan()
	FindByIdSatuan()
}
