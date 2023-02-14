package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type SatuanRepository interface {
	InsertSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan) error
	UpdateSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan) error
	DeleteSatuan(ctx context.Context, tx *sql.Tx, id int32) error
	FindAllSatuan(ctx context.Context, tx *sql.Tx) ([]entity.Satuan, error)
	FindByIdSatuan(ctx context.Context, tx *sql.Tx, id int32) (entity.Satuan, error)
}
