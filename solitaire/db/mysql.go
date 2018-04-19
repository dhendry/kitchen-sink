package db

import (
	"database/sql"

	"github.com/dhendry/kitchen-sink/solitaire/model"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// TODO: Fix root user obviously
	rawDb, err := sql.Open("mysql", "root:@/solitaire")
	if err != nil {
		panic(err)
	}

	Db = &mySqlSolitaireDataAccess{rawDb:rawDb}
}

type mySqlSolitaireDataAccess struct {
	rawDb *sql.DB
}

func (sda *mySqlSolitaireDataAccess) SaveNewGameState(gs *model.GameState) error {
	return nil
}
