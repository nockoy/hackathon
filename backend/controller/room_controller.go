package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"
)

func SearchRoom(w http.ResponseWriter, r *http.Request) {

	roomID := r.URL.Query().Get("roomID")

	if roomID == "" {
		log.Println("fail: roomID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.SearchRoom(roomID)

	if err != nil {

		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterRoom(w http.ResponseWriter, r *http.Request) {

	var room model.Rooms

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		log.Println("fail: Error3")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := RegisterRoomCheck(room.Name); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.RegisterRoom(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterRoomCheck(name string) bool {

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
