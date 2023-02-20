package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/master/model/entity"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/web"
)

type BahanbakuServiceimpl struct {
	BahanRepo repository.BahanBakuRepository
	DB        *sql.DB
}

func NewBahanbakuService(bakuRepository repository.BahanBakuRepository, db *sql.DB) BahanbakuService {
	return &BahanbakuServiceimpl{
		BahanRepo: bakuRepository,
		DB:        db,
	}
}

func (service *BahanbakuServiceimpl) GetSatuan(ctx context.Context) []web.SatuanResponse {
	tx, err := service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	satuan := repository.NewSatuanRepository()
	satuans, err := satuan.FindAllSatuan(ctx, tx)
	return helperMain.ToSatuanResponses(satuans)
}

func (Service *BahanbakuServiceimpl) Save(ctx context.Context, request web.BahanbakuRequest) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()
	bahan := entity.BahanBaku{
		IdSatuan: request.IdSatuan,
		Nama:     request.Nama,
	}
	Service.BahanRepo.Insert(ctx, tx, bahan)
}

func (Service *BahanbakuServiceimpl) Update(ctx context.Context, response web.BahanbakuRequest) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	bahan := entity.BahanBaku{
		Id:       response.Id,
		IdSatuan: response.IdSatuan,
		Nama:     response.Nama,
	}
	Service.BahanRepo.Update(ctx, tx, bahan)
}

func (Service *BahanbakuServiceimpl) Delete(ctx context.Context, id int) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	Service.BahanRepo.Delete(ctx, tx, id)
}

func (Service *BahanbakuServiceimpl) FindAll(ctx context.Context) []web.BahanbakuFullResponse {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	all := Service.BahanRepo.FindAll(ctx, tx)

	return helperMain.ToBahanRresponses(all)
}

func (Service *BahanbakuServiceimpl) FindById(ctx context.Context, id int) web.BahanbakuResponse {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)
	byId := Service.BahanRepo.FindById(ctx, tx, id)

	return web.BahanbakuResponse{
		Id:       byId.Id,
		Nama:     byId.Nama,
		IdSatuan: byId.IdSatuan,
	}
}
