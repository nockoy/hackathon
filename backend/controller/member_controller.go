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
