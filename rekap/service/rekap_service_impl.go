package service

import (
	"context"
	"database/sql"
	BiayaRespon "es_teh_mewa/biaya/web"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/rekap/model/repository"
	"es_teh_mewa/rekap/web"
	"strconv"
)

type RekapServiceImpl struct {
	RepositoryRekap repository.RekapRepository
	DB              *sql.DB
}

func NewRekapServiceImpl(repo repository.RekapRepository, DB *sql.DB) RekapService {
	return &RekapServiceImpl{
		RepositoryRekap: repo,
		DB:              DB,
	}
}

func (repository *RekapServiceImpl) GetAllPesananToday(ctx context.Context) ([]web.AllPesananRekap, int) {
	tx, err := repository.DB.Begin()
	helperMain.PanicIfError(err)
	data := repository.RepositoryRekap.PesananNonRekapByDate(ctx, tx)
	total := 0
	for _, datum := range data {
		total += datum.TotalPesanan
	}
	return data, total
}

func (repository *RekapServiceImpl) GetAllBiayaToday(ctx context.Context) ([]BiayaRespon.GetBiayaTodayRespon, int) {
	tx, err := repository.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := repository.RepositoryRekap.BiayaNonRekapByDate(ctx, tx)
	var total int
	for _, datum := range data {
		total += datum.Total
	}
	return data, total
}

func (repository *RekapServiceImpl) GetAllBiayaPesananToday(ctx context.Context) []BiayaRespon.GetBiayaTodayRespon {
	tx, err := repository.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	data := repository.RepositoryRekap.BiayaNonRekapByDatePesanan(ctx, tx)
	return data
}

func (repository *RekapServiceImpl) Create(ctx context.Context, keterangan string) (int, error) {
	tx, err := repository.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	id, err := repository.RepositoryRekap.Create(ctx, tx, keterangan)
	return id, err
}

func (repository *RekapServiceImpl) UpdateRekapPesananBiaya(ctx context.Context, id int, id_pesanan []string, id_biaya []string, id_biayaPesanan []string) error {
	tx, err := repository.DB.Begin()
	defer helperMain.ErrorTx(tx)
	helperMain.PanicIfError(err)
	var pesanan []int
	for _, s := range id_pesanan {
		idP, _ := strconv.Atoi(s)
		pesanan = append(pesanan, idP)
	}

	var biaya []int
	for _, s := range id_biaya {
		idP, _ := strconv.Atoi(s)
		biaya = append(biaya, idP)
	}
	var BP []int
	for _, s := range id_biayaPesanan {
		idBP, _ := strconv.Atoi(s)
		BP = append(BP, idBP)
	}
	err = repository.RepositoryRekap.UpdateIdRekapPesananBiaya(ctx, tx, id, pesanan, biaya, BP)
	return err
}
