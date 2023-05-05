package service

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	repositoryBahan "es_teh_mewa/master/repository"
	webBahan "es_teh_mewa/master/web"
	"es_teh_mewa/produk/model/entity"
	"es_teh_mewa/produk/model/repository"
	"es_teh_mewa/produk/model/web"
)

type ProdukServiceImpl struct {
	ProdukRepository repository.ProdukRepository
	DB               *sql.DB
}

func NewProdukService(produkRepository repository.ProdukRepository, DB *sql.DB) ProdukService {
	return &ProdukServiceImpl{
		ProdukRepository: produkRepository,
		DB:               DB,
	}
}

func (service *ProdukServiceImpl) GetBahan(ctx context.Context) ([]webBahan.BahanbakuFullResponse, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	repo := repositoryBahan.NewBahanRepository()
	bahan := repo.FindAll(context.Background(), tx)
	bahans := helperMain.ToBahanRresponses(bahan)
	return bahans, nil

}

func (service *ProdukServiceImpl) FindByAll(ctx context.Context) ([]web.ResponseProdukFull, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	produks := service.ProdukRepository.FindAllProduk(ctx, tx)
	helperMain.PanicIfError(err)
	return produks, err
}

func (service *ProdukServiceImpl) FindAllByBahan(ctx context.Context, barang int) ([]web.ResponseProdukFull, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	produks := service.ProdukRepository.FindAllProdukByBahan(ctx, tx, barang)
	helperMain.PanicIfError(err)
	return produks, err
}

func (service *ProdukServiceImpl) CheckByBahan(ctx context.Context, barang int) (entity.Produk, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	produks, err := service.ProdukRepository.FindProdukByBahan(ctx, tx, barang)
	helperMain.PanicIfError(err)
	return produks, err
}

func (service *ProdukServiceImpl) FindById(ctx context.Context, id int) (entity.Produk, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	var produk entity.Produk
	produk = service.ProdukRepository.FindProdukById(ctx, tx, id)
	return produk, err
}

func (service *ProdukServiceImpl) Create(ctx context.Context, request entity.Produk) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	service.ProdukRepository.CreateProduk(ctx, tx, request)
}

func (service *ProdukServiceImpl) Update(ctx context.Context, produk entity.Produk) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	service.ProdukRepository.UpdateProduk(ctx, tx, produk)
}

func (service *ProdukServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	service.ProdukRepository.DeleteProduk(ctx, tx, id)
	return err
}
