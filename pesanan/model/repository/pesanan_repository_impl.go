package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/pesanan/model/entity"
	"es_teh_mewa/pesanan/web"
	"time"
)

type PesananRepositoryImpl struct{}

func NewPesananRepository() PesananRepository {
	return &PesananRepositoryImpl{}
}

func (PesananRepositoryImpl) GetIdProduk(ctx context.Context, tx *sql.Tx) []string {
	SQL := "Select id from produk_jual;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var ids []string
	for row.Next() {
		var id string
		err := row.Scan(&id)
		helperMain.PanicIfError(err)
		ids = append(ids, id)
	}
	return ids
}

func (PesananRepositoryImpl) GetProdukJualsAll(ctx context.Context, tx *sql.Tx) []web.ProdukJual {
	SQL := "SELECT produk_jual.id, bahan_baku.nama, produk_jual.harga FROM produk_jual INNER JOIN bahan_baku ON " +
		"bahan_baku.id = produk_jual.id_bahan_baku;"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var produks []web.ProdukJual
	for row.Next() {
		produk := web.ProdukJual{}
		err = row.Scan(&produk.Id, &produk.Bahan_baku, &produk.Harga)
		produks = append(produks, produk)
	}
	return produks
}

func (PesananRepositoryImpl) GetProdukJual(ctx context.Context, tx *sql.Tx, id int) web.ProdukJual {
	SQL := "SELECT produk_jual.id, bahan_baku.nama, produk_jual.harga FROM produk_jual INNER JOIN bahan_baku ON " +
		"bahan_baku.id = produk_jual.id_bahan_baku WHERE produk_jual.id = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	produk := web.ProdukJual{}
	if row.Next() {
		row.Scan(&produk.Id, &produk.Bahan_baku, &produk.Harga)
	}
	return produk
}

func (PesananRepositoryImpl) CreateBiayaPesanan(ctx context.Context, tx *sql.Tx, jumlah int, bahan int, harga int, id_detail int) (error, int) {
	currentTime := time.Now()
	SQL := "INSERT INTO detail_biaya(id_bahan_baku, id_detail_pesanan, jumlah, harga_satuan, total, tanggal) values(?, ?, ?, ?, ?, ?);"
	succ, err := tx.ExecContext(ctx, SQL, bahan, id_detail, jumlah, harga,
		jumlah*harga, currentTime.Format("2006-01-02 15:04:05"))
	idSu, _ := succ.LastInsertId()
	return err, int(idSu)
}

func (PesananRepositoryImpl) GetIdBahan(ctx context.Context, tx *sql.Tx, id int) int {
	SQL := "select id_bahan_baku from produk_jual where id = ?;"
	data, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var idBahan int
	if data.Next() {
		data.Scan(&idBahan)
	}
	return idBahan
}

func (PesananRepositoryImpl) GetHargaBiaya(ctx context.Context, tx *sql.Tx, id int) int {
	SQL := "select biaya from pembelian where id_bahan_baku = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var harga int
	if row.Next() {
		row.Scan(&harga)
	}
	return harga
}

func (PesananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) web.PesananRequestUpdate {
	SQL := "SELECT id, tanggal, pembayaran FROM pesanan WHERE id = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	defer helperMain.CloseRow(row)
	pesanan := web.PesananRequestUpdate{}
	//fmt.Println(row.Next())
	if row.Next() {
		err := row.Scan(&pesanan.Id, &pesanan.Date, &pesanan.Pembayaran)
		helperMain.PanicIfError(err)
		return pesanan
	} else {
		return pesanan
	}
}

func (i PesananRepositoryImpl) FindAllPesananDetail(ctx context.Context, tx *sql.Tx) []web.PesananRequestSum {
	SQL := "SELECT pesanan.id , cast(pesanan.tanggal as time) as jam, pesanan.pembayaran, detail_pesanan.SB as total FROM pesanan INNER JOIN " +
		"(SELECT id_pesanan, SUM(total) as SB FROM detail_pesanan GROUP BY id_pesanan) detail_pesanan ON detail_pesanan.id_pesanan = pesanan.id WHERE DATE(pesanan.tanggal) = CURDATE();"
	row, err := tx.QueryContext(ctx, SQL)
	helperMain.PanicIfError(err)
	var pesanans []web.PesananRequestSum
	for row.Next() {
		pesanan := web.PesananRequestSum{}
		err := row.Scan(&pesanan.Id, &pesanan.Time, &pesanan.Pembayaran, &pesanan.Total)
		helperMain.PanicIfError(err)
		pesanans = append(pesanans, pesanan)
	}
	row.Close()
	return pesanans
}

