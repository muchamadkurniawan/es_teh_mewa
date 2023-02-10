CREATE TABLE produk_jual
(
    id INT NOT NULL PRIMARY KEY,
    id_user INT NOT NULL,
    id_bahan_baku INT NOT NULL,
    harga INT NOT NULL default(0),
    stock BOOLEAN default(FALSE),
    CONSTRAINT FK_UserProdukJual FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT FK_BahanBakuProduk FOREIGN KEY (id_bahan_baku) REFERENCES bahan_baku(id)
) ENGINE = InnoDB;