package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Player struct {
	Name string
}

var player Player
const (
	templateDir = "template/"
	templateBase = templateDir + "base.html"
)

func RenderTemplate(w http.ResponseWriter,  page string, data any)  {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir + page))
	

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error:", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "new-game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return 
		}

		player.Name = r.Form.Get("name")
	}
	RenderTemplate(w, "game.html", player)
	
}
func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Play")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html", nil)
}