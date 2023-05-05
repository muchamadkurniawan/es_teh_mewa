package master

import (
	"context"
	"es_teh_mewa/db"
	"es_teh_mewa/master/repository"
	"es_teh_mewa/master/service"
	"es_teh_mewa/master/web"
	"fmt"
	"testing"
)

func TestSaveService(t *testing.T) {
	db := db.GetConnect()
	bakuRepository := repository.NewBahanRepository()
	bahanbakuService := service.NewBahanbakuService(bakuRepository, db)
	request := web.BahanbakuRequest{
		IdSatuan: 1,
		Nama:     "coba bahan baku",
	}
	bahanbakuService.Save(context.Background(), request)
}

func TestFindAll(t *testing.T) {
	db := db.GetConnect()
	bakuRepository := repository.NewBahanRepository()
	bahanbakuService := service.NewBahanbakuService(bakuRepository, db)
	all := bahanbakuService.FindAll(context.Background())
	fmt.Println(all)
}

func TestFindById(t *testing.T) {
	db := db.GetConnect()
	bakuRepository := repository.NewBahanRepository()
	bahanbakuService := service.NewBahanbakuService(bakuRepository, db)
	bahan := bahanbakuService.FindById(context.Background(), 1)
	fmt.Println(bahan)
}

func TestUpdateSerive(t *testing.T) {
	db := db.GetConnect()
	bakuRepository := repository.NewBahanRepository()
	bahanbakuService := service.NewBahanbakuService(bakuRepository, db)
	responses := web.BahanbakuResponse{
		Id:       1,
		IdSatuan: 1,
		Nama:     "update bahan baku",
	}
	bahanbakuService.Update(context.Background(), responses)
}

func TestDeleteSerive(t *testing.T) {
	db := db.GetConnect()
	bakuRepository := repository.NewBahanRepository()
	bahanbakuService := service.NewBahanbakuService(bakuRepository, db)
	bahanbakuService.Delete(context.Background(), 2)
}
