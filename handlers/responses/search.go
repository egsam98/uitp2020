package responses

import (
	"html/template"
)

type Search struct {
	Text    string
	Matches []template.HTML
}
