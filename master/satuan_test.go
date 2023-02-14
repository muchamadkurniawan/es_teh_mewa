package master

import (
	"context"
	"eh_teh_mewa/db"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/service"
	"eh_teh_mewa/master/web"
	"fmt"
	"testing"
)

func TestReadAllService(t *testing.T) {
	db := db.GetConnect()
	satuanRepo := repository.NewSatuanRepository()
	satuanService := service.NewSatuanService(satuanRepo, db)
	all := satuanService.FindAll(context.Background())
	fmt.Println(all)
}
func TestReadByIdService(t *testing.T) {
	db := db.GetConnect()
	satuanRepo := repository.NewSatuanRepository()
	satuanService := service.NewSatuanService(satuanRepo, db)
	all := satuanService.FindById(context.Background(), 1)
	fmt.Println(all)
}

func TestCreateService(t *testing.T) {
	db := db.GetConnect()
	satuanRepo := repository.NewSatuanRepository()
	satuanService := service.NewSatuanService(satuanRepo, db)
	satuan := web.SatuanRequest{Nama: "kaleng"}
	satuanService.Save(context.Background(), satuan)
}

func TestUpdateService(t *testing.T) {
	db := db.GetConnect()
	satuanRepo := repository.NewSatuanRepository()
	satuanService := service.NewSatuanService(satuanRepo, db)
	satuanResponce := web.SatuanResponse{
		Id:   2,
		Name: "sachet",
	}
	satuanService.Update(context.Background(), satuanResponce)
}

func TestDeleteService(t *testing.T) {
	db := db.GetConnect()
	satuanRepo := repository.NewSatuanRepository()
	satuanService := service.NewSatuanService(satuanRepo, db)
	satuanService.Delete(context.Background(), 2)
}
