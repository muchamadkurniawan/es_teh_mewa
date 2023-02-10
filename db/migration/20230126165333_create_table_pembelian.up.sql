CREATE TABLE pembelian
(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    id_bahan_baku INT NOT NULL,
    tanggal DATE NOT NULL default(CURRENT_DATE),
    biaya INT NOT NULL default(0),
    use_pembelian BOOLEAN default(FALSE),
    CONSTRAINT FK_UserPembelian FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT FK_BahanBakuPembelian FOREIGN KEY (id_bahan_baku) REFERENCES bahan_baku(id)
) ENGINE = InnoDB;