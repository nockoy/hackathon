package main

import (
	"db/dao"
	"db/router"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func init() {
	dao.DBInit()
}

func main() {

	http.HandleFunc("/user", router.UserHandler)
	http.HandleFunc("/user2", router.UserHandler2)
	http.HandleFunc("/members", router.MemberHandler)
	http.HandleFunc("/message", router.MessageHandler)
	http.HandleFunc("/reply", router.ReplyHandler)
	http.HandleFunc("/channel", router.ChannelIDHandler)
	http.HandleFunc("/channel/join", router.UserChannelHandler)
	http.HandleFunc("/channel/other", router.OtherChannelHandler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.DBClose()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
