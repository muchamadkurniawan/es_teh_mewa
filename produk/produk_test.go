package produk

import (
	"context"
	"es_teh_mewa/db"
	"es_teh_mewa/produk/model/entity"
	"es_teh_mewa/produk/model/repository"
	"es_teh_mewa/produk/service"
	"fmt"
	"testing"
)

func TestCreateService(t *testing.T) {
	db := db.GetConnect()
	repo := repository.NewProdukRepo()
	service := service.NewProdukService(repo, db)
	produk := entity.Produk{
		Id:       0,
		Id_User:  1,
		Id_Bahan: 1,
		Harga:    200,
		Stock:    false,
	}
	create, err := service.Create(context.Background(), produk)
	if err != nil {
		panic(err)
	}
	fmt.Println(create)
}

func TestFindAll(t *testing.T) {
	db := db.GetConnect()
	repo := repository.NewProdukRepo()
	service := service.NewProdukService(repo, db)
	all, err := service.FindByAll(context.Background())
	if err != nil {
		return
	}
	fmt.Println(all)
}

func TestFindById(t *testing.T) {
	db := db.GetConnect()
	repo := repository.NewProdukRepo()
	service := service.NewProdukService(repo, db)
	id, err := service.FindById(context.Background(), 1)
	if err != nil {
		return
	}
	fmt.Println(id)
}

func TestUpdate(t *testing.T) {
	db := db.GetConnect()
	repo := repository.NewProdukRepo()
	service := service.NewProdukService(repo, db)
	produk := entity.Produk{
		Id:       2,
		Id_User:  1,
		Id_Bahan: 1,
		Harga:    300,
		Stock:    true,
	}
	update, err := service.Update(context.Background(), produk)
	if err != nil {
		return
	}
	fmt.Println(update)
}

func TestDelete(t *testing.T) {
	db := db.GetConnect()
	repo := repository.NewProdukRepo()
	service := service.NewProdukService(repo, db)
	err := service.Delete(context.Background(), 2)
	if err != nil {
		return
	}
}
