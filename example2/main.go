package example2

import (
	"fmt"

	bleve "github.com/blevesearch/bleve/v2"
	"github.com/davecgh/go-spew/spew"
)

func Run() {
	indexPath := "markdown_index"
	dirPath := "testdata/markdown"

	bindex := indexMarkdownFiles(indexPath, dirPath)
	defer bindex.Close()

	query := bleve.NewQueryStringQuery("golang")
	searchRequest := bleve.NewSearchRequest(query)
	searchResults, err := bindex.Search(searchRequest)
	if err != nil {
		fmt.Println("Error searching index:", err)
		panic(err)
	}

	fmt.Println("Search Results:", spew.Sdump(searchResults))
}
