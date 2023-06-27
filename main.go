package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my wonderful site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<h1>Contact Page</h1>
	<p>
		To get in touch, email me at <a href="mailto:abc@gmail.com">abc@gmail.com</a>
	</p>
	`)
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

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(notfoundHandler)

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
