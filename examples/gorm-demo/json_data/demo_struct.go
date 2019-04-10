package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

/*
以结构体的方式存入Json
结构体需要提前指定，完全没有灵活性可言
 */
type Media struct {
	Path string
	Type string
}

type MediaList []*Media

var _ sql.Scanner = &MediaList{}
var _ driver.Valuer = MediaList{}

//继承Scanner（Scan接受的是指针类型）
func (m *MediaList) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	b, _ := src.([]byte)
	return json.Unmarshal(b, m)
}

//继承Valuer（INSERT时，Valuer不接受指针类型）
func (m MediaList) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type Weibo struct {
	ID        int
	Author    string
	Content   string
	Media     MediaList
	CreatedAt *time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gormdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//插入
	weibo := &Weibo{
		Author:  "joe",
		Content: "nice too meet you!",
		Media: MediaList{
			{"http://www.google.com", "search"},
			{"http://cn.bing.com", "search"},
		},
	}
	_, err = db.Exec("INSERT INTO contents (author,content,media) VALUES (?,?,?)",
		weibo.Author, weibo.Content, weibo.Media) //这里会调用Valuer(不接受指针类型)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("insert success.")
	}

	//查询
	weibo = &Weibo{}
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
	log.Println(weibo)
	log.Println(weibo.Media[0])
}

