package service

import (
	"context"
	"es_teh_mewa/biaya/web"
)

type BiayaService interface {
	GetBahanBakuNonProdukServ(ctx context.Context) []web.GetBahanBakuNonProdukRespon
	GetBiayaTodayServ(ctx context.Context) []web.GetBiayaTodayRespon
	CreateBiaya(ctx context.Context, create web.BiayaRequestCreate) error
	FindById(ctx context.Context, id int) web.GetBiayaRespon
	Delete(ctx context.Context, id int) error
}
