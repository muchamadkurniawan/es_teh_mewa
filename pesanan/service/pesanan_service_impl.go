package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/pesanan/helper"
	"eh_teh_mewa/pesanan/model/repository"
	"eh_teh_mewa/pesanan/web"
)

type PesananServiceImpl struct {
	PesananRepository repository.PesananRepository
	db                *sql.DB
}

func NewPesananServie(repo repository.PesananRepository, DB *sql.DB) PesananService {
	return &PesananServiceImpl{
		PesananRepository: repo,
		db:                DB,
	}
}

func (service *PesananServiceImpl) GetIdProduk(ctx context.Context) []string {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.GetIdProduk(ctx, tx)
	return find
}

func (service *PesananServiceImpl) GetProdukJualsAll(ctx context.Context) []web.ProdukJual {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.GetProdukJualsAll(ctx, tx)
	return find
}

func (service *PesananServiceImpl) GetProdukJual(ctx context.Context, id int) web.ProdukJual {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.GetProdukJual(ctx, tx, id)
	return find
}

func (service *PesananServiceImpl) FindById(ctx context.Context, id int) (web.PesananResponseDateString, error) {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.FindById(ctx, tx, id)
	return helper.PesananEntityToResponseString(find), nil
}

func (service *PesananServiceImpl) FindAll(ctx context.Context) ([]web.PesananResponseDateString, error) {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	all := service.PesananRepository.FindAll(ctx, tx)
	var pesanans []web.PesananResponseDateString
	for _, pesanan := range all {
		newPesanan := helper.PesananEntityToResponseString(pesanan)
		pesanans = append(pesanans, newPesanan)
	}
	return pesanans, nil
}

func (service *PesananServiceImpl) CreatePesanan(ctx context.Context, request web.PesananRequestDateString) int {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	id := service.PesananRepository.CreatePesanan(ctx, tx, request)
	return id
}

func (service *PesananServiceImpl) UpdatePembayaran(ctx context.Context, id int, pembayaran bool) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.UpdatePembayaran(ctx, tx, id, pembayaran)
	helperMain.PanicIfError(err)
	return err
}

func (service *PesananServiceImpl) UpdateRekap(ctx context.Context, id int, id_rekap int) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.UpdateRekap(ctx, tx, id, id_rekap)
	helperMain.PanicIfError(err)
	return err
}

func (service *PesananServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.Delete(ctx, tx, id)
	helperMain.PanicIfError(err)
	return err
}
