DROP DATABASE IF EXISTS cris;
CREATE DATABASE cris CHARACTER SET utf8mb4;

USE cris;

CREATE TABLE urls (
  uuid CHAR(64) PRIMARY KEY,
  url VARCHAR(255) NOT NULL UNIQUE
) ENGINE=InnoDB;
