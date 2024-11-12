package route

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html")
}

func Projects(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "projects.html")
}

func Skills(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "skills.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact.html")
}
