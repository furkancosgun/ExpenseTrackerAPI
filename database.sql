-- Veritabanı oluştur
CREATE DATABASE expense_tracker;

-- Kullanıcı tablosu oluştur
CREATE TABLE users (
    email VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(255),
    account_confirmed BOOLEAN
);

