package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// данные для подключения к бд
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

var Db *sql.DB

func InitDb() error {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	var err error
	Db, err = sql.Open("postgres", psqlConn)
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
