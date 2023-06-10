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

func (service *BiayaServiceImpl) GetBahanBaku(ctx context.Context) []web.GetBahanBakuNonProdukRespon {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	bahan := service.repo.GetBahanBaku(ctx, tx)
	return bahan
}

func (service *BiayaServiceImpl) GetBiayaTodayServ(ctx context.Context) []web.GetBiayaTodayRespon {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	bahan := service.repo.GetBiayaTodayRepo(ctx, tx)
	return bahan
}

func (service *BiayaServiceImpl) CreateBiaya(ctx context.Context, create web.BiayaRequestCreate) error {
	tx, err := service.DB.Begin()
	defer helperMain.ErrorTx(tx)
	err = service.repo.CreateBiaya(ctx, tx, create)
	helperMain.PanicIfError(err)
	return err
}

func (service *BiayaServiceImpl) FindById(ctx context.Context, id int) web.GetBiayaRespon {
	tx, err := service.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := service.repo.FindById(ctx, tx, id)
	return data
}

func (service *BiayaServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	return service.repo.Delete(ctx, tx, id)
}
