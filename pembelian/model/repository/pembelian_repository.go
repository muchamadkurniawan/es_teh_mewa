package repository

import (
	"context"
	"database/sql"
	webBahanBakuResponse "es_teh_mewa/master/web"
	"es_teh_mewa/pembelian/model/entity"
	"es_teh_mewa/pembelian/model/web"
)

type PembelianRepository interface {
	GetAllBahanBaku(ctx context.Context, tx *sql.Tx) []webBahanBakuResponse.BahanbakuResponse
	UpdatePembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error)
	DeletePembelian(ctx context.Context, tx *sql.Tx, id string) error
	InsertPembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error)
	UpdateStok(ctx context.Context, tx *sql.Tx, id_bahan int, jumlah int)
	FindByIdPembelian(ctx context.Context, tx *sql.Tx, id string) (entity.Pembelian, error)
	FindByAllPembelian(ctx context.Context, tx *sql.Tx) ([]web.PembelianResponseFull, error)
	FindByAllPembelianByDate(ctx context.Context, tx *sql.Tx, filterAwal string, filterAkhir string) ([]web.PembelianResponseFull, error)
}
