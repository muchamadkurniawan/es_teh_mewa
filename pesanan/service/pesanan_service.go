package service

import (
	"context"
	"eh_teh_mewa/pesanan/web"
)

type PesananService interface {
	FindById(ctx context.Context, id int) (web.PesananResponseDateString, error)
	FindAll(ctx context.Context) ([]web.PesananResponseDateString, error)
	GetProdukJualsAll(ctx context.Context) []web.ProdukJual
	GetProdukJual(ctx context.Context, id int) web.ProdukJual
	GetIdProduk(ctx context.Context) []string
	CreatePesanan(ctx context.Context, request web.PesananRequestDateString) int
	UpdatePembayaran(ctx context.Context, id int, pembayaran bool) error
	UpdateRekap(ctx context.Context, id int, id_rekap int) error
	Delete(ctx context.Context, id int) error
}
