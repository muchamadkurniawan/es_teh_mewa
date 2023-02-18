package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type BahanBakuRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, baku entity.BahanBaku)
	Update(ctx context.Context, tx *sql.Tx, baku entity.BahanBaku) entity.BahanBaku
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.BahanBakuFull
	FindById(ctx context.Context, tx *sql.Tx, id int) entity.BahanBakuFull
}
