package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/master/model/entity"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/web"
)

type SatuanServiceImpl struct {
	SatuanRepo repository.SatuanRepository
	DB         *sql.DB
}

func NewSatuanService(satuanRepository repository.SatuanRepository, DB *sql.DB) SatuanService {
	return &SatuanServiceImpl{
		SatuanRepo: satuanRepository,
		DB:         DB,
	}
}

func (Service *SatuanServiceImpl) Save(ctx context.Context, request web.SatuanRequest) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)
	satuan := entity.Satuan{
		Nama: request.Nama,
	}
	Service.SatuanRepo.InsertSatuan(ctx, tx, satuan)
}

func (Service *SatuanServiceImpl) Update(ctx context.Context, response web.SatuanResponse) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	satuan := entity.Satuan{
		Id:   response.Id,
		Nama: response.Name,
	}
	Service.SatuanRepo.UpdateSatuan(ctx, tx, satuan)
}

func (Service *SatuanServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	Service.SatuanRepo.DeleteSatuan(ctx, tx, id)
}

func (Service *SatuanServiceImpl) FindAll(ctx context.Context) []web.SatuanResponse {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	satuans, err := Service.SatuanRepo.FindAllSatuan(ctx, tx)
	if err != nil {
		return nil
	}
	satuanResponses := helperMain.ToSatuanResponses(satuans)
	return satuanResponses
}

func (Service *SatuanServiceImpl) FindById(ctx context.Context, id int) web.SatuanResponse {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	satuan, err := Service.SatuanRepo.FindByIdSatuan(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	satuanResponse := helperMain.ToSatuanResponse(satuan)
	return satuanResponse
}
