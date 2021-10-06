package bot

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, e := sql.Open("mysql", "dbuser:dbpass@tcp(127.0.0.1:3306)/dbname")
	ErrorCheck(e)

	e = db.Ping()
	ErrorCheck(e)

	return db
}
