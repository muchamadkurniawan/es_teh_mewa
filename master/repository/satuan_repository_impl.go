package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type SatuanRepoImpl struct {
	DB *sql.DB
}

func NewSatuanRepository() *SatuanRepoImpl {
	return &SatuanRepoImpl{}
}

func (SatuanRepoImpl) InsertSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan) error {
	SQL := "INSERT INTO satuan (nama) VALUES (?)"
	_, err := tx.ExecContext(ctx, SQL, satuan.Nama)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func (SatuanRepoImpl) UpdateSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan) error {
	SQL := "UPDATE satuan SET nama = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, satuan.Nama, satuan.Id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (SatuanRepoImpl) DeleteSatuan(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "DELETE FROM satuan WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return err
}

func (SatuanRepoImpl) FindAllSatuan(ctx context.Context, tx *sql.Tx) ([]entity.Satuan, error) {
	SQL := "SELECT id, nama FROM satuan"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var satuans []entity.Satuan
	for rows.Next() {
		satuan := entity.Satuan{}
		err := rows.Scan(&satuan.Id, &satuan.Nama)
		if err != nil {
			panic(err)
			return nil, err
		}
		satuans = append(satuans, satuan)
	}
	return satuans, nil
}

func (SatuanRepoImpl) FindByIdSatuan(ctx context.Context, tx *sql.Tx, id int) (entity.Satuan, error) {
	satuan := entity.Satuan{}
	SQL := "SELECT id, nama FROM satuan where id=?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&satuan.Id, &satuan.Nama)
		if err != nil {
			panic(err)
		}
	}
	return satuan, nil
}
