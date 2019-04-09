package main

import (
	"fmt"
	"github.com/go-pg/pg"
)

type UserPlain struct {
	Id     int64
	Name   string
}

func (u UserPlain) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Name)
}


func insertUserPlain(db *pg.Tx, user UserPlain) {
	db.ExecOne("INSERT INTO users (name) values(?)", user.Name)
	return
}

func getUsersByNamePlain(db *pg.Tx, name string) (users []UserPlain) {
	db.QueryOne(&users, "SELECT id, name, name || ' unexpected' from users where name = ?", name)
	return
}

func printNamesWithUnKnownFieldsinSelect(db *pg.Tx) {
	var results []struct {
		Name        string
		FirstLetter string
	}

	db.Query(&results, "SELECT name, substring(name, 1,1) as first_letter from users")
	for _, r := range results {
		fmt.Printf("%s: %s\n", r.FirstLetter, r.Name)
	}
}

func prepPlain(db *pg.Tx) (stmt *pg.Stmt) {
	 stmt, _ = db.Prepare("SELECT id from users where id = $1")
	 return
}

func main() {
	db := pg.Connect(&pg.Options{
		PoolSize: 1,
		User: "postgres",
		Password: "postgres",
		Database: "go_demo",
	})
	defer db.Close()

	tx, _ := db.Begin()
	createSchemaPlain(tx)
	tx.Commit()

	tx, _ = db.Begin()
	insertUserPlain(tx, UserPlain{Name: "Godot"})
	tx.Commit()

	db.RunInTransaction(func(transaction *pg.Tx ) error {
		for _, u := range getUsersByNamePlain(transaction, "Godot") {
			fmt.Println(u.String())
		}
		printNamesWithUnKnownFieldsinSelect(tx)
		return nil
	})

	var id int64

	tx, _ = db.Begin()
	prepPlain(tx).QueryOne(&id, 1)
	fmt.Println(id)
	tx.Commit()
	tx, _ = db.Begin()
	insertUserPlain(tx, UserPlain{Name: "A"})
	fmt.Println("length in tx :", len(getUsersByNamePlain(tx, "A")))
	tx.Rollback()
	tx, _ = db.Begin()
	fmt.Println("length after tx :", len(getUsersByNamePlain(tx, "A")))
	tx.Commit()

}

func createSchemaPlain(db *pg.Tx) {
		db.ExecOne("CREATE TABLE users (" +
			"id bigserial primary key not null," +
			"name text)")
		return
}