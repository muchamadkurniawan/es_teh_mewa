package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/biaya/web"
	"es_teh_mewa/helperMain"
	"time"
)

type BiayaRepositoryImpl struct{}

func NewBiayaRepository() BiayaRepository {
	return &BiayaRepositoryImpl{}
}

func (repo *BiayaRepositoryImpl) GetBahanBakuNonProdukRepo(ctx context.Context, tx *sql.Tx) []web.GetBahanBakuNonProdukRespon {
	SQL := "select bahan_baku.id as id, satuan.nama as satuan, bahan_baku.nama as nama from bahan_baku " +
		"left join satuan on bahan_baku.id_satuan = satuan.id " +
		"left join produk_jual on produk_jual.id_bahan_baku = bahan_baku.id " +
		"where produk_jual.id_bahan_baku is null;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var bahans []web.GetBahanBakuNonProdukRespon
	for row.Next() {
		bahan := web.GetBahanBakuNonProdukRespon{}
		row.Scan(&bahan.Id, &bahan.Satuan, &bahan.Nama)
		bahans = append(bahans, bahan)
	}
	return bahans
}

func (repo *BiayaRepositoryImpl) GetBiayaTodayRepo(ctx context.Context, tx *sql.Tx) []web.GetBiayaTodayRespon {
	currentTime := time.Now()
	SQL := "SELECT detail_biaya.id AS id, concat(bahan_baku.nama, '-', satuan.nama) AS bahan, detail_biaya.jumlah AS jumlah, " +
		"detail_biaya.harga_satuan AS harga_satuan, detail_biaya.total AS total FROM detail_biaya " +
		"LEFT JOIN bahan_baku ON bahan_baku.id = detail_biaya.id_bahan_baku " +
		"LEFT JOIN satuan ON bahan_baku.id_satuan = satuan.id " +
		"WHERE CAST(detail_biaya.tanggal AS DATE) = ? AND detail_biaya.id_detail_pesanan IS NULL order by detail_biaya.tanggal DESC;"
	row, err := tx.QueryContext(ctx, SQL, currentTime.Format("2006-01-02"))
	helperMain.PanicIfError(err)
	var biayas []web.GetBiayaTodayRespon
	for row.Next() {
		biaya := web.GetBiayaTodayRespon{}
		row.Scan(&biaya.Id, &biaya.Barang, &biaya.Jumlah, &biaya.HargaSatuan, &biaya.Total)
		biayas = append(biayas, biaya)
	}

	return biayas
}

func (repo *BiayaRepositoryImpl) CreateBiaya(ctx context.Context, tx *sql.Tx, request web.BiayaRequestCreate) error {
	currentTime := time.Now()
	SQL := "insert into detail_biaya(id_bahan_baku, jumlah, harga_satuan, total, tanggal) VALUES(?, ?, ?, ?, ?);"
	_, err := tx.ExecContext(ctx, SQL, request.Barang, request.Jumlah, request.HargaSatuan, request.Total, currentTime.Format("2006-01-02 15:04:05"))
	return err
}

func (repo *BiayaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) web.GetBiayaRespon {
	SQL := "select id, id_bahan_baku, jumlah, harga_satuan, total, tanggal from detail_biaya " +
		"where id = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var data web.GetBiayaRespon
	if row.Next() {
		row.Scan(&data.Id, &data.Barang, &data.Jumlah, &data.HargaSatuan, &data.Total, &data.Tanggal)
	}
	return data
}

func (repo *BiayaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "delete from detail_biaya where id = ?;"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	return nil
}
