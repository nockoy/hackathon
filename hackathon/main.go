package main

import (
	"db/controller"
	"db/dao"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func init() {
	dao.DBInit()
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.SearchUser(w, r)

	case http.MethodPost:
		controller.RegisterUser(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.GetMessage(w, r)
	case http.MethodPost:
		controller.SendMessage(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func roomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller.RegisterRoom(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {

	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/newroom", roomHandler)

}
