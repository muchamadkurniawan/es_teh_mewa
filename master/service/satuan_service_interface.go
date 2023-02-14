package service

import (
	"context"
	"eh_teh_mewa/master/web"
)

type SatuanService interface {
	Save(ctx context.Context, request web.SatuanRequest)
	Update(ctx context.Context, response web.SatuanResponse)
	Delete(ctx context.Context, id int32)
	FindAll(ctx context.Context) []map[string]interface{}
	FindById(ctx context.Context, id int32) map[string]interface{}
}
