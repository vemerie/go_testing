package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)



func main() {

	type Todo struct {
    Title string
    Done  bool
}


type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}


	tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":80", nil)
	
}