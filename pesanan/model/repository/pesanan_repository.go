package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/pesanan/model/entity"
	"eh_teh_mewa/pesanan/web"
)

type PesananRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) entity.PesananEntity
	FindAll(ctx context.Context, tx *sql.Tx) []entity.PesananEntity
	GetProdukJualsAll(ctx context.Context, tx *sql.Tx) []web.ProdukJual
	GetProdukJual(ctx context.Context, tx *sql.Tx, id int) web.ProdukJual
	Create(ctx context.Context, tx *sql.Tx, request web.PesananRequestDateString) error
	UpdatePembayaran(ctx context.Context, tx *sql.Tx, id int, pembayaran bool) error
	UpdateRekap(ctx context.Context, tx *sql.Tx, id int, id_rekap int) error
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
