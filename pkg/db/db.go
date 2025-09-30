package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `CREATE TABLE scheduler (
id INTEGER PRIMARY KEY AUTOINCREMENT,
date CHAR(8) NOT NULL DEFAULT "",
title VARCHAR(256) NOT NULL DEFAULT "",
comment TEXT,
repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX scheduler_date ON scheduler (date);`

var DB *sql.DB

func Init(dbFile string) error {

	var err error
	var install bool

	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	_, err = os.Stat(dbFile)

	if err != nil {
		install = true
	}
	if install {
		_, err = DB.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}
