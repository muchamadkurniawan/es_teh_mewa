CREATE TABLE bahan_baku(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_satuan INT NOT NULL,
    nama VARCHAR(20) NOT NULL,
    CONSTRAINT FK_SatuanBahanBaku FOREIGN KEY (id_satuan) REFERENCES satuan(id)
) ENGINE = InnoDB;