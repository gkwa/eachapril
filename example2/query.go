package example2

import (
	"fmt"

	bleve "github.com/blevesearch/bleve/v2"
)

func Query() {
	indexPath := "markdown_index"

	bindex, err := bleve.Open(indexPath)
	if err != nil {
		fmt.Println("Error opening index:", err)
		panic(err)
	}
	defer bindex.Close()

	query := bleve.NewQueryStringQuery("golang")
	searchRequest := bleve.NewSearchRequest(query)
	searchResults, err := bindex.Search(searchRequest)
	if err != nil {
		fmt.Println("Error searching index:", err)
		panic(err)
	}

	fmt.Printf("Search Results: (%d matches)\n", searchResults.Total)
	for _, hit := range searchResults.Hits {
		file, ok := hit.Fields["Path"].(string)
		if ok {
			fmt.Printf("File: %s, Score: %f\n", file, hit.Score)
		}
	}
}
