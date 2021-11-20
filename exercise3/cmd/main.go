package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/quangduc/exercise3/models"
)

func main() {
	port := flag.Int("port", 3030, "port for the web application")
	filename := flag.String("file", "gopher.json", "the json file")
	flag.Parse()
	fmt.Printf("using the story from the file %s\n", *filename)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := models.JsonStory(f)

	if err != nil {
		panic(err)
	}

	handler := models.NewHandler(story)
	fmt.Printf("starting a new server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story./"):]
}
