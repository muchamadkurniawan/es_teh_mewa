package service

import (
	"context"
	BiayaRespon "es_teh_mewa/biaya/web"
	"es_teh_mewa/rekap/web"
)

type RekapService interface {
	GetAllPesananToday(ctx context.Context) ([]web.AllPesananRekap, int)
	GetAllBiayaToday(ctx context.Context) ([]BiayaRespon.GetBiayaTodayRespon, int)
	Create(ctx context.Context, keterangan string) (int, error)
	UpdateRekapPesananBiaya(ctx context.Context, id int, id_pesanan []string, id_biaya []string) error
}
