package service

import (
	"context"
	"es_teh_mewa/stok/web"
)

type StokService interface {
	Update(ctx context.Context, id_bahan int, total int) error
	GetAll(ctx context.Context) []web.ResponseStok
	GetAllDetail(ctx context.Context) []web.ResponseDetailStok
	GetById(ctx context.Context, id int) web.ResponseStok
}
