CREATE TABLE detail_pesanan(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_produk_jual INT NOT NULL,
    id_pesanan INT NOT NULL,
    jumlah INT DEFAULT 1,
    harga_satuan INT DEFAULT 0,
    total INT DEFAULT 0,
    CONSTRAINT FK_produkPesanan FOREIGN KEY (id_produk_jual) REFERENCES produk_jual(id),
    CONSTRAINT FK_pesananDetail FOREIGN KEY (id_pesanan) REFERENCES pesanan(id)
) ENGINE = InnoDB;