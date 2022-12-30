package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title":   "Personal web",
	"IsLogin": false,
}

// Array of objects
// nama = []string{"Abel", "Dandi", "Ilham", "Jody"}

// This is interface
// type persegi interface {
// 	panjang() float64
// 	lebar() float64
// }

// type Project struct {
// 	Title        string
// 	date_start   string
// 	date_end     string
// 	Description  string
// 	technologies []string
// 	duration     string
// }

// var Projects = []Project{
// 	{
// 		Title:        "Dumbways Mobile App 2022",
// 		date_start:   "1 Des 2022",
// 		date_end:     "9 Des 2022",
// 		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
// 		technologies: []string{"nodejs", "nextjs", "reactjs", "Typescript"},
// 		duration:     " 1 Minggu",
// 	},
// }

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", index).Methods("GET")
	route.HandleFunc("/project", projectForm).Methods("GET")
	route.HandleFunc("/project/{id}", projectDetail).Methods("GET")
	route.HandleFunc("/project", projectAdd).Methods("POST")
	route.HandleFunc("/contact", contactMe).Methods("GET")

	// port := 5000
	fmt.Println("Server running at localhost:5000")
	http.ListenAndServe("localhost:5000", route)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// var tmpl, err = template.ParseFiles("views/index.html")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func projectForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/project-add.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	respData := map[string]interface{}{
		"Data": Data,
		"Id":   id,
		// "Projects": Projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func projectAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	// techstack := r.Form["project-tech"]
	fmt.Println("Name :" + r.PostForm.Get("project-name"))
	fmt.Println("Start :" + r.PostForm.Get("project-start"))
	fmt.Println("End :" + r.PostForm.Get("project-end"))
	fmt.Println("Description :" + r.PostForm.Get("project-description"))
	fmt.Println("Tech Stack :", r.Form["project-tech"])

	http.Redirect(w, r, "/project", http.StatusMovedPermanently)
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}
