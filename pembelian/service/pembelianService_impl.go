package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/pembelian/helper"
	"eh_teh_mewa/pembelian/model/repository"
	"eh_teh_mewa/pembelian/model/web"
)

type PembelianServiceImpl struct {
	PembelianRepository repository.PembelianRepository
	DB                  *sql.DB
}

func NewPembelianService(pembelianRepository repository.PembelianRepository, DB *sql.DB) PembelianService {
	return &PembelianServiceImpl{
		PembelianRepository: pembelianRepository,
		DB:                  DB,
	}
}

func (service *PembelianServiceImpl) Add(ctx context.Context, index int, id int) int {
	return index + id
}

func (service *PembelianServiceImpl) FindById(ctx context.Context, id int) (web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.FindByIdPembelian(ctx, tx, int32(id))

	return helper.ToPembelianResponse(pembelian), err
}

func (service *PembelianServiceImpl) FindByAll(ctx context.Context) ([]web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.FindByAllPembelian(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToPembelianAllResponse(pembelian), nil
}

func (service *PembelianServiceImpl) Store(ctx context.Context, request web.PembelianCreateRequest) (web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.InsertPembelian(ctx, tx, request)
	helper.PanicIfError(err)
	return helper.RequestToResponse(pembelian), nil
}

func (service *PembelianServiceImpl) Update(ctx context.Context, request web.PembelianCreateRequest) (web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.UpdatePembelian(ctx, tx, request)
	helper.PanicIfError(err)
	return helper.RequestToResponse(pembelian), nil
}
func (service *PembelianServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ErrorTx(tx)

	err = service.PembelianRepository.DeletePembelian(ctx, tx, id)
	helper.PanicIfError(err)
	return nil
}
