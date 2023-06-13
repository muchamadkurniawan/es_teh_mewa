package service

import (
	"context"
	"database/sql"
	BiayaRespon "es_teh_mewa/biaya/web"
	"es_teh_mewa/dashboard/repository"
	"es_teh_mewa/dashboard/web"
	"es_teh_mewa/helperMain"
	web3 "es_teh_mewa/rekap/web"
)

type DashboardServiceImpl struct {
	repo repository.DashboardRepository
	DB   *sql.DB
}

func NewDashboardService(repo repository.DashboardRepository, DB *sql.DB) DashboardService {
	return &DashboardServiceImpl{
		repo: repo,
		DB:   DB,
	}
}

func (d *DashboardServiceImpl) GetRekap(ctx context.Context) []web.ResponseDashboard {
	tx, err := d.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := d.repo.GetRekap(ctx, tx)
	return data
}

func (d *DashboardServiceImpl) GetRekapById(ctx context.Context, id int) web.ResponseDashboard {
	tx, err := d.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := d.repo.GetRekapById(ctx, tx, id)
	return data
}

func (d *DashboardServiceImpl) GetBiayaRekapById(ctx context.Context, id int) ([]BiayaRespon.GetBiayaTodayRespon, int) {
	tx, err := d.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := d.repo.GetBiayaRekapById(ctx, tx, id)
	var total int
	for _, datum := range data {
		total += datum.Total
	}
	return data, total
}

func (d *DashboardServiceImpl) GetPesananRekapById(ctx context.Context, id int) ([]web3.AllPesananRekap, int) {
	tx, err := d.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := d.repo.GetPesananRekapById(ctx, tx, id)
	var total int
	for _, datum := range data {
		total += datum.TotalPesanan
	}
	return data, total
}

//func (d *DashboardServiceImpl) Cetak(ctx context.Context, id int) error {
//	tx, err := d.DB.Begin()
//
//	if err != nil {
//		return err
//	}
//	return nil
//}
