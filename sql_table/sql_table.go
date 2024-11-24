package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Sql_table() {
	db, _ := sql.Open("sqlite3", "./scheduler.db")

	statement, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL,
		title TEXT NOT NULL,
		comment TEXT,
		repeat TEXT CHECK(length(repeat) <= 128)
	);
	CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	statement.Exec()

}