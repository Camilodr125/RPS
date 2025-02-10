package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
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
	diffValue()
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	diffValue()
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

	//Redirect to another route
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	RenderTemplate(w, "game.html", player)
	
}
func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	diffValue()
	RenderTemplate(w, "about.html", nil)
}

func diffValue()  {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}