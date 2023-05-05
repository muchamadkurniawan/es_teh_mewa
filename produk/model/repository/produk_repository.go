package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/produk/model/entity"
	"es_teh_mewa/produk/model/web"
)

type ProdukRepository interface {
	FindAllProduk(ctx context.Context, tx *sql.Tx) []web.ResponseProdukFull
	FindProdukById(ctx context.Context, tx *sql.Tx, id int) entity.Produk
	FindProdukByBahan(ctx context.Context, tx *sql.Tx, bahan int) (entity.Produk, error)
	CreateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk)
	UpdateProduk(ctx context.Context, tx *sql.Tx, produk entity.Produk)
	DeleteProduk(ctx context.Context, tx *sql.Tx, id int)
	FindAllProdukByBahan(ctx context.Context, tx *sql.Tx, barang int) []web.ResponseProdukFull
}
