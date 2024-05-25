package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve/v2"
)

func Run() {
	mapping := bleve.NewIndexMapping()

	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	markdownDir := "markdown"
	err = filepath.Walk(markdownDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			err = index.Index(path, string(content))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	query1 := bleve.NewMatchQuery("Go")
	query1.SetField("_all")
	query2 := bleve.NewMatchQuery("projects")
	query2.SetField("_all")

	conjunctionQuery := bleve.NewConjunctionQuery(query1, query2)

	searchRequest := bleve.NewSearchRequest(conjunctionQuery)

	searchResults, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Search Results:")
	if len(searchResults.Hits) > 0 {
		for _, hit := range searchResults.Hits {
			fmt.Printf("Document: %s, Score: %f\n", hit.ID, hit.Score)
		}
	} else {
		fmt.Println("No results found")
	}
}
