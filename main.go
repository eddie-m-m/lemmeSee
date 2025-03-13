package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setContentTypeHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// http.NotFound(w, r)
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	galleryID := chi.URLParam(r, "galleryID")
	setContentTypeHeader(w)
	fmt.Fprint(w, "<h1>Gallery ID: "+galleryID+"</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeHeader(w)
	fmt.Fprint(w, `<h1>FAQ</h1>`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeHeader(w)
	fmt.Fprint(w, `<h1>Contact Page</h1><p>Wanna connect? email me at
	 <a href="mailto:eddie.m.menefee@gmail.com">eddie.m.menefee@gmail.com</p>`)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeHeader(w)
	fmt.Fprint(w, "<h1>Test 2!</h1>")
}

func main() {
	// http.HandleFunc("/", pathHandler)
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
