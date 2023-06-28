package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles(filepath)

	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error executing template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, templatePath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<h3>Q: Is there a free version?</h3>
	<p>A: Yes! We offer a free trial for 30 days on any paid plan.</p>
	<br>
	<h3>Q: What are your support hours?</h3>
	<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends</p>
	<br>
	<h3>Q: How do I contact support?</h3>
	<p>A: Email us - support@lenslocked.com</p>
	`)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusNotFound)
	//fmt.Fprint(w, `<h1>Page Not Found</h1>`)
	http.Error(w, `Page Not Found`, http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	// adding middleware logger
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(notfoundHandler)

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
