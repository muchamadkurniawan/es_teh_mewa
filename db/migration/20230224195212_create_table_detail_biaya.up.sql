CREATE TABLE detail_biaya(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_bahan_baku INT NOT NULL,
    id_detail_pesanan INT DEFAULT NULL,
    id_rekap INT DEFAULT NULL,
    jumlah INT NOT NULL DEFAULT 1,
    harga_satuan INT NOT NULL DEFAULT 0,
    total INT NOT NULL DEFAULT 0,
    tanggal DATETIME NOT NULL default(CURRENT_TIMESTAMP),
    CONSTRAINT FK_bahanBiaya FOREIGN KEY (id_bahan_baku) REFERENCES bahan_baku(id),
    CONSTRAINT FK_pesananBiaya FOREIGN KEY (id_detail_pesanan) REFERENCES detail_pesanan(id),
    CONSTRAINT FK_rekapBiaya FOREIGN KEY (id_rekap) REFERENCES rekap(id)
)ENGINE = InnoDB;