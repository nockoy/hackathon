package controller

import (
	"db/usecase"
	"log"
	"net/http"
)

type MessageGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
