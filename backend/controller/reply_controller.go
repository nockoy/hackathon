package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func SendReply(w http.ResponseWriter, r *http.Request) {

	var re model.Replies

	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		log.Println("fail: decode error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := SendMessageCheck(re.Text); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.SendReply(re)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func ReplyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	switch r.Method {
	case http.MethodPost:
		SendReply(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
