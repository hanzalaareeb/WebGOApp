package handlers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// I dont know that I am going to do
//	  err != nil {
//	  error(w, err.Error)
//	  }

/*
	func byeHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "bye World!")
	}
*/
