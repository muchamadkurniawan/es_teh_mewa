package repository

import (
	"context"
	"database/sql"
	web2 "es_teh_mewa/biaya/web"
	"es_teh_mewa/dashboard/web"
	web3 "es_teh_mewa/rekap/web"
)

type DashboardRepository interface {
	GetRekap(ctx context.Context, tx *sql.Tx) []web.ResponseDashboard
	GetRekapById(ctx context.Context, tx *sql.Tx, id int) web.ResponseDashboard
	GetBiayaRekapById(ctx context.Context, tx *sql.Tx, id int) []web2.GetBiayaTodayRespon
	GetPesananRekapById(ctx context.Context, tx *sql.Tx, id int) []web3.AllPesananRekap
}
