package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
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

func (service *PembelianServiceImpl) FindById(ctx context.Context, id string) (web.PembelianUpdateResponse, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.FindByIdPembelian(ctx, tx, id)

	return helper.ToPembelianUpdateResponse(pembelian), err
}

func (service *PembelianServiceImpl) FindByAll(ctx context.Context, filterAwal string, filterAkhir string) ([]web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helper.ErrorTx(tx)
	if filterAwal == "" || filterAkhir == "" {
		pembelian, err := service.PembelianRepository.FindByAllPembelian(ctx, tx)
		helperMain.PanicIfError(err)
		return helper.ToPembelianAllResponse(pembelian), nil
	} else {
		pembelian, err := service.PembelianRepository.FindByAllPembelianByDate(ctx, tx, filterAwal, filterAkhir)
		helperMain.PanicIfError(err)
		return helper.ToPembelianAllResponse(pembelian), nil
	}

}

func (service *PembelianServiceImpl) Store(ctx context.Context, request web.PembelianCreateRequest) (web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.InsertPembelian(ctx, tx, request)
	helperMain.PanicIfError(err)
	return helper.RequestToResponse(pembelian), nil
}

func (service *PembelianServiceImpl) Update(ctx context.Context, request web.PembelianCreateRequest) (web.PembelianResponse, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helper.ErrorTx(tx)

	pembelian, err := service.PembelianRepository.UpdatePembelian(ctx, tx, request)
	helperMain.PanicIfError(err)
	return helper.RequestToResponse(pembelian), nil
}
func (service *PembelianServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helper.ErrorTx(tx)

	err = service.PembelianRepository.DeletePembelian(ctx, tx, id)
	helperMain.PanicIfError(err)
	return nil
}
