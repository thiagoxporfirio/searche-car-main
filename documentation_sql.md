# SQL

Colocarei as documentações das tabelas do banco, para caso alguma duvida possa consultar como está os formatos de entrada para o banco.


## TABELA USERINFO

CREATE TABLE userInfo (
    userId varchar(255) PRIMARY KEY,
    username varchar(12) NOT NULL,
    password varchar(16) NOT NULL,
    telefone varchar(15),
    name     varchar(50) NOT NULL,
    email    varchar(100) NOT NULL ,
    permission varchar(1) NOT NULL,
    cars int NOT NULL,
);

CREATE TABLE car (
    placa varchar(12) PRIMARY KEY,
    userId varchar(255) NOT NULL,
    renavam varchar(50) NOT NULL,
    state varchar(2) NOT NULL,
    marcaEModelo varchar(50) NOT NULL,
    municipio varchar(25) NOT NULL,
    anoDoCarro varchar(4) NOT NULL,
    cor varchar(15) NOT NULL,
    chassi varchar(50) NOT NULL,
    nome varchar(45) NOT NULL
);

