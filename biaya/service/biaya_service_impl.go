package service

import (
	"context"
	"database/sql"
	"es_teh_mewa/biaya/model/repository"
	"es_teh_mewa/biaya/web"
	"es_teh_mewa/helperMain"
)

type BiayaServiceImpl struct {
	repo repository.BiayaRepository
	DB   *sql.DB
}

func NewBiayaService(repo repository.BiayaRepository, DB *sql.DB) BiayaService {
	return &BiayaServiceImpl{
		repo: repo,
		DB:   DB,
	}
}

func (service *BiayaServiceImpl) GetBahanBakuNonProdukServ(ctx context.Context) []web.GetBahanBakuNonProdukRespon {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	bahan := service.repo.GetBahanBakuNonProdukRepo(ctx, tx)
	return bahan
}

func (service *BiayaServiceImpl) GetBiayaTodayServ(ctx context.Context) []web.GetBiayaTodayRespon {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	bahan := service.repo.GetBiayaTodayRepo(ctx, tx)
	return bahan
}
