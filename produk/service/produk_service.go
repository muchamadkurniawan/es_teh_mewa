package service

import (
	"context"
	"eh_teh_mewa/produk/model/entity"
)

type ProdukService interface {
	FindByAll(ctx context.Context) ([]entity.Produk, error)
	FindById(ctx context.Context, id int) (entity.Produk, error)
	Create(ctx context.Context, produk entity.Produk) (entity.Produk, error)
	Update(ctx context.Context, produk entity.Produk) (entity.Produk, error)
	Delete(ctx context.Context, id int) error
}
