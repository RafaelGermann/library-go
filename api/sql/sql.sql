CREATE DATABASE IF NOT EXISTS libraryGo;
USE libraryGo;

DROP TABLE IF EXISTS Seguidores;
DROP TABLE IF EXISTS Usuarios;


CREATE TABLE Usuarios
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    nome     VARCHAR(50)  NOT NULL,
    nick     VARCHAR(50)  NOT NULL UNIQUE,
    email    VARCHAR(50)  NOT NULL UNIQUE,
    senha    VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT current_timestamp()
) ENGINE = INNODB;

CREATE TABLE Seguidores
(
    usuario_id  int NOT NULL,
    FOREIGN key (usuario_id) REFERENCES Usuarios (id) ON DELETE CASCADE,
    seguidor_id int,
    FOREIGN key (seguidor_id) REFERENCES Usuarios (id) ON DELETE CASCADE,
    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE = INNODB;