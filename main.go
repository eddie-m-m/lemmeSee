package main

import (
	"fmt"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>Wanna connect? email me at
	 <a href="mailto:eddie.m.menefee@gmail.com">eddie.m.menefee@gmail.com</p>`)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Test 2!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
