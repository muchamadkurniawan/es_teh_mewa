package repository

import (
	"context"
	"database/sql"
	BiayaRespon "es_teh_mewa/biaya/web"
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

func (r *RekapRepositoryImpl) BiayaNonRekapByDate(ctx context.Context, tx *sql.Tx) []BiayaRespon.GetBiayaTodayRespon {
	currentTime := time.Now()
	SQL := "SELECT detail_biaya.id AS id, concat(bahan_baku.nama, '-', satuan.nama) AS bahan, detail_biaya.jumlah AS jumlah, " +
		"detail_biaya.harga_satuan AS harga_satuan, detail_biaya.total AS total FROM detail_biaya " +
		"LEFT JOIN bahan_baku ON bahan_baku.id = detail_biaya.id_bahan_baku " +
		"LEFT JOIN satuan ON bahan_baku.id_satuan = satuan.id " +
		"WHERE CAST(detail_biaya.tanggal AS DATE) = ? AND detail_biaya.id_detail_pesanan IS NULL;"
	row, err := tx.QueryContext(ctx, SQL, currentTime.Format("2006-01-02"))
	helperMain.PanicIfError(err)
	var biayas []BiayaRespon.GetBiayaTodayRespon
	for row.Next() {
		biaya := BiayaRespon.GetBiayaTodayRespon{}
		row.Scan(&biaya.Id, &biaya.Barang, &biaya.Jumlah, &biaya.HargaSatuan, &biaya.Total)
		biayas = append(biayas, biaya)
	}
	return biayas
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
