package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func InternalError(w http.ResponseWriter, err error) {
	log.Printf("%+v\n", err)
	w.WriteHeader(http.StatusInternalServerError)
}

func BadRequestError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func Template(w http.ResponseWriter, path string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		InternalError(w, errors.Wrap(err, "failed to create template"))
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		InternalError(w, errors.Wrap(err, "failed to execute template"))
	}
}
