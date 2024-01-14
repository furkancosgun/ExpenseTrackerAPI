CREATE DATABASE expense_tracker;

--Users Table
CREATE TABLE users (
    email VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(255),
    account_confirmed BOOLEAN
);


--Users Tokens
CREATE TABLE tokens (
    email VARCHAR(255) PRIMARY KEY,
    token VARCHAR(255),
    expires_at timestamp
)


