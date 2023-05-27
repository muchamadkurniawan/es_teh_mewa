package service

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/stok/repository"
	"es_teh_mewa/stok/web"
)

type StokServiceImpl struct {
	DB   *sql.DB
	repo repository.StokRepository
}

func NewStokService(db *sql.DB, repo repository.StokRepository) StokService {
	return &StokServiceImpl{
		DB:   db,
		repo: repo,
	}
}

func (s *StokServiceImpl) Update(ctx context.Context, id_bahan int, total int) error {
	tx, err := s.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	err = s.repo.Update(ctx, tx, id_bahan, total)
	return err
}

func (s *StokServiceImpl) GetAll(ctx context.Context) []web.ResponseStok {
	tx, err := s.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	stoks := s.repo.GetByAll(ctx, tx)
	return stoks
}

func (s *StokServiceImpl) GetById(ctx context.Context, id int) web.ResponseStok {
	tx, err := s.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	stoks := s.repo.GetById(ctx, tx, id)
	return stoks
}

func (s *StokServiceImpl) GetAllDetail(ctx context.Context) []web.ResponseDetailStok {
	tx, err := s.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	stoks := s.repo.GetByAllDetail(ctx, tx)
	return stoks
}
