package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"uitp/business"
	"uitp/handlers"
)

const (
	defaultPort  = 8080
	staticDir    = "static"
	templatesDir = "templates"
)

var (
	uitpFilename = filepath.Join("static", "Otvety_UITP_2020.html")
)

func main() {
	reader := business.NewUITPReader(uitpFilename)

	hs := handlers.Handlers{
		UITPReader:   reader,
		TemplatesDir: templatesDir,
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	http.HandleFunc("/", hs.Index)
	http.HandleFunc("/search", hs.Search)

	port, err := port()
	if err != nil {
		log.Fatal(err)
	}
	addr := ":" + strconv.Itoa(port)
	log.Println("Starting server " + addr + "...")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func port() (int, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort, nil
	}
	return strconv.Atoi(port)
}
