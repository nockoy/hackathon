package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"
)

func SearchUser(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.SearchUser(name)
	if err != nil {
		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var u model.Rooms

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Println("fail: decode err")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := RegisterUserCheck(u.Name, u.Email); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.RegisterUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterUserCheck(name string, email string) bool {

	if name == "" {
		log.Println("fail: name is empty")
		return false
	}

	if utf8.RuneCountInString(name) > 50 {
		log.Println("fail: name length is over 50")
		return false
	}

	if email == "" {
		log.Println("fail: email is empty")
		return false
	}

	return true
}
