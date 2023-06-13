package service

import (
	"context"
	web2 "es_teh_mewa/biaya/web"
	"es_teh_mewa/dashboard/web"
	web3 "es_teh_mewa/rekap/web"
)

type DashboardService interface {
	GetRekap(ctx context.Context) []web.ResponseDashboard
	GetRekapById(ctx context.Context, id int) web.ResponseDashboard
	GetBiayaRekapById(ctx context.Context, id int) ([]web2.GetBiayaTodayRespon, int)
	GetPesananRekapById(ctx context.Context, id int) ([]web3.AllPesananRekap, int)
	//Cetak(ctx context.Context, id int) error
}
