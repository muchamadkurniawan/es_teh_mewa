package repository

import (
	"context"
	"database/sql"
	BiayaRespon "es_teh_mewa/biaya/web"
	"es_teh_mewa/dashboard/web"
	"es_teh_mewa/helperMain"
	web3 "es_teh_mewa/rekap/web"
)

type DashboardRepositoryImpl struct{}

func NewDashboardRepository() DashboardRepository {
	return &DashboardRepositoryImpl{}
}

func (d DashboardRepositoryImpl) GetRekap(ctx context.Context, tx *sql.Tx) []web.ResponseDashboard {
	SQL := "select id, tanggal, keterangan from rekap order by tanggal DESC;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var respons []web.ResponseDashboard
	for row.Next() {
		respon := web.ResponseDashboard{}
		row.Scan(&respon.Id, &respon.Tanggal, &respon.Keterangan)
		respons = append(respons, respon)
	}
	return respons
}

func (d DashboardRepositoryImpl) GetRekapById(ctx context.Context, tx *sql.Tx, id int) web.ResponseDashboard {
	SQL := "select id, tanggal, keterangan from rekap where id = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	respon := web.ResponseDashboard{}
	if row.Next() {
		row.Scan(&respon.Id, &respon.Tanggal, &respon.Keterangan)
	}
	return respon
}

func (d DashboardRepositoryImpl) GetBiayaRekapById(ctx context.Context, tx *sql.Tx, id int) []BiayaRespon.GetBiayaTodayRespon {
	SQL := "SELECT detail_biaya.id AS id, concat(bahan_baku.nama, '-', satuan.nama) AS bahan, detail_biaya.jumlah AS jumlah, " +
		"detail_biaya.harga_satuan AS harga_satuan, detail_biaya.total AS total FROM detail_biaya " +
		"LEFT JOIN bahan_baku ON bahan_baku.id = detail_biaya.id_bahan_baku " +
		"LEFT JOIN satuan ON bahan_baku.id_satuan = satuan.id " +
		"WHERE detail_biaya.id_rekap = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var biayas []BiayaRespon.GetBiayaTodayRespon
	for row.Next() {
		biaya := BiayaRespon.GetBiayaTodayRespon{}
		row.Scan(&biaya.Id, &biaya.Barang, &biaya.Jumlah, &biaya.HargaSatuan, &biaya.Total)
		biayas = append(biayas, biaya)
	}
	return biayas
}

func (d DashboardRepositoryImpl) GetPesananRekapById(ctx context.Context, tx *sql.Tx, id int) []web3.AllPesananRekap {
	SQL := "select pesanan.id, pesanan.tanggal, pesanan.pembayaran, " +
		"(select sum(detail_pesanan.total) from detail_pesanan where detail_pesanan.id_pesanan = pesanan.id) as total " +
		"from pesanan WHERE pesanan.id_rekap = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var pesanans []web3.AllPesananRekap
	for row.Next() {
		pesanan := web3.AllPesananRekap{}
		row.Scan(&pesanan.Id, &pesanan.Tanggal, &pesanan.Pembayaran, &pesanan.TotalPesanan)
		pesanans = append(pesanans, pesanan)
	}
	return pesanans
}
