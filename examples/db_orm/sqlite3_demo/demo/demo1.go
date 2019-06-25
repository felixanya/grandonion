package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./demo1_foo.db")
	if err != nil {
		log.Fatalf("open db error, %q", err)
	}
	defer db.Close()

	//err = CreateTable()
	//if err != nil {
	//	log.Fatalf("create table error, %q", err)
	//}

	//err = InsertIntoData()
	//if err != nil {
	//	log.Fatalf("insert into error, %q", err)
	//}

	rows, err := QueryData()
	if err != nil {
		log.Fatalf("query data error, %q", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var dept string
		err = rows.Scan(&id, &name, &dept)
		if err != nil {
			log.Fatalf("scan data error, %q", err)
		}
		fmt.Printf("No.%d, Name: %s, Dept: %s\n", id, name, dept)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("rows.Err: %q", err)
	}

}

func CreateTable() error {
	sqlStatement := `
		CREATE TABLE users(id integer not null primary key, name text, dept text);
		DELETE FROM users;
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		//log.Fatalf("%q: %s \n", err, sqlStatement)
		return err
	}

	return nil
}

func InsertIntoData() error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%q", err)
	}

	stmt, err := tx.Prepare("INSERT INTO users(id, name, dept) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("%03d", i), "autoD")
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

func QueryData() (*sql.Rows, error) {
	rows, err := db.Query("SELECT id, name, dept FROM users")
	if err != nil {
		return nil, err
	}

	return rows, nil
}
