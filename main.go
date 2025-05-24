package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type AppEntry struct {
	Id       string `db:"id"`
	UserName string `db:"username"`
	Email    string `db:"email"`
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/big_event")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 执行查询
	rows, err := db.Query("SELECT id, username,email FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 读取查询结果
	var appEntries []AppEntry
	for rows.Next() {
		var entry AppEntry
		if err := rows.Scan(&entry.Id, &entry.UserName, &entry.Email); err != nil {
			log.Fatal(err)
		}
		appEntries = append(appEntries, entry)
	}

	// 检查查询是否出错
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 打印查询结果
	for _, entry := range appEntries {
		fmt.Println(entry.Id, entry.UserName, entry.Email)
	}
}
