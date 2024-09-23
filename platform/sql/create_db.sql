DROP DATABASE IF EXISTS mobileapp_rest;
CREATE DATABASE mobileapp_rest;

SET TIMEZONE="Europe/Moscow";

\c mobileapp_rest

-- Create register user table
CREATE TABLE users (
    id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    jwt VARCHAR (255) NOT NULL,
    fname VARCHAR (255) NOT NULL,
    email VARCHAR (255) NOT NULL,
    phonenumber VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    gender VARCHAR (255) NOT NULL,
    birthday VARCHAR NOT NULL,
    balance INT NOT NULL
);

-- Create donates table
CREATE TABLE transactions (
    id INT NOT NULL,
    iduser INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    phonenumber VARCHAR (255) NOT NULL,
    summary INT NOT NULL
);