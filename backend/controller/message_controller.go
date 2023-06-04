package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"unicode/utf8"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {

	var m model.Messages

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Println("fail: Error2")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := SendMessageCheck(m.Text); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.SendMessage(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func SendMessageCheck(text string) bool {

	if text == "" {
		log.Println("fail: text is empty")
		return false
	}

	if utf8.RuneCountInString(text) > 500 {
		log.Println("fail: name length is over 500")
		return false
	}

	return true
}

func GetMessage(w http.ResponseWriter, r *http.Request) {

	roomId := r.URL.Query().Get("room_id")

	bytes, err := usecase.GetMessage(roomId)
	if err != nil {
		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	switch r.Method {
	case http.MethodGet:
		GetMessage(w, r)
	case http.MethodPost:
		SendMessage(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
