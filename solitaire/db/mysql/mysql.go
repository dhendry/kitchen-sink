package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func foo() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	_ = db
}
