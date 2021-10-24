package main

import (
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type PqConfig struct {
	User string
	DbName string
	Password string
	Host string
	SslMode	string
}

func NewConfig() *PqConfig {
	return &PqConfig{
		User: "test_user",
		Password: "test_pass",
		Host: "testdb",
		DbName: "test",
		SslMode: "disable",
	}
}

func (pqc *PqConfig) Validate() error {
	if pqc.User == "" {
		return errors.New("no user defined for configuring database connection")
	}
	if pqc.DbName == "" {
		return errors.New("no db-name defined for configuring database connection")
	}
	if pqc.Password == "" {
		return errors.New("no password defined for configuring database connection")
	}
	if pqc.Host == "" {
		return errors.New("no host defined for configuring database connection")
	}
	return nil
}

func (pqc *PqConfig) FormatDSN() string {
	err := pqc.Validate()
	if err != nil {
		panic(err.Error())
	}

	conf := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", pqc.Host, pqc.User, pqc.Password, pqc.DbName)
	if pqc.SslMode != "" {
		conf += fmt.Sprintf(" sslmode=%s", pqc.SslMode)
	}
	return conf
}