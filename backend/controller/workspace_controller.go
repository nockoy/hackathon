package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"
)

func WorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}

	switch r.Method {
	case http.MethodPost:
		RegisterWorkspace(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func RegisterWorkspace(w http.ResponseWriter, r *http.Request) {

	var workspaces model.Workspaces

	if err := json.NewDecoder(r.Body).Decode(&workspaces); err != nil {
		log.Println("fail: Error4")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := RegisterWorkspaceCheck(workspaces.Name); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.RegisterWorkspace(workspaces)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterWorkspaceCheck(name string) bool {

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
