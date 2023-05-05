package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/rekap/web"
	"time"
)

type RekapRepositoryImpl struct{}

func NewRekapRepository() RekapRepository {
	return &RekapRepositoryImpl{}
}

func (r *RekapRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (web.RekapResponseTime, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RekapRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]web.RekapResponseTime, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RekapRepositoryImpl) PesananNonRekapByDate(ctx context.Context, tx *sql.Tx) []web.AllPesananRekap {
	currentTime := time.Now()
	SQL := "select pesanan.id, pesanan.tanggal, pesanan.pembayaran, " +
		"(select sum(detail_pesanan.total) from detail_pesanan where detail_pesanan.id_pesanan = pesanan.id) as total " +
		"from pesanan WHERE CAST(pesanan.tanggal AS DATE) = ? AND pesanan.id_rekap is null;"
	row, err := tx.QueryContext(ctx, SQL, currentTime.Format("2006-01-02"))
	helperMain.PanicIfError(err)
	var pesanans []web.AllPesananRekap
	for row.Next() {
		pesanan := web.AllPesananRekap{}
		row.Scan(&pesanan.Id, &pesanan.Tanggal, &pesanan.Pembayaran, &pesanan.TotalPesanan)
		pesanans = append(pesanans, pesanan)
	}
	return pesanans
}

func (r *RekapRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error {
	//TODO implement me
	panic("implement me")
}

func (r *RekapRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error {
	//TODO implement me
	panic("implement me")
}

func (r *RekapRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	//TODO implement me
	panic("implement me")
}
