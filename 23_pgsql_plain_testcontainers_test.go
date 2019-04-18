package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/go-pg/pg"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"testing"
	"time"
)

//for now, docker api version 1.39+ is not supported by go testcontainers
//env DOCKER_API_VERSION=1.38 go test 23_pgsql_plain.go 23_pgsql_plain_testcontainers_test.go

type WaitForPostgresStrategy struct {
	Port nat.Port
	Database string
	User string
	Password string
}

func (this WaitForPostgresStrategy) WaitUntilReady(ctx context.Context, target wait.StrategyTarget) error {
	port, err := target.MappedPort(ctx, this.Port)
	host, err := target.Host(ctx)
	if err != nil {
		return errors.New("container startup failed")
	}


	db := pg.Connect(&pg.Options{
		Addr: fmt.Sprintf("%s:%s", host, port.Port()),
		Database: this.Database,
		User: this.User,
		Password: this.Password,
	})

	for i := 0; i < 60 ; i++ {
		res, err := db.ExecOne("SELECT 1")
		if err != nil || res.RowsReturned() != 1 {
			time.Sleep(100 * time.Millisecond)
		} else {
			return nil
		}
	}
	return errors.New("could not connect to postgres testcontainer in time")
}

func Test_insertUserPlainTestContainers(t *testing.T) {
	databaseName := "go_demo"
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		WaitingFor: WaitForPostgresStrategy{
			Port: "5432/tcp",
			User: "postgres",
			Password: "postgres",
			Database: databaseName,
		},
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{"POSTGRES_DB": databaseName},
	}

	container, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	defer container.Terminate(ctx)

	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "5432/tcp")

	fmt.Println(host)
	fmt.Println(port)

	db := pg.Connect(&pg.Options{
		Addr: fmt.Sprintf("%s:%s", host, port.Port()),
		Database: "go_demo",
		User: "postgres",
	})


	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer 	tx.Rollback()
	tx.ExecOne("create table users (id bigserial primary key not null, name text)")
	tx.ExecOne("insert into users (name) values (?)", "Jorge")
	var count int
	tx.QueryOne(&count, "select count(*) from users")
	fmt.Println(count)
}