package handlers

import (
	"html/template"
	"net/http"
)

func error400(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/400.html")
	if err != nil {
		http.Error(w, "Error 400\nBad request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, http.StatusBadRequest)
}

func error404(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "Error 404\nPage not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, http.StatusNotFound)
}

func error405(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/405.html")
	if err != nil {
		http.Error(w, "Error 405\nMethod Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, http.StatusMethodNotAllowed)
}

func error500(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/500.html")
	if err != nil {
		http.Error(w, "Error 500\nInternal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, http.StatusInternalServerError)
}
