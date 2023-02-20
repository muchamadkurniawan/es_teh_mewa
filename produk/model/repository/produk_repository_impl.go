package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/produk/model/entity"
	"eh_teh_mewa/produk/model/web"
)

type ProdukRepositoryImpl struct {
	DB *sql.DB
}

func NewProdukRepo() *ProdukRepositoryImpl {
	return &ProdukRepositoryImpl{}
}

func (ProdukRepositoryImpl) FindAllProduk(ctx context.Context, tx *sql.Tx) []web.ResponseProdukFull {
	SQL := "SELECT bahan_baku.id, users.username, bahan_baku.nama, bahan_baku.harga, bahan_baku.stock FROM produk_jual JOIN users ON user.id = produk_jual.id_user " +
		"JOIN bahan_baku ON bahan_baku.id = produk_jual.id_bahan_baku;"
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
	SQL := "INSERT INTO produk_jual (id_user, id_bahan_baku, harga, stock) VALUES (?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, produk.Id_User, produk.Id_Bahan, produk.Harga, produk.Stock)
	if err != nil {
		panic(err)
	}
}

func (ProdukRepositoryImpl) UpdateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk) entity.Produk {
	SQL := "UPDATE produk_jual SET id_user = ?, id_bahan_baku = ?, harga = ?, stock = ? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, produk.Id_User, produk.Id_Bahan, produk.Harga, produk.Stock, produk.Id)
	if err != nil {
		panic(err)
		return entity.Produk{}
	}
	return produk
}

func (ProdukRepositoryImpl) DeleteProduk(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM produk_jual WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}
