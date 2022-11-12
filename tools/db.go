package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DBConnector struct {
	host, port, username, password string
}

func NewDBConnector(username, password, host, port, db string) *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, db)
	connectedDB, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatalf("DB connect failed %s, db info %s\n", err, url)
	}
	err = connectedDB.Ping()
	if err != nil {
		log.Fatalf("DB Ping failed %s, db info %s\n", err, url)

	}
	log.Printf("Connected database %s:%s@tcp(%s:%s)/%s\n", username, "***", host, port, db)

	return connectedDB
}
