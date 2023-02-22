package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func newTestDB(t *testing.T) (*sql.DB, func()) {

	dbUser, dbPass := os.Getenv("DBUSER"), os.Getenv("DBPASS")
	conf := fmt.Sprintf("%s:%s@/%s?parseTime=true&multiStatements=true", dbUser, dbPass, "test_snippetbox")
	db, err := sql.Open("mysql", conf)
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	}
}
