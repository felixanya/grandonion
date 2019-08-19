package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main() {
    db, err := sql.Open("mysql", "root:3306@localhost/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //result, err := db.Exec("SELECT * FROM `table_name` WHERE condition1 = ?", 123)

    rows, err := db.Query("SELECT * FROM `table_name` WHERE condition1 = ?", 123)
    if err != nil {
        log.Fatal(err)
    }

    for {
        if !rows.Next() {
            break
        }

    }
}
