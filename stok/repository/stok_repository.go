package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/stok/web"
)

type StokRepository interface {
	Update(ctx context.Context, tx *sql.Tx, id_bahan int, total int) error
	GetByAll(ctx context.Context, tx *sql.Tx) []web.ResponseStok
	GetByAllDetail(ctx context.Context, tx *sql.Tx) []web.ResponseDetailStok
	GetById(ctx context.Context, tx *sql.Tx, id int) web.ResponseStok
}
