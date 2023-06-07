package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB


func DBInit() {
		// DB接続のための準備
		mysqlUser := os.Getenv("MYSQL_USER")
		mysqlPwd := os.Getenv("MYSQL_PWD")
		mysqlHost := os.Getenv("MYSQL_HOST")
		mysqlDatabase := os.Getenv("MYSQL_DATABASE")

		connStr := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
		_db, err := sql.Open("mysql", connStr)

		if err != nil {
			log.Fatalf("fail: sql.Open, %v\n", err)
		}
		// ①-3
		if err := _db.Ping(); err != nil {
			log.Fatalf("fail: _db.Ping, %v\n", err)
		}
		db = _db
	}



// func DBInit() {
// 	// ①-1
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf("読み込み出来ませんでした: %v", err)
// 	}
// 	mysqlUser := os.Getenv("MYSQL_USER")
// 	mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
// 	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
// 	// ①-2
// 	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s?parseTime=true", mysqlUser, mysqlUserPwd, mysqlDatabase))
// 	if err != nil {
// 		log.Fatalf("fail: sql.Open, %v\n", err)
// 	}
// 	// ①-3
// 	if err := _db.Ping(); err != nil {
// 		log.Fatalf("fail: _db.Ping, %v\n", err)
// 	}
// 	db = _db
// }

func DBClose() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
