package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/produk/model/entity"
)

type ProdukRepositoryImpl struct {
	DB *sql.DB
}

func NewProdukRepo() *ProdukRepositoryImpl {
	return &ProdukRepositoryImpl{}
}

func (ProdukRepositoryImpl) FindAllProduk(ctx context.Context, tx *sql.Tx) []entity.Produk {
	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual;"
	rows, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var produks []entity.Produk
	if rows.Next() {
		produk := entity.Produk{}
		err := rows.Scan(&produk.Id, &produk.Id_User, &produk.Id_Bahan, &produk.Harga, &produk.Stock)
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

//type PembelianRepositoryImpl struct{}
//
//func NewProdukRepository() ProdukRepository {
//	return &PembelianRepositoryImpl{}
//}
//
//func (repository *PembelianRepositoryImpl) FindAllProduk(ctx context.Context, tx *sql.Tx) ([]web.ResponseProduk, error) {
//	var respon []web.ResponseProduk
//	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual;"
//	rows, err := tx.QueryContext(ctx, SQL)
//	helperMain.PanicIfError(err)
//	for rows.Next() {
//		newRespon := web.ResponseProduk{}
//		rows.Scan(&newRespon.Id, &newRespon.Id_User, &newRespon.Id_Barang, &newRespon.Harga, &newRespon.Stock)
//		respon = append(respon, newRespon)
//	}
//	return respon, nil
//}
//
//func (repository *PembelianRepositoryImpl) FindAllProdukById(ctx context.Context, tx *sql.Tx, id int) (web.ResponseProduk, error) {
//	SQL := "SELECT id, id_user, id_bahan_baku, harga, stock FROM produk_jual WHERE id = ?;"
//	rows, err := tx.QueryContext(ctx, SQL, id)
//	helperMain.PanicIfError(err)
//	respon := web.ResponseProduk{}
//	if rows.Next() {
//		rows.Scan(&respon.Id, &respon.Id_User, &respon.Id_Barang, &respon.Harga, &respon.Stock)
//		return respon, nil
//	} else {
//		return respon, errors.New("It's Not Found")
//	}
//}
//
//func (repository *PembelianRepositoryImpl) CreateProduk(ctx context.Context, tx *sql.Tx, produk web.RequestProduk) (web.RequestProduk, error) {
//	//SQL := "INSERT "
//	panic("implement me")
//}
//
//func (repository *PembelianRepositoryImpl) UpdateProduk(ctx context.Context, tx *sql.Tx, produk web.RequestProduk) (web.RequestProduk, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (repository *PembelianRepositoryImpl) DeleteProduk(ctx context.Context, tx *sql.Tx, id int) error {
//	//TODO implement me
//	panic("implement me")
//}
