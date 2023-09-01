package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	
)

type Film struct {
  Title string
  Director string
}

func main() {
  fmt.Println("hello world")

  h1 := func (w http.ResponseWriter, r *http.Request)  {
    tmp1 := template.Must(template.ParseFiles("index.html"))
    films := map[string][]Film{
      "Films" : {
        {Title: "The Godfather", Director: "Francis Ford Coppola"},
        {Title: "Blade Runner", Director: "Ridley Scott"},
        {Title: "The Thing", Director: "John Carpenter"},
      },
    }
    tmp1.Execute(w, films)
  }
  h2 := func (w http.ResponseWriter, r *http.Request)  {
    title := r.PostFormValue("title")  
    director := r.PostFormValue("director")
    fmt.Println(title)
    fmt.Println(director)
  }

  http.HandleFunc("/",h1)
  http.HandleFunc("/add-film/",h2)

  log.Fatal(http.ListenAndServe(":8000", nil))

}
