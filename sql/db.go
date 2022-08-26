package sql

import (
	"database/sql"
	"fmt"
	"log"
)

var dB *sql.DB

func InitializeDB(dbName, root, pwd, ipAndPort, charset string) {
	daraSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True", root, pwd, ipAndPort, dbName, charset)
	db, err := sql.Open("mysql", daraSourceName)
	if err != nil {
		log.Fatal(err)
	}
	dB = db
}
