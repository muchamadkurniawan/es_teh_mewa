CREATE TABLE stok(
    id int not null auto_increment primary key,
    id_bahan_baku INT NOT NULL,
    total int default 0,
    CONSTRAINT FK_stokbahan FOREIGN KEY (id_bahan_baku) REFERENCES bahan_baku(id)
)ENGINE = InnoDB;