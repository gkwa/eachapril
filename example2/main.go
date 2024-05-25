package example2

import (
	"fmt"
	"os"
	"path/filepath"

	bleve "github.com/blevesearch/bleve/v2"
	"github.com/davecgh/go-spew/spew"
)

type MarkdownFile struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

func indexMarkdownFiles(indexPath, dirPath string) bleve.Index {
	defaultIndexMapping := bleve.NewIndexMapping()
	bindex, err := bleve.New(indexPath, defaultIndexMapping)
	if err != nil {
		fmt.Println("Error creating index:", err)
		panic(err)
	}

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		mdFile := MarkdownFile{
			Title:    info.Name(),
			Category: filepath.Base(filepath.Dir(path)),
			Content:  string(content),
		}

		err = bindex.Index(path, mdFile)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error indexing markdown files:", err)
		panic(err)
	}

	return bindex
}

func Run() {
	indexPath := "markdown_index"
	dirPath := "markdown_files"

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
