package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/pembelian/model/entity"
	"eh_teh_mewa/pembelian/model/web"
)

type PembelianRepository interface {
	UpdatePembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error)
	DeletePembelian(ctx context.Context, tx *sql.Tx, id int) error
	InsertPembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error)
	FindByIdPembelian(ctx context.Context, tx *sql.Tx, id int32) (entity.Pembelian, error)
	FindByAllPembelian(ctx context.Context, tx *sql.Tx) ([]entity.Pembelian, error)
}
