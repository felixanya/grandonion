package main

import (
	"database/sql"
	"github.com/aaronize/grandonion/examples/gorm-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	dbMap *sql.DB
)

func main() {
	defer dbMap.Close()
	// 存储
	mb := &model.MicroBlog{
		Author:  "aaron.chen",
		Content: "Hello, have a nice day.",
		Attribute: model.Attribute{
			map[string]interface{}{
				"tags":       []string{"technical", "personal", "golang", "database", "json"},
				"lang":       "zh",
				"location":   "shanghai,cn",
				"word_count": 109,
			},
		},
	}

	_, err := dbMap.Exec("INSERT INTO micro_blog (Author,Content,Attribute) VALUES (?,?,?)",
		mb.Author, mb.Content, mb.Attribute)
	if err != nil {
		log.Fatal("执行插入数据错误，", err)
		return
	}

	// 查询
	var mbq = &model.MicroBlog{}
	rows, err := dbMap.Query("SELECT * FROM micro_blog WHERE ID = ? LIMIT 1", 1)
	if err != nil {
		log.Fatal("查询数据库错误，", err.Error())
		return
	}
	for rows.Next() {
		err := rows.Scan(&mbq.ID, &mbq.Author, &mbq.Content, &mbq.Attribute, &mbq.CreateAt, &mbq.UpdateAt)
		if err != nil {
			log.Fatal("获取查询数据错误", err.Error())
			return
		}
	}
	for _, m := range mbq.Attribute {
		log.Println("tags: ", m["tags"])
		log.Println("location: ", m["location"])
		log.Println("language: ", m["lang"])
		log.Println("word_count: ", m["word_count"])
	}
}

func init() {
	var err error
	dbMap, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1)/testdb?parseTime=true")
	if err != nil {
		log.Fatal("连接数据库错误，", err)
		os.Exit(-1)
	} else {
		log.Println("数据库连接成功")
	}
}
