package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/pkg/errors"

	"uitp/handlers/responses"
	"uitp/utils/web"
)

func (hs *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := filepath.Join(hs.TemplatesDir, "index.html")
	web.Template(w, path, nil)
}

func (hs *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		web.BadRequestError(w, errors.New("\"text\" form param must be provided"))
		return
	}

	matches, err := hs.UITPReader.SearchQuestion(text)
	if err != nil {
		web.InternalError(w, err)
		return
	}

	var matchesTmpls []template.HTML
	if len(matches) == 0 {
		matchesTmpls = []template.HTML{"По запросу ничего не найдено"}
	} else {
		matchesTmpls = make([]template.HTML, len(matches))
		for i, match := range matches {
			matchesTmpls[i] = template.HTML(match)
		}
	}
	web.Template(w, filepath.Join(hs.TemplatesDir, "index.html"), responses.Search{
		Text:    text,
		Matches: matchesTmpls,
	})
}
