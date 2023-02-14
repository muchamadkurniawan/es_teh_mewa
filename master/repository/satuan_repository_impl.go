package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type SatuanRepoImpl struct {
	DB *sql.DB
}

func (SatuanRepoImpl) InsertSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan) error {
	//TODO implement me
	panic("implement me")
}

func (SatuanRepoImpl) UpdateSatuan(ctx context.Context, tx *sql.Tx, satuan entity.Satuan) error {
	//TODO implement me
	panic("implement me")
}

func (SatuanRepoImpl) DeleteSatuan(ctx context.Context, tx *sql.Tx, id int32) error {
	//TODO implement me
	panic("implement me")
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

func (SatuanRepoImpl) FindByIdSatuan(ctx context.Context, tx *sql.Tx, id int32) (entity.Satuan, error) {
	satuan := entity.Satuan{}

	return satuan, nil
}
