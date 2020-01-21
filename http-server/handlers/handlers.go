package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := []byte("<h1>Hello</h1>")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}

func NumberHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		data := []byte(fmt.Sprintf("<h1>Page %s</h1>", params["pageNumber"]))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}

func UserHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type User struct {
			Name    string
			Company string
		}
		res, _ := json.Marshal(&User{Name: "Ravi", Company: "TW"})
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(res)
	}
}
