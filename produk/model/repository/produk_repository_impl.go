package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/produk/model/entity"
	"es_teh_mewa/produk/model/web"
)

type ProdukRepositoryImpl struct {
}

func NewProdukRepo() ProdukRepository {
	return &ProdukRepositoryImpl{}
}

func (ProdukRepositoryImpl) FindAllProduk(ctx context.Context, tx *sql.Tx) []web.ResponseProdukFull {
	SQL := "SELECT produk_jual.id, users.username, bahan_baku.nama, produk_jual.harga, produk_jual.stock FROM produk_jual INNER JOIN users ON users.id = produk_jual.id_user " +
		"INNER JOIN bahan_baku ON bahan_baku.id = produk_jual.id_bahan_baku;"
	rows, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var produks []web.ResponseProdukFull
	for rows.Next() {
		produk := web.ResponseProdukFull{}
		err := rows.Scan(&produk.Id, &produk.User, &produk.Barang, &produk.Harga, &produk.Stock)
		if err != nil {
			helperMain.PanicIfError(err)
		}
		produks = append(produks, produk)
	}
	return produks
}

func (ProdukRepositoryImpl) FindAllProdukByBahan(ctx context.Context, tx *sql.Tx, barang int) []web.ResponseProdukFull {
	SQL := "SELECT produk_jual.id, users.username, bahan_baku.nama, produk_jual.harga, produk_jual.stock FROM produk_jual INNER JOIN users ON users.id = produk_jual.id_user " +
		"INNER JOIN bahan_baku ON bahan_baku.id = produk_jual.id_bahan_baku WHERE bahan_baku.id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, barang)
	helperMain.PanicIfError(err)
	var produks []web.ResponseProdukFull
	for rows.Next() {
		produk := web.ResponseProdukFull{}
		err := rows.Scan(&produk.Id, &produk.User, &produk.Barang, &produk.Harga, &produk.Stock)
		if err != nil {
			helperMain.PanicIfError(err)
		}
		produks = append(produks, produk)
	}
	return produks
}

func (ProdukRepositoryImpl) FindProdukByBahan(ctx context.Context, tx *sql.Tx, bahan int) (entity.Produk, error) {
	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual WHERE id_bahan_baku=?;"
	rows, err := tx.QueryContext(ctx, SQL, bahan)
	helperMain.PanicIfError(err)
	produk := entity.Produk{}
	if rows.Next() {
		err := rows.Scan(&produk.Id, &produk.Id_User, &produk.Id_Bahan, &produk.Harga, &produk.Stock)
		helperMain.PanicIfError(err)
	}
	return produk, nil
}

func (ProdukRepositoryImpl) FindProdukById(ctx context.Context, tx *sql.Tx, id int) entity.Produk {
	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual WHERE id=?;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	produk := entity.Produk{}
	if rows.Next() {
		err := rows.Scan(&produk.Id, &produk.Id_User, &produk.Id_Bahan, &produk.Harga, &produk.Stock)
		helperMain.PanicIfError(err)
	}
	return produk
}

func (ProdukRepositoryImpl) CreateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk) {
	SQL := "INSERT INTO produk_jual(id_user, id_bahan_baku, harga, stock) VALUES(?,?,?,?);"
	save, err := tx.ExecContext(ctx, SQL, produk.Id_User, produk.Id_Bahan, produk.Harga, produk.Stock)
	helperMain.PanicIfError(err)
	id, err := save.LastInsertId()
	produk.Id = int(id)
}

func (ProdukRepositoryImpl) UpdateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk) {
	SQL := "UPDATE produk_jual SET id_user = ?, id_bahan_baku = ?, harga = ?, stock = ? WHERE id=?;"
	_, err := tx.ExecContext(ctx, SQL, produk.Id_User, produk.Id_Bahan, produk.Harga, produk.Stock, produk.Id)
	helperMain.PanicIfError(err)
}

func (ProdukRepositoryImpl) DeleteProduk(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM produk_jual WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}
