package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func RegisterMember(w http.ResponseWriter, r *http.Request) {

	var m model.Members

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Println("fail: Error5")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := usecase.RegisterMember(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func DeleteMember(w http.ResponseWriter, r *http.Request) {

	var m model.Members

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Println("fail: decode error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := usecase.DeleteMember(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func MemberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	switch r.Method {
	case http.MethodPost:
		RegisterMember(w, r)
	case http.MethodDelete:
		DeleteMember(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
