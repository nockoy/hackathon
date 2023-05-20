package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

type UserResForHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var u model.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Println("fail: Error1")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := RegisterUserCheck(u.Name, u.Age); isOk != true {
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

func RegisterUserCheck(name string, age int) bool {

	if name == "" {
		log.Println("fail: name is empty")
		return false
	}

	if utf8.RuneCountInString(name) > 50 {
		log.Println("fail: name length is over 50")
		return false
	}

	if age < 20 || age > 80 {
		log.Println("fail: age is not between 20-80")
		return false
	}

	return true
}
