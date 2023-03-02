package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/rekap/web"
)

type RekapRepositoryImpl struct{}

func (RekapRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (web.RekapResponseTime, error) {
	//TODO implement me
	panic("implement me")
}

func (RekapRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]web.RekapResponseTime, error) {
	//TODO implement me
	panic("implement me")
}

func (RekapRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error {
	//TODO implement me
	panic("implement me")
}

func (RekapRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error {
	//TODO implement me
	panic("implement me")
}

func (RekapRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	//TODO implement me
	panic("implement me")
}
