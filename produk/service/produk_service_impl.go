package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/produk/model/entity"
	"eh_teh_mewa/produk/model/repository"
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

func (service *ProdukServiceImpl) FindByAll(ctx context.Context) ([]entity.Produk, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	var produks []entity.Produk
	produks = service.ProdukRepository.FindAllProduk(ctx, tx)
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

func (service *ProdukServiceImpl) Create(ctx context.Context, produk entity.Produk) (entity.Produk, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	//defer helperMain.ErrorTx(tx)
	defer tx.Commit()
	service.ProdukRepository.CreateProduk(ctx, tx, produk)
	return produk, err
}

func (service *ProdukServiceImpl) Update(ctx context.Context, produk entity.Produk) (entity.Produk, error) {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	updateProduk := service.ProdukRepository.UpdateProduk(ctx, tx, produk)
	return updateProduk, err
}

func (service *ProdukServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	service.ProdukRepository.DeleteProduk(ctx, tx, id)
	return err
}
