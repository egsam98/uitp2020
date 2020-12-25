package business

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type UITPReader struct {
	filename string
}

func NewUITPReader(filename string) *UITPReader {
	return &UITPReader{filename: filename}
}

func (r *UITPReader) SearchQuestion(substring string) ([]string, error) {
	f, err := os.Open(r.filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open "+r.filename)
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file by goquery")
	}

	h1 := doc.Find("h1")
	matches := make([]string, 0, h1.Length())
	h1.Each(func(_ int, s *goquery.Selection) {
		if strings.Contains(strings.ToLower(s.Text()), strings.ToLower(substring)) {
			question, _ := s.Html()
			answer, _ := s.Next().Html()
			matches = append(matches, question+answer)
		}
	})
	return matches, nil
}
