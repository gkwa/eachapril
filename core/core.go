package core

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
)

func Run() {
	mapping := bleve.NewIndexMapping()

	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	data := struct {
		Title string
		Body  string
	}{
		Title: "Hello World",
		Body:  "This is a sample document.",
	}

	err = index.Index("doc1", data)
	if err != nil {
		log.Fatal(err)
	}

	query := bleve.NewMatchQuery("hello")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(searchResults)
}
