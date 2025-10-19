package power4

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func loadTemplates(files ...string) (*template.Template, error) {
	return template.ParseFiles(
		append(files,
			"templates/start.html",
			"templates/tab.html",
			"templates/result.html",
		)...,
	)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var pageFile string
	if !game.Started {
		pageFile = "pages/home.html"
	} else if game.Winner != 0 {
		pageFile = "templates/result.html"
	} else {
		pageFile = "pages/play.html"
	}

	tmpl, err := loadTemplates(pageFile)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	name := filepath.Base(pageFile)
	if name == "result.html" {
		name = "result"
	}
	if err := tmpl.ExecuteTemplate(w, name, game); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	p1 := r.FormValue("p1")
	p2 := r.FormValue("p2")
	diff := r.FormValue("difficulty")

	mu.Lock()
	game = newGame(diff, p1, p2)
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	col, err := strconv.Atoi(r.FormValue("col"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if !game.Started || game.Winner != 0 || col < 0 || col >= game.Cols {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if !dropToken(&game, col, game.Current) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	game.Moves++
	if checkWin(game.Grid, game.Rows, game.Cols, game.Current) {
		game.Winner = game.Current
	} else if game.Moves == game.Rows*game.Cols {
		game.Winner = 3
	} else {
		if game.Current == 1 {
			game.Current = 2
		} else {
			game.Current = 1
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	game = Game{}
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func replayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	mu.Lock()
	prev := game
	game = newGame(prev.Difficulty, prev.Player1, prev.Player2)
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
