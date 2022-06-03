package main

import (
	"log"
	"html/template"
	"net/http"
)

const PORT string = ":3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving request '/'")
		render(w, "test.page.gohtml")
	})
	
	log.Printf("Starting frontend service on port %v", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	templatesDir := "./cmd/web/templates/"
	partials := []string{
		templatesDir + "base.layout.gohtml",
		templatesDir + "header.partial.gohtml",
		templatesDir + "footer.partial.gohtml",
	}
	templates := []string{
		templatesDir + t,
	}
	templates = append(templates, partials...)
	
	templ, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := templ.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

