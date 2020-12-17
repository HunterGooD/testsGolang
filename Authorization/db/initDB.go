package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(dbname string) (*sql.DB, error) {
	if !utils.FileIsExist(dbname) {
		err = utils.CreateFile(dbname)
		if err != nil {
			panic("can't create database")
		}
	}
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	return db, nil
}
