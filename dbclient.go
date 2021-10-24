package main

import "database/sql"

type DBClient struct {
	db *sql.DB
}

func NewDBClient() *DBClient {
	cfg := NewConfig()
	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		panic("db connection failed: " +  err.Error())
	}
	return &DBClient{
		db: db,
	}
}