package repository

import (
	"context"
	"database/sql"
	BiayaRespon "es_teh_mewa/biaya/web"
	"es_teh_mewa/rekap/web"
)

type RekapRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) (web.RekapResponseTime, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]web.RekapResponseTime, error)
	PesananNonRekapByDate(ctx context.Context, tx *sql.Tx) []web.AllPesananRekap
	BiayaNonRekapByDate(ctx context.Context, tx *sql.Tx) []BiayaRespon.GetBiayaTodayRespon
	Create(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error
	Update(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
