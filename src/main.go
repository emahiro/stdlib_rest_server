package main

import (
	"fmt"
	"net/http"

	"handler"
	mw "middleware"
)

var port = 5000

type AppHandler struct {
	RootHandler *handler.RootHandler
	UserHandler *handler.UserHandler
}

func(h *AppHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	var head string
	head, r.URL.Path = handler.ShiftPath(r.URL.Path)

	switch head {
	case "":
		h.RootHandler.ServeHTTP(w, r)
		return
	case "users":
		h.UserHandler.ServeHTTP(w, r)
		return
	default:
		http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func main() {
	fmt.Println("Server Start....")
	h := &AppHandler{
		UserHandler: &handler.UserHandler{},
	}

	router := mw.RequestLog(h)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		panic(err)
	}
}
