package service

import (
	"context"
	"es_teh_mewa/biaya/web"
)

type BiayaService interface {
	GetBahanBakuNonProdukServ(ctx context.Context) []web.GetBahanBakuNonProdukRespon
	GetBiayaTodayServ(ctx context.Context) []web.GetBiayaTodayRespon
}
