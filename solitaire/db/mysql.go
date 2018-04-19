package db

import (
	"database/sql"

	"github.com/dhendry/kitchen-sink/solitaire/model"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *MySqlSolitaireDataAccess
)

func foo() {
	// TODO: Fix root user obviously
	db, err := sql.Open("mysql", "root:@/solitaire")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

type MySqlSolitaireDataAccess struct {
	Db *sql.DB
}

func (db *MySqlSolitaireDataAccess) SaveNewGameState(gs *model.GameState) error {
	return nil
}
