package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/pesanan/model/entity"
	"eh_teh_mewa/pesanan/web"
)

type PesananRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) entity.PesananEntity
	FindAll(ctx context.Context, tx *sql.Tx) []entity.PesananEntity
	Create(ctx context.Context, tx *sql.Tx, request web.PesananRequestDateString) error
	Update(ctx context.Context, tx *sql.Tx, request web.PesananResponseDateString) error
	UpdateRekap(ctx context.Context, tx *sql.Tx, id int) error
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
