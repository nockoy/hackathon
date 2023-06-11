package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func GetReplies(w http.ResponseWriter, r *http.Request) {

	replyToId := r.URL.Query().Get("reply_to_id")

	bytes, err := usecase.GetReplies(replyToId)
	if err != nil {
		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

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

func EditReply(w http.ResponseWriter, r *http.Request) {

	var re model.Replies

	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("fail: Decode error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := usecase.EditReply(re)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func DeleteReply(w http.ResponseWriter, r *http.Request) {

	var re model.Replies

	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("fail: Decode error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := usecase.DeleteReply(re)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
