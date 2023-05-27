package service

import (
	"context"
	"es_teh_mewa/pesanan/web"
)

type PesananService interface {
	FindById(ctx context.Context, id int) (web.PesananRequestUpdate, error)
	FindAll(ctx context.Context) ([]web.PesananResponseDateString, error)
	FindAllPesananDetail(ctx context.Context) []web.PesananDetail
	FindPesananDetail(ctx context.Context, id int) web.PesananDetailUpdate
	GetProdukJualsAll(ctx context.Context) []web.ProdukJual
	GetProdukJual(ctx context.Context, id int) web.ProdukJual
	GetIdProduk(ctx context.Context) []string
	CreatePesanan(ctx context.Context, request web.PesananRequestDateString) int
	CreateDetail(ctx context.Context, request web.DetailRequest) (error, int)
	CreateBiaya(ctx context.Context, bahan int, jumlah int, harga int, id_detail int) error
	GetIdBahan(ctx context.Context, id int) int
	GetHargaBiaya(ctx context.Context, id int) int
	UpdatePembayaran(ctx context.Context, id int, pembayaran bool) error
	UpdateRekap(ctx context.Context, id int, id_rekap int) error
	Delete(ctx context.Context, id int) error
	Cetak(ctx context.Context, id int) error
	UpdateStok(ctx context.Context, BP int, jumlah int)
}
