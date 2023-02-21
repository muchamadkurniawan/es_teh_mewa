package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/pembelian/model/entity"
	"eh_teh_mewa/pembelian/model/web"
	"errors"
	"time"
)

type PembelianRespositoryImpl struct {
}

func NewPembelianRepository() PembelianRepository {
	return &PembelianRespositoryImpl{}
}

func (repository *PembelianRespositoryImpl) UpdatePembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error) {
	SQL := "UPDATE pembelian SET id_user = ?, id_bahan_baku = ?, tanggal = ?, jumlah=?, biaya = ?, use_pembelian = ? WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, pembelian.Id_user, pembelian.Id_bahan_baku, pembelian.Tanggal, pembelian.Jumlah, pembelian.Biaya, pembelian.Use_pembelian, pembelian.Id)
	if err != nil {
		return pembelian, err
	}
	return pembelian, nil
}

func (repository *PembelianRespositoryImpl) DeletePembelian(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM pembelian WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (repository *PembelianRespositoryImpl) InsertPembelian(ctx context.Context, tx *sql.Tx, pembelian web.PembelianCreateRequest) (web.PembelianCreateRequest, error) {
	SQL := "INSERT INTO pembelian(id_user, id_bahan_baku, tanggal, jumlah, biaya, use_pembelian) VALUES(?, ?, ?, ?, ?, ?);"
	result, err := tx.ExecContext(ctx, SQL, pembelian.Id_user, pembelian.Id_bahan_baku, pembelian.Tanggal, pembelian.Jumlah, pembelian.Biaya, pembelian.Use_pembelian)
	helperMain.PanicIfError(err)

	id, err := result.LastInsertId()
	helperMain.PanicIfError(err)
	pembelian.Id = int(id)
	return pembelian, nil
}

func (repository *PembelianRespositoryImpl) FindByIdPembelian(ctx context.Context, tx *sql.Tx, id string) (entity.Pembelian, error) {
	var pembelian entity.Pembelian
	SQL := "SELECT id, id_user, id_bahan_baku, tanggal, jumlah, biaya, use_pembelian FROM pembelian WHERE id = ? LIMIT 1;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	// defer rows.Close()
	if err != nil {
		return pembelian, err
	}
	if rows.Next() {
		err := rows.Scan(&pembelian.Id, &pembelian.IdUser, &pembelian.IdBahan_Baku, &pembelian.Tanggal, &pembelian.Jumlah, &pembelian.Biaya, &pembelian.UsePembelian)
		if err != nil {
			return pembelian, err
		}
		return pembelian, nil
	} else {
		return pembelian, errors.New("id Not Found")
	}
}
func (repository *PembelianRespositoryImpl) FindByAllPembelian(ctx context.Context, tx *sql.Tx) ([]web.PembelianResponseFull, error) {
	var pembelian []web.PembelianResponseFull
	date := time.Now().AddDate(0, 0, 1).Format("2006-01-22")
	now := time.Now().Format("2006-01-21")
	SQL := "SELECT pembelian.id, pembelian.id_user, bahan_baku.nama, pembelian.tanggal, pembelian.jumlah, " +
		"pembelian.biaya, pembelian.use_pembelian FROM pembelian INNER JOIN bahan_baku ON bahan_baku.id = pembelian.id_bahan_baku WHERE pembelian.tanggal BETWEEN ? AND ?;"
	rows, err := tx.QueryContext(ctx, SQL, date, now)
	helperMain.PanicIfError(err)
	for rows.Next() {
		newPembelian := web.PembelianResponseFull{}
		err := rows.Scan(&newPembelian.Id, &newPembelian.Id_user, &newPembelian.Id_bahan_baku, &newPembelian.Tanggal, &newPembelian.Jumlah, &newPembelian.Biaya, &newPembelian.Use_pembelian)
		helperMain.PanicIfError(err)

		pembelian = append(pembelian, newPembelian)
	}
	return pembelian, nil
}

func (repository *PembelianRespositoryImpl) FindByAllPembelianByDate(ctx context.Context, tx *sql.Tx, filterAwal string, filterAkhir string) ([]web.PembelianResponseFull, error) {
	var pembelian []web.PembelianResponseFull
	SQL := "SELECT pembelian.id, pembelian.id_user, bahan_baku.nama, pembelian.tanggal, pembelian.jumlah, pembelian.biaya, " +
		"pembelian.use_pembelian FROM pembelian INNER JOIN bahan_baku ON bahan_baku.id = pembelian.id_bahan_baku WHERE pembelian.tanggal BETWEEN ? AND ?;"
	rows, err := tx.QueryContext(ctx, SQL, filterAwal, filterAkhir)
	// defer rows.Close()
	if err != nil {
		return pembelian, err
	}
	for rows.Next() {
		newPembelian := web.PembelianResponseFull{}
		err := rows.Scan(&newPembelian.Id, &newPembelian.Id_user, &newPembelian.Id_bahan_baku, &newPembelian.Tanggal, &newPembelian.Jumlah, &newPembelian.Biaya, &newPembelian.Use_pembelian)
		if err != nil {
			return pembelian, err
		}

		pembelian = append(pembelian, newPembelian)
	}
	return pembelian, nil
}
