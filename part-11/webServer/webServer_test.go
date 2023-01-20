package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const CONNECT_STR = "user=postgres dbname=s2 sslmode=disable"

func getDbConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", CONNECT_STR)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func create_table() {
	db, err := getDbConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	const query = `
		create table if not exists users (
			id serial primary key,
			first_name text,
			last_name text
	)`

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func drop_table() {
	db, err := getDbConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("drop table if exists users")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func instert_record(query string) {
	db, err := getDbConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insert_user(first_name, last_name string) {
	instert_record(fmt.Sprintf("insert into users (first_name, last_name) values ('%s', '%s')", first_name, last_name))
}

func Test_count(t *testing.T) {
	var count int
	create_table()
	defer drop_table()

	insert_user("Misha", "1")
	insert_user("Misha", "2")
	insert_user("Misha", "3")

	db, err := getDbConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	row := db.QueryRow("select count(*) from users")
	err = row.Scan(&count)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Close()
	if count != 3 {
		t.Errorf("Select query returned %d", count)
	}
}

func Test_queryDB(t *testing.T) {
	var col1 int
	var col2 string
	var col3 string

	create_table()
	defer drop_table()
	db, err := getDbConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	insert_user("Misha", "1")

	rows, err := db.Query(`SELECT * FROM users WHERE last_name=$1`, `1`)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		rows.Scan(&col1, &col2, &col3)
	}
	if col2 != "Misha" {
		t.Errorf("first_name returned %s", col2)
	}
	if col3 != "1" {
		t.Errorf("last_name returned %s", col3)
	}
}

func Test_record(t *testing.T) {
	create_table()
	defer drop_table()

	insert_user("Misha", "1")
	req, err := http.NewRequest("GET", "/getdata", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getDataHandler)
	handler.ServeHTTP(recorder, req)
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned %v", status)
	}
	if recorder.Body.String() != "1, Misha, 1\n" {
		t.Errorf("Wrong server response!: '%s'", recorder.Body.String())
	}
}
