CREATE TABLE pesanan
(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    id_rekap INT DEFAULT NULL,
    tanggal DATETIME NOT NULL default(CURRENT_TIMESTAMP),
    pembayaran BOOL DEFAULT(FALSE),
    CONSTRAINT FK_UserPesanan FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT FK_RekapPesanan FOREIGN KEY (id_rekap) REFERENCES rekap(id)
) ENGINE = InnoDB;