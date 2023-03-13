package service

import (
	"context"
	"eh_teh_mewa/pesanan/web"
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
	CreateDetail(ctx context.Context, request web.DetailRequest) error
	UpdatePembayaran(ctx context.Context, id int, pembayaran bool) error
	UpdateRekap(ctx context.Context, id int, id_rekap int) error
	Delete(ctx context.Context, id int) error
}
