package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/produk/model/web"
	"errors"
)

type PembelianRepositoryImpl struct{}

func NewProdukRepository() ProdukRepository {
	return &PembelianRepositoryImpl{}
}

func (repository *PembelianRepositoryImpl) FindAllProduk(ctx context.Context, tx *sql.Tx) ([]web.ResponseProduk, error) {
	var respon []web.ResponseProduk
	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual;"
	rows, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	for rows.Next() {
		newRespon := web.ResponseProduk{}
		rows.Scan(&newRespon.Id, &newRespon.Id_User, &newRespon.Id_Barang, &newRespon.Harga, &newRespon.Stock)
		respon = append(respon, newRespon)
	}
	return respon, nil
}

func (repository *PembelianRepositoryImpl) FindProdukById(ctx context.Context, tx *sql.Tx, id int) (web.ResponseProduk, error) {
	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual WHERE id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	respon := web.ResponseProduk{}
	if rows.Next() {
		rows.Scan(&respon.Id, &respon.Id_User, &respon.Id_Barang, &respon.Harga, &respon.Stock)
		return respon, nil
	} else {
		return respon, errors.New("It's Not Found")
	}
}

func (repository *PembelianRepositoryImpl) CreateProduk(ctx context.Context, tx *sql.Tx, produk web.RequestProduk) (web.RequestProduk, error) {
	SQL := "INSERT INTO produk_jual(id_user, id_bahan_baku, harga, stock) VALUES(?, ?, ?, ?);"
	result, err := tx.ExecContext(ctx, SQL, produk.Id_User, produk.Id_Barang, produk.Harga, produk.Stock)
	helperMain.PanicIfError(err)
	id, err := result.LastInsertId()
	helperMain.PanicIfError(err)
	produk.Id = int(id)
	return produk, nil
}

func (repository *PembelianRepositoryImpl) UpdateProduk(ctx context.Context, tx *sql.Tx, produk web.RequestProduk) (web.RequestProduk, error) {
	SQL := "UPDATE produk_jual SET id_user = ?, id_bahan_baku = ?, harga = ?, stock = ? WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, produk.Id_User, produk.Id_Barang, produk.Harga, produk.Stock)
	helperMain.PanicIfError(err)
	return produk, nil
}

func (repository *PembelianRepositoryImpl) DeleteProduk(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "DELETE FROM produk_jual WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	return nil
}
