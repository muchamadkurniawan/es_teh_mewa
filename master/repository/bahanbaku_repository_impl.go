package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/master/model/entity"
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

func (BahanRepository) Update(ctx context.Context, tx *sql.Tx, baku entity.BahanBaku) {
	SQL := "UPDATE bahan_baku SET id_satuan = ?, nama = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, baku.IdSatuan, baku.Nama, baku.Id)
	if err != nil {
		panic(err)
	}
}

func (BahanRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM bahan_baku WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}

func (BahanRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.BahanBakuFull {
	SQL := "SELECT bahan_baku.id, satuan.nama, bahan_baku.nama FROM bahan_baku INNER JOIN satuan ON bahan_baku.id_satuan = satuan.id ORDER BY bahan_baku.id DESC;"
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

func (BahanRepository) FindById(ctx context.Context, tx *sql.Tx, id int) entity.BahanBaku {
	SQL := "SELECT id, id_satuan, nama FROM bahan_baku WHERE bahan_baku.id=?;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	bahan := entity.BahanBaku{}
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		err := rows.Scan(&bahan.Id, &bahan.IdSatuan, &bahan.Nama)
		if err != nil {
			return entity.BahanBaku{}
		}
	}
	return bahan
}
