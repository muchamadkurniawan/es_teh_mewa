package service

import (
	"context"
	"eh_teh_mewa/produk/model/web"
)

type ProdukService interface {
	FindByAll(ctx context.Context) ([]web.ResponseProduk, error)
	FindById(ctx context.Context, id int) (web.ResponseProduk, error)
	Create(ctx context.Context, produk web.RequestProduk) (web.ResponseProduk, error)
	Update(ctx context.Context, produk web.ResponseProduk) (web.ResponseProduk, error)
	Delete(ctx context.Context, id int) error
}
