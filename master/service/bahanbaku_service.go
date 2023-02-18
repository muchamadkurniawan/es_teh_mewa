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

func (Service *BahanbakuServiceimpl) Update(ctx context.Context, response web.BahanbakuResponse) {
	tx, err := Service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

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

func (Service *BahanbakuServiceimpl) FindAll(ctx context.Context) []web.BahanbakuResponse {
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

	return helperMain.ToBahanResponse(byId)
}
