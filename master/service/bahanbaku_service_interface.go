package service

import (
	"context"
	"eh_teh_mewa/master/web"
)

type BahanbakuService interface {
	Save(ctx context.Context, request web.BahanbakuRequest)
	Update(ctx context.Context, response web.BahanbakuRequest)
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []web.BahanbakuFullResponse
	FindById(ctx context.Context, id int) web.BahanbakuResponse
	GetSatuan(ctx context.Context) []web.SatuanResponse
}
