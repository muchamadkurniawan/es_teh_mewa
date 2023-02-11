package web

type PembelianUpdateRequest struct {
	Id            int
	Id_user       int
	Id_bahan_baku int
	Tanggal       string
	Biaya         int
	Jumlah        int
	Use_pembelian bool
}
