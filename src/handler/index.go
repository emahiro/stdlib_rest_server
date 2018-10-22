package handler

import "net/http"

type RootHandler struct {}

func (h *RootHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	if head == "" && r.Method == "GET" {
		Index(w)
	}
}

func Index(w http.ResponseWriter) {
	w.Write([]byte("root handler"))
}
