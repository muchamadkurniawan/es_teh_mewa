package service

import (
	"context"
	webBahan "eh_teh_mewa/master/web"
	"eh_teh_mewa/produk/model/entity"
	"eh_teh_mewa/produk/model/web"
)

type ProdukService interface {
	FindByAll(ctx context.Context) ([]web.ResponseProdukFull, error)
	FindById(ctx context.Context, id int) (entity.Produk, error)
	Create(ctx context.Context, produk entity.Produk)
	Update(ctx context.Context, produk entity.Produk)
	Delete(ctx context.Context, id int) error
	FindAllByBahan(ctx context.Context, barang int) ([]web.ResponseProdukFull, error)
	CheckByBahan(ctx context.Context, barang int) (entity.Produk, error)
	GetBahan(ctx context.Context) ([]webBahan.BahanbakuFullResponse, error)
}
