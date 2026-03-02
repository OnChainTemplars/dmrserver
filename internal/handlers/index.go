package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"

	"github.com/pd0mz/go-dmr/homebrew"
)

//go:embed index.html
var templateHTML string

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/index.html")

	h := r.Context().Value("homebrew").(*homebrew.Homebrew)

	data := struct {
		Callsign string
	}{
		Callsign: h.Config.Callsign,
	}

	t := template.Must(template.New("index").Parse(templateHTML))
	t.Execute(w, data)
}
