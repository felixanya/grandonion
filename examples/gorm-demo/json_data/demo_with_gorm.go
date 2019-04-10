package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type MediaLists []map[string]string

var _ sql.Scanner = &MediaLists{}
var _ driver.Valuer = MediaLists{}

//继承Scanner（Scan接受的是指针类型）
func (m *MediaLists) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	b, _ := src.([]byte)
	return json.Unmarshal(b, m)
}

//继承Valuer（INSERT时，Valuer不接受指针类型）
func (m MediaLists) Value() (driver.Value, error) {
	return json.Marshal(m)
	//return m, nil
}

type Weibos struct {
	ID        int
	Author    string
	Content   string
	Media     MediaLists
	CreatedAt *time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gormdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//插入
	//weibo := &Weibos{
	//	Author:  "joe",
	//	Content: "nice too meet you!",
	//	//Media: []map[string]string{
	//	//	map[string]string{"path": "http://www.google.com", "type": "1"},
	//	//	map[string]string{"path": "http://cn.bing.com", "type": "1"},
	//	//},
	//	Media: MediaLists{
	//		map[string]string{"path": "http://www.google.com", "type": "search"},
	//		map[string]string{"path": "http://cn.bing.com", "type": "search"},
	//	},
	//}
	//_, err = db.Exec("INSERT INTO contents (author,content,media) VALUES (?,?,?)",
	//	weibo.Author, weibo.Content, weibo.Media) //这里会调用Valuer(不接受指针类型)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println("insert success.")
	//}

	//查询
	weibo := &Weibos{}
	rows, err := db.Query("SELECT * FROM contents WHERE id = ? LIMIT 1", 10)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&weibo.ID, &weibo.Author, &weibo.Content, &weibo.Media, &weibo.CreatedAt) //这里会调用Scanner
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("query: ", weibo)
	log.Println("json: ", weibo.Media)
	for _, m := range weibo.Media {
		log.Println("json: ", m)
	}
}

