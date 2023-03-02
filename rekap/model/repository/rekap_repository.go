package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/rekap/web"
)

type RekapRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) (web.RekapResponseTime, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]web.RekapResponseTime, error)
	Create(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error
	Update(ctx context.Context, tx *sql.Tx, request web.RekapRequestDateString) error
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
