package main

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	sqlite3conn := make([]*sqlite3.SQLiteConn, 0)
	sql.Register("sqlite3_with_hook_example",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				sqlite3conn = append(sqlite3conn, conn)
				conn.RegisterUpdateHook(func(op int, db string, table string, rowid int64) {
					switch op {
					case sqlite3.SQLITE_INSERT:
						log.Printf("Notified of insert on db: %s, table: %s, rowid: %d", db, table, rowid)
					case sqlite3.SQLITE_UPDATE:
						log.Printf("Notified of update on db: %s, table: %s, rowid: %d", db, table, rowid)
					}
				})
				return nil
			},
		})
	os.Remove("./sqlite3_hook_foo.db")
	os.Remove("./sqlite3_hook_bar.db")

	srcDb, err := sql.Open("sqlite3_with_hook_example", "./sqlite3_hook_foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer srcDb.Close()
	srcDb.Ping()

	_, err = srcDb.Exec("CREATE TABLE host_info (id int, ip text, sn text, status text, create_at datetime, update_at datetime, info text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = srcDb.Exec("INSERT INTO host_info (id, ip, sn, status, create_at, update_at, info) VALUES (1000, '192.168.168.10', 'SN12345567', 'cleaning', '2019-06-03', '2019-06-03', 'clean messages')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = srcDb.Exec("INSERT INTO host_info (id, ip, sn, status, create_at, update_at, info) VALUES (1001, '192.168.152.174', 'SN190345525', 'reboot', '2019-06-03', '2019-06-03', 'reboot messages')")
	if err != nil {
		log.Fatal(err)
	}

	_, err = srcDb.Exec("CREATE TABLE port_info (id int, ip text, sn text, switch_man_ip text, switch_name text, port_name text, mac text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = srcDb.Exec("INSERT INTO port_info (id, ip, sn, switch_man_ip, switch_name, port_name, mac) VALUES (1042, '192.168.168.10', 'SN12345567', '192.168.11.1', 'SW.N-01', 'GE1/0/12', '52:54:00:59:85:50')")
	if err != nil {
		log.Fatal(err)
	}
}
