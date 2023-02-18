package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type BahanRepository struct {
	DB *sql.DB
}

func NewBahanRepository() *BahanRepository {
	return &BahanRepository{}
}

func (BahanRepository) Insert(ctx context.Context, tx *sql.Tx, baku entity.BahanBaku) {
	SQL := "INSERT INTO bahan_baku (id_Satuan, nama) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, baku.IdSatuan, baku.Nama)
	if err != nil {
		panic(err)
	}
}

func (BahanRepository) Update(ctx context.Context, tx *sql.Tx, baku entity.BahanBaku) entity.BahanBaku {
	SQL := "UPDATE bahan_baku SET id_satuan = ?, nama = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, baku.IdSatuan, baku.Nama, baku.Id)
	if err != nil {
		panic(err)
		return entity.BahanBaku{}
	}
	return baku
}

func (BahanRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM bahan_baku WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}

func (BahanRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.BahanBakuFull {
	SQL := "SELECT bahan_baku.id, satuan.nama, bahan_baku.nama FROM bahan_baku INNER JOIN satuan ON bahan_baku.id_satuan = satuan.id;"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	var bahans []entity.BahanBakuFull
	for rows.Next() {
		bahan := entity.BahanBakuFull{}
		err := rows.Scan(&bahan.Id, &bahan.IdSatuan, &bahan.Nama)
		if err != nil {
			panic(err)
		}
		bahans = append(bahans, bahan)
	}
	return bahans
}

func (BahanRepository) FindById(ctx context.Context, tx *sql.Tx, id int) entity.BahanBakuFull {
	SQL := "SELECT id, satuan.nama, nama FROM bahan_baku INNER JOIN satuan ON bahan_baku.id_satuan = satuan.id WHERE bahan_baku.id=?;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	bahan := entity.BahanBakuFull{}
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var id int
		err := rows.Scan(&bahan.Id, &id, &bahan.Nama)
		if err != nil {
			return entity.BahanBakuFull{}
		}
	}
	return bahan
}
