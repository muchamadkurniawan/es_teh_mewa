package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/pesanan/model/entity"
	"eh_teh_mewa/pesanan/web"
)

type PesananRepositoryImpl struct{}

func (PesananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) entity.PesananEntity {
	SQL := "SELECT id, id_user, id_rekap, tanggal, pembayaran FROM pesanan WHERE id = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var pesanan entity.PesananEntity
	if row.Next() {
		err := row.Scan(&pesanan.Id, &pesanan.Id_user, &pesanan.Id_rekap, &pesanan.Tanggal, &pesanan.Pembayaran)
		helperMain.PanicIfError(err)
		return pesanan
	} else {
		return pesanan
	}
}

func (PesananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.PesananEntity {
	SQL := "SELECT id, id_user, id_rekap, tanggal, pembayaran FROM pesanan;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var pesanan []entity.PesananEntity
	for row.Next() {
		newPesanan := entity.PesananEntity{}
		err := row.Scan(&newPesanan.Id, &newPesanan.Id_user, &newPesanan.Id_rekap, &newPesanan.Tanggal, &newPesanan.Pembayaran)
		helperMain.PanicIfError(err)
		pesanan = append(pesanan, newPesanan)
	}
	return pesanan
}

func (PesananRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request web.PesananRequestDateString) error {
	//TODO implement me
	panic("implement me")
}

func (PesananRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request web.PesananResponseDateString) error {
	//TODO implement me
	panic("implement me")
}

func (PesananRepositoryImpl) UpdateRekap(ctx context.Context, tx *sql.Tx, id int) error {
	//TODO implement me
	panic("implement me")
}

func (PesananRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	//TODO implement me
	panic("implement me")
}
