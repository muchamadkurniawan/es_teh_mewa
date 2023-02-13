package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/produk/model/web"
)

type ProdukRepository interface {
	FindAllProduk(ctx context.Context, tx *sql.Tx) ([]web.ResponseProduk, error)
	FindAllProdukById(ctx context.Context, tx *sql.Tx, id int) (web.ResponseProduk, error)
	CreateProduk(ctx context.Context, tx *sql.Tx, produk web.RequestProduk) (web.RequestProduk, error)
	UpdateProduk(ctx context.Context, tx *sql.Tx, produk web.RequestProduk) (web.RequestProduk, error)
	DeleteProduk(ctx context.Context, tx *sql.Tx, id int) error
}
