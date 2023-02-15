package service

import (
	"context"
	"eh_teh_mewa/master/web"
)

type BahanbakuService interface {
	Save(ctx context.Context, request web.BahanbakuRequest)
	Update(ctx context.Context, response web.BahanbakuResponse)
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []map[string]interface{}
	FindById(ctx context.Context, id int) map[string]interface{}
}
