package service

import (
	"context"
	"database/sql"
	BiayaRespon "es_teh_mewa/biaya/web"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/rekap/model/repository"
	"es_teh_mewa/rekap/web"
)

type RekapServiceImpl struct {
	RepositoryRekap repository.RekapRepository
	DB              *sql.DB
}

func NewRekapServiceImpl(repo repository.RekapRepository, DB *sql.DB) RekapService {
	return &RekapServiceImpl{
		RepositoryRekap: repo,
		DB:              DB,
	}
}

func (repository *RekapServiceImpl) GetAllPesananToday(ctx context.Context) ([]web.AllPesananRekap, int) {
	tx, err := repository.DB.Begin()
	helperMain.PanicIfError(err)
	data := repository.RepositoryRekap.PesananNonRekapByDate(ctx, tx)
	total := 0
	for _, datum := range data {
		total += datum.TotalPesanan
	}
	return data, total
}

func (repository *RekapServiceImpl) GetAllBiayaToday(ctx context.Context) ([]BiayaRespon.GetBiayaTodayRespon, int) {
	tx, err := repository.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := repository.RepositoryRekap.BiayaNonRekapByDate(ctx, tx)
	var total int
	for _, datum := range data {
		total += datum.Total
	}
	return data, total
}
