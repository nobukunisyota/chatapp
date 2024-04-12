package model

import "database/sql"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=**** port=**** user=**** password=**** dbname=**** sslmode=disable")
	if err != nil {
		panic(err)
	}
}
