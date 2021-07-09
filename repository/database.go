package repository

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func OpenDB(url string) {
	// Connecting to DB
	DB = sqlx.MustOpen("mssql", url)

	err := DB.Ping()
	if err != nil {
		fmt.Println(3)
		panic(err)
	}
	fmt.Println("Connection to DB was established!")
}

func CloseDB() error {
	return DB.Close()
}
