package bot

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, e := sql.Open("mysql", "smirkwoodia:#Vlacas23@tcp(127.0.0.1:3306)/robindb")
	ErrorCheck(e)

	e = db.Ping()
	ErrorCheck(e)

	return db
}
