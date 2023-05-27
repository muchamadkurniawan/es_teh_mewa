package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/stok/web"
)

type StokRepositoryImpl struct{}

func NewStokRepository() StokRepository {
	return StokRepositoryImpl{}
}

func (s StokRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id_bahan int, total int) error {
	SQL := "update stock set total = ? where id_bahan = ?;"
	_, err := tx.ExecContext(ctx, SQL, total, id_bahan)
	return err
}

func (s StokRepositoryImpl) GetByAll(ctx context.Context, tx *sql.Tx) []web.ResponseStok {
	SQL := "select bahan_baku.nama, stok.total from stok inner join bahan_baku on bahan_baku.id = stok.id_bahan_baku;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var stoks []web.ResponseStok
	for row.Next() {
		stok := web.ResponseStok{}
		row.Scan(&stok.Bahan, &stok.Total)
		stoks = append(stoks, stok)
	}
	return stoks
}

func (s StokRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) web.ResponseStok {
	SQL := "select bahan_baku.nama, stok.total from stok " +
		"inner join bahan_baku on bahan_baku.id = stok.id_bahan_baku " +
		"where stok.id = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var stoks web.ResponseStok
	if row.Next() {
		row.Scan(&stoks.Bahan, &stoks.Total)
	}
	return stoks
}

func (s StokRepositoryImpl) GetByAllDetail(ctx context.Context, tx *sql.Tx) []web.ResponseDetailStok {
	SQL := "select bahan_baku.nama, detail_stok.type, detail_stok.jumlah from detail_stok " +
		"inner join bahan_baku on bahan_baku.id = detail_stok.id_bahan_baku;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var stoks []web.ResponseDetailStok
	for row.Next() {
		stok := web.ResponseDetailStok{}
		row.Scan(&stok.Bahan, &stok.Type, &stok.Total)
		stoks = append(stoks, stok)
	}
	return stoks
}
