package models

import (
	"database/sql"
	"io/ioutil"
	"testing"
)

const dbScriptPath = "../../../database/"

func newTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=yarc_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	script, err := ioutil.ReadFile(dbScriptPath + "setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		script, err := ioutil.ReadFile(dbScriptPath + "teardown.sql")
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
