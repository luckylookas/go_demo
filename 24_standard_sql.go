package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //_ = import for effect, meaning you do not want to USE any of the functions or types, but you need it there
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "go_demo"
)

func getAllNames(db *sql.DB) (names []string) {
	rows, _ := db.Query("select name from users;")
	var name string
	for rows.Next() {
		rows.Scan(&name)
		names = append(names, name)
	}
	return
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Exec("create table if not exists users(id bigserial primary key not null, name text);")
	db.Exec("insert into users (name) values($1);", "Rogelio")

	fmt.Println(getAllNames(db))
}
