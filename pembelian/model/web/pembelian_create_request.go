package web

type PembelianCreateRequest struct {
	Id            int
	Id_user       int
	Id_bahan_baku int
	Tanggal       string
	Biaya         int
	Jumlah        int
	Total         int
	Use_pembelian bool
}
