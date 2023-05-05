package service

import (
	"context"
	"database/sql"
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

func (repository *RekapServiceImpl) GetAllPesananToday(ctx context.Context) []web.AllPesananRekap {
	tx, err := repository.DB.Begin()
	helperMain.PanicIfError(err)
	data := repository.RepositoryRekap.PesananNonRekapByDate(ctx, tx)
	return data
}
