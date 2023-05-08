package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/biaya/web"
)

type BiayaRepository interface {
	GetBahanBakuNonProdukRepo(ctx context.Context, tx *sql.Tx) []web.GetBahanBakuNonProdukRespon
	GetBiayaTodayRepo(ctx context.Context, tx *sql.Tx) []web.GetBiayaTodayRespon
	CreateBiaya(ctx context.Context, tx *sql.Tx, request web.BiayaRequestCreate) error
}
