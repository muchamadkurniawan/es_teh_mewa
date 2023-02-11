package service

import (
	"context"
	"eh_teh_mewa/pembelian/model/web"
)

type PembelianService interface {
	FindById(ctx context.Context, id string) (web.PembelianUpdateResponse, error)
	FindByAll(ctx context.Context, filterAwal string, filterAkhir string) ([]web.PembelianResponse, error)
	Store(ctx context.Context, request web.PembelianCreateRequest) (web.PembelianResponse, error)
	Update(ctx context.Context, response web.PembelianCreateRequest) (web.PembelianResponse, error)
	Delete(ctx context.Context, id string) error
}
