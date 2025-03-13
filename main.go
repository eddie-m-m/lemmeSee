package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setContentTypeHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	galleryID := chi.URLParam(r, "galleryID")
	setContentTypeHeader(w)
	fmt.Fprint(w, "<h1>Gallery ID: "+galleryID+"</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeHeader(w)
	executeTemplate(w, "faq.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeHeader(w)
	executeTemplate(w, "contact.gohtml")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeHeader(w)
	executeTemplate(w, "home.gohtml")
}

func executeTemplate(w http.ResponseWriter, filePath string) {
	tplPath := filepath.Join("templates", filePath)
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{galleryID}", galleryHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
