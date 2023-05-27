package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/pesanan/model/entity"
	"es_teh_mewa/pesanan/web"
)

type PesananRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) web.PesananRequestUpdate
	FindAll(ctx context.Context, tx *sql.Tx) []entity.PesananEntity
	FindAllPesananDetail(ctx context.Context, tx *sql.Tx) []web.PesananRequestSum
	FindAllDetailPesananId(ctx context.Context, tx *sql.Tx, id int, limit int) []web.PesananNamaBarang
	ShowAllDetailPesananId(ctx context.Context, tx *sql.Tx, id int) []web.DetailRespon
	GetProdukJualsAll(ctx context.Context, tx *sql.Tx) []web.ProdukJual
	GetProdukJual(ctx context.Context, tx *sql.Tx, id int) web.ProdukJual
	GetIdProduk(ctx context.Context, tx *sql.Tx) []string
	CreatePesanan(ctx context.Context, tx *sql.Tx, request web.PesananRequestDateString) int
	CreateDetail(ctx context.Context, tx *sql.Tx, request web.DetailRequest) (error, int)
	CreateBiayaPesanan(ctx context.Context, tx *sql.Tx, jumlah int, bahan int, harga int, id_detail int) (error, int)
	GetIdBahan(ctx context.Context, tx *sql.Tx, id int) int
	GetHargaBiaya(ctx context.Context, tx *sql.Tx, id int) int
	UpdatePembayaran(ctx context.Context, tx *sql.Tx, id int, pembayaran bool) error
	UpdateRekap(ctx context.Context, tx *sql.Tx, id int, id_rekap int) error
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	CreateDetailStok(ctx context.Context, tx *sql.Tx, BP int, jumlah int) error
}
