package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

func insertUser(db *pg.DB, user User) {
	db.Insert(&user)
}

func getUsersByName(db *pg.DB, name string) (users []User) {
	db.Model(&users).
		Where("name = ?", name).
		Select()
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
	createSchema(db)
	insertUser(db, User{Name: "Godot", Emails:[]string{"waiting@forhimto.com"}})

	for _, u := range getUsersByName(db, "Godot") {
		fmt.Println(u.String())
	}
}

func createSchema(db *pg.DB) error {
		err := db.CreateTable((*User)(nil), &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	return nil
}