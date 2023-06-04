package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"
)

func ChannelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}

	switch r.Method {
	case http.MethodGet:
		SearchJoinChannelsByUserID(w, r)
	case http.MethodPost:
		RegisterChannel(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func NotJoinChannelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	switch r.Method {
	case http.MethodGet:
		SearchNotJoinChannelsByUserID(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func SearchJoinChannelsByUserID(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		log.Println("fail: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.GetJoinChannelsByUserID(userID)

	if err != nil {

		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func SearchNotJoinChannelsByUserID(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		log.Println("fail: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.GetNotJoinChannelsByUserID(userID)

	if err != nil {

		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterChannel(w http.ResponseWriter, r *http.Request) {

	var channel model.Channels

	if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
		log.Println("fail: Error3")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := RegisterChannelCheck(channel.Name); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.RegisterChannel(channel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterChannelCheck(name string) bool {

	if name == "" {
		log.Println("fail: name is empty")
		return false
	}

	if utf8.RuneCountInString(name) > 16 {
		log.Println("fail: name length is over 16")
		return false
	}

	return true
}
