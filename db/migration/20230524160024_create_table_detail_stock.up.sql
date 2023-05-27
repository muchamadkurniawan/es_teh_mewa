CREATE TABLE detail_stok(
     id int not null auto_increment primary key,
     id_bahan_baku INT NOT NULL,
     type enum('out', 'in'),
     jumlah int default 0,
     CONSTRAINT FK_detailstokbahan FOREIGN KEY (id_bahan_baku) REFERENCES bahan_baku(id)
)ENGINE = InnoDB;