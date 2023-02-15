package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/produk/model/entity"
)

type ProdukRepository interface {
	FindAllProduk(ctx context.Context, tx *sql.Tx) []entity.Produk
	FindProdukById(ctx context.Context, tx *sql.Tx, id int) entity.Produk
	CreateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk)
	UpdateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk) entity.Produk
	DeleteProduk(ctx context.Context, tx *sql.Tx, id int)
}
