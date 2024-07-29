CREATE DATABASE IF NOT EXISTS adotapet;
USE adotapet;

DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique, 
    telefone varchar(50) not null,
    senha varchar(50) not null
) ENGINE = INNODB;

CREATE TABLE pets(
    id int auto_increment primary key,
    nome varchar(50) not null,
    raca varchar(50) not null, 
    cor  varchar(50) not null,
    tamanho  varchar(20) not null,
    sexo  varchar(20) not null,
    idade TINYINT UNSIGNED not null,
    peso float not null,
    usuarioID int UNSIGNED NOT NULL,
    FOREIGN KEY (usuarioID) REFERENCES usuarios(id)
) ENGINE = INNODB;


