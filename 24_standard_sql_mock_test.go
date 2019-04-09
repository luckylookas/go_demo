package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/assert"
	"testing"
)

func TestGetAllNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	mock.
		ExpectQuery("select name from users;").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob"))

	names := getAllNames(db)

	assert.Assert(t, len(names) == 2)
	assert.Assert(t, contains(names, "Alice"))
	assert.Assert(t, contains(names, "Bob"))

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}