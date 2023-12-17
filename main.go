package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

func TicTacToe(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/TicTacToe.html",
		"templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "TicTacToe", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html",
		"templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func HandleFunc() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/Tic-Tac-Toe/", TicTacToe).Methods("GET")
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleFunc()

}
