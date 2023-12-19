package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

var posts = []Article{}
var showPost = Article{}

type Article struct {
	Id                                     uint16
	Title, Anons, Full_text, Image, Author string
}

func articles(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/articles.html",
		"templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//Подключение к БД
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/VladislavsPotapkinsPortfolio")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Выборка данных
	res, err := db.Query("SELECT * FROM `Articles`")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text, &post.Image, &post.Author)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)

	}

	defer res.Close()

	t.ExecuteTemplate(w, "articles", posts)

}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html",
		"templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)

}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	image := r.FormValue("image")
	author := r.FormValue("author")

	if title == "" || anons == "" || full_text == "" || author == "" {
		fmt.Fprintf(w, "Not all data is filled in")
	} else {

		//Подключение к БД
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/VladislavsPotapkinsPortfolio")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		//Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `Articles` (`title`, `anons`, `full_text`,`image`,`author`) VALUES('%s', '%s', '%s','%s', '%s')", title, anons, full_text, image, author))
		if err != nil {
			panic(err)
		}

		defer insert.Close()

		http.Redirect(w, r, "/articles/", http.StatusSeeOther)
	}
}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html",
		"templates/header.html", "templates/footer.html")

	//Подключение к БД
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/VladislavsPotapkinsPortfolio")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Выборка данных
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `Articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	showPost = Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text, &post.Image, &post.Author)
		if err != nil {
			panic(err)
		}

		showPost = post

	}

	defer res.Close()

	t.ExecuteTemplate(w, "show", showPost)

}

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
	r.HandleFunc("/articles/", articles).Methods("GET")
	r.HandleFunc("/articles/{create}/", create).Methods("GET")
	r.HandleFunc("/articles/{create}/{save_article}/", save_article).Methods("POST")
	r.HandleFunc("/articles/{posts}/{id:[0-9]+}/", show_post).Methods("GET")
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleFunc()

}
