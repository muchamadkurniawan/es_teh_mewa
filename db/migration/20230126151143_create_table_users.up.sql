CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    password VARCHAR(20) NOT NULL,
    type ENUM('admin', 'kasir', 'super_user') NOT NULL default('admin')
) ENGINE = InnoDB;