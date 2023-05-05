package service

import (
	"context"
	"es_teh_mewa/rekap/web"
)

type RekapService interface {
	GetAllPesananToday(ctx context.Context) []web.AllPesananRekap
}
