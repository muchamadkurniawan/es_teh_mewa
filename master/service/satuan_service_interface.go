package service

import (
	"context"
	"eh_teh_mewa/master/web"
)

type SatuanService interface {
	Save(ctx context.Context, request web.SatuanRequest)
	Update(ctx context.Context, response web.SatuanResponse)
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []web.SatuanResponse
	FindById(ctx context.Context, id int) web.SatuanResponse
}
