package db

import (
	"database/sql"
	"io/ioutil"

	"github.com/HunterGooD/testsGolang/Authorization/utils"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(dbname string) (*sql.DB, error) {
	if !utils.FileIsExist(dbname) {
		err := utils.CreateFile(dbname)
		if err != nil {
			panic("can't create database")
		}
	}
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}

	sql, err := ioutil.ReadFile(dbname)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(string(sql))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectDB(databasename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databasename)
	if err != nil {
		return nil, err
	}
	return db, err
}
