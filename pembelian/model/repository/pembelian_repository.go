package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/pembelian/model/entity"
	"eh_teh_mewa/pembelian/model/web"
)

type PembelianRepository interface {
	UpdatePembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error)
	DeletePembelian(ctx context.Context, tx *sql.Tx, id string) error
	InsertPembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error)
	FindByIdPembelian(ctx context.Context, tx *sql.Tx, id string) (entity.Pembelian, error)
	FindByAllPembelian(ctx context.Context, tx *sql.Tx) ([]entity.Pembelian, error)
	FindByAllPembelianByDate(ctx context.Context, tx *sql.Tx, filterAwal string, filterAkhir string) ([]entity.Pembelian, error)
}
