package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {}

func (h *UserHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	// main.goのAppHandlerで一度ShiftPathしているので headが空 = /users と同義
	if head == "" && r.Method == "GET"{
		GetUsers(w)
		return
	}

	id, err := strconv.ParseInt(head, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid params. head: %s", head), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		GetUser(w, id)
		return
	}

}

func GetUsers(w http.ResponseWriter){
	w.Write([]byte("get users"))
}

func GetUser(w http.ResponseWriter,id int64){
	w.Write([]byte(fmt.Sprintf("get user. id: %d", id)))
}


