package example2

import (
	"fmt"
	"os"
	"path/filepath"

	bleve "github.com/blevesearch/bleve/v2"
)

func indexMarkdownFiles(indexPath, dirPath string) bleve.Index {
	defaultIndexMapping := bleve.NewIndexMapping()
	bindex, err := bleve.New(indexPath, defaultIndexMapping)
	if err != nil {
		if err == bleve.ErrorIndexPathExists {
			bindex, err = bleve.Open(indexPath)
			if err != nil {
				fmt.Println("Error opening existing index:", err)
				panic(err)
			}
		} else {
			fmt.Println("Error creating index:", err)
			panic(err)
		}
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
			Path:     path,
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

func Index(forceIndex bool) {
	indexPath := "markdown_index"
	dirPath := "markdown_files"

	if forceIndex {
		err := os.RemoveAll(indexPath)
		if err != nil {
			fmt.Println("Error removing existing index:", err)
			panic(err)
		}
	}

	indexMarkdownFiles(indexPath, dirPath)
	fmt.Println("Indexing completed.")
}