func (i PesananRepositoryImpl) FindAllDetailPesananId(ctx context.Context, tx *sql.Tx, id int, limit int) []web.PesananNamaBarang {
	SQL := "SELECT detail_pesanan.id as id, detail_pesanan.id_pesanan, bahan_baku.nama as nama FROM detail_pesanan INNER JOIN " +
		"produk_jual ON produk_jual.id = detail_pesanan.id_produk_jual INNER JOIN bahan_baku ON produk_jual.id_bahan_baku = bahan_baku.id WHERE detail_pesanan.id_pesanan = ? LIMIT ?;"
	row, err := tx.QueryContext(ctx, SQL, id, limit)
	helperMain.PanicIfError(err)
	var pesanans []web.PesananNamaBarang
	for row.Next() {
		pesanan := web.PesananNamaBarang{}
		err := row.Scan(&pesanan.Id_detail, &pesanan.Id_pesanan, &pesanan.Nama)
		helperMain.PanicIfError(err)
		pesanans = append(pesanans, pesanan)
	}
	row.Close()
	return pesanans
}

func (i PesananRepositoryImpl) ShowAllDetailPesananId(ctx context.Context, tx *sql.Tx, id int) []web.DetailRespon {
	SQL := "SELECT detail_pesanan.id, detail_pesanan.id_produk_jual, detail_pesanan.jumlah, detail_pesanan.harga_satuan, detail_pesanan.total, bahan_baku.nama FROM detail_pesanan INNER JOIN " +
		"produk_jual ON produk_jual.id = detail_pesanan.id_produk_jual INNER JOIN bahan_baku ON produk_jual.id_bahan_baku = bahan_baku.id WHERE detail_pesanan.id_pesanan = ?;"
	row, err := tx.QueryContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	var pesanans []web.DetailRespon
	for row.Next() {
		pesanan := web.DetailRespon{}
		err := row.Scan(&pesanan.Id, &pesanan.Id_produk, &pesanan.Jumlah, &pesanan.Harga, &pesanan.Total, &pesanan.Nama_produk)
		helperMain.PanicIfError(err)
		pesanans = append(pesanans, pesanan)
	}
	return pesanans
}

func (PesananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.PesananEntity {
	now := time.Now().Format("2006-01-02")
	SQL := "SELECT id, id_user, id_rekap, tanggal, pembayaran FROM pesanan WHERE tanggal = ?;"
	row, err := tx.QueryContext(ctx, SQL, now)
	helperMain.PanicIfError(err)
	var pesanan []entity.PesananEntity
	for row.Next() {
		newPesanan := entity.PesananEntity{}
		err := row.Scan(&newPesanan.Id, &newPesanan.Id_user, &newPesanan.Id_rekap, &newPesanan.Tanggal, &newPesanan.Pembayaran)
		helperMain.PanicIfError(err)
		pesanan = append(pesanan, newPesanan)
	}
	row.Close()
	return pesanan
}

//func (PesananRepositoryImpl) ShowIdDetailPesanan(ctx context.Context, tx *sql.Tx, id int) []int {
//	//TODO implement me
//	panic("implement me")
//}

func (PesananRepositoryImpl) CreatePesanan(ctx context.Context, tx *sql.Tx, request web.PesananRequestDateString) int {
	SQL := "INSERT INTO pesanan(id_user, pembayaran) VALUES(?, ?);"
	i, err := tx.ExecContext(ctx, SQL, request.Id_user, request.Pembayaran)
	helperMain.PanicIfError(err)
	id, err := i.LastInsertId()
	return int(id)
}

func (PesananRepositoryImpl) CreateDetail(ctx context.Context, tx *sql.Tx, request web.DetailRequest) (error, int) {
	SQL := "INSERT INTO detail_pesanan(id_produk_jual, id_pesanan, jumlah, harga_satuan, total) VALUES(?, ?, ?, ?, ?);"
	id, err := tx.ExecContext(ctx, SQL, request.Id_produk, request.Id_pesanan, request.Jumlah, request.Harga, request.Total)
	helperMain.PanicIfError(err)
	Id, _ := id.LastInsertId()
	return err, int(Id)
}

func (PesananRepositoryImpl) UpdatePembayaran(ctx context.Context, tx *sql.Tx, id int, pembayaran bool) error {
	SQL := "UPDATE pesanan SET pembayaran = ? WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, pembayaran, id)
	helperMain.PanicIfError(err)
	return nil
}

func (PesananRepositoryImpl) UpdateRekap(ctx context.Context, tx *sql.Tx, id int, id_rekap int) error {
	SQL := "UPDATE pesanan SET id_rekap = ? WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, id_rekap, id)
	helperMain.PanicIfError(err)
	return nil
}

func (PesananRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "DELETE FROM pesanan WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, id)
	helperMain.PanicIfError(err)
	return nil
}

func (PesananRepositoryImpl) CreateDetailStok(ctx context.Context, tx *sql.Tx, BP int, jumlah int) error {
	SQL := "insert into detail_stok(id_bahan_baku, type, jumlah) values(?, 'out', ?)"
	_, err := tx.ExecContext(ctx, SQL, BP, jumlah)
	SQL2 := "update stok set total = total - ? where id_bahan_baku = ?;"
	_, err = tx.ExecContext(ctx, SQL2, jumlah, BP)
	helperMain.PanicIfError(err)
	return nil
}
