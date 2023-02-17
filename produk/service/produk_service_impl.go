package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/produk/model/repository"
	"eh_teh_mewa/produk/model/web"
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

func (service *ProdukServiceImpl) FindByAll(ctx context.Context) ([]web.ResponseProduk, error) {
	//tx, err := service.DB.Begin()
	//helperMain.PanicIfError(err)
	//defer helperMain.ErrorTx(tx)
	//produk, err := service.ProdukRepository.FindAllProduk(ctx, tx)
	//helperMain.PanicIfError(err)
	//return produk, nil
	//TODO implement me
	panic("implement me")
}

func (service *ProdukServiceImpl) FindById(ctx context.Context, id int) (web.ResponseProduk, error) {
	//tx, err := service.DB.Begin()
	//helperMain.PanicIfError(err)
	//defer helperMain.ErrorTx(tx)
	//produk, err := service.ProdukRepository.FindProdukById(ctx, tx, id)
	//return produk, nil
	//TODO implement me
	panic("implement me")
}

func (service *ProdukServiceImpl) Create(ctx context.Context, produk web.RequestProduk) (web.ResponseProduk, error) {
	//TODO implement me
	panic("implement me")
}

func (service *ProdukServiceImpl) Update(ctx context.Context, produk web.ResponseProduk) (web.ResponseProduk, error) {
	//TODO implement me
	panic("implement me")
}

func (service *ProdukServiceImpl) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
