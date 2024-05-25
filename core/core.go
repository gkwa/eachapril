package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

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

	// Create a new match query
	matchQuery := bleve.NewMatchQuery("Go")
	matchQuery.SetField("_all")
	searchRequest := bleve.NewSearchRequest(matchQuery)
	searchResults, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Match Query Results:")
	printResults(searchResults)

	// Create a new match phrase query
	phraseQuery := bleve.NewMatchPhraseQuery("Go projects")
	phraseQuery.SetField("_all")
	searchRequest = bleve.NewSearchRequest(phraseQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Match Phrase Query Results:")
	printResults(searchResults)

	// Create a new term range query
	startTerm := "Go"
	endTerm := "Gopher"
	includeStart := true
	includeEnd := true
	termRangeQuery := bleve.NewTermRangeInclusiveQuery(startTerm, endTerm, &includeStart, &includeEnd)
	termRangeQuery.SetField("_all")
	searchRequest = bleve.NewSearchRequest(termRangeQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Term Range Query Results:")
	printResults(searchResults)

	// Create a new date range query
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
	dateRangeQuery := bleve.NewDateRangeQuery(startDate, endDate)
	dateRangeQuery.SetField("_all")
	searchRequest = bleve.NewSearchRequest(dateRangeQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Date Range Query Results:")
	printResults(searchResults)

	// Create a new prefix query
	prefixQuery := bleve.NewPrefixQuery("Go")
	prefixQuery.SetField("_all")
	searchRequest = bleve.NewSearchRequest(prefixQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Prefix Query Results:")
	printResults(searchResults)

	// Create a new wildcard query
	wildcardQuery := bleve.NewWildcardQuery("Go*")
	wildcardQuery.SetField("_all")
	searchRequest = bleve.NewSearchRequest(wildcardQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wildcard Query Results:")
	printResults(searchResults)

	// Create a new fuzzy query
	fuzzyQuery := bleve.NewFuzzyQuery("Goo")
	fuzzyQuery.SetField("_all")
	searchRequest = bleve.NewSearchRequest(fuzzyQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fuzzy Query Results:")
	printResults(searchResults)

	// Create a new boolean query
	booleanQuery := bleve.NewBooleanQuery()
	booleanQuery.AddMust(bleve.NewMatchQuery("Go"))
	booleanQuery.AddMust(bleve.NewMatchQuery("program"))
	searchRequest = bleve.NewSearchRequest(booleanQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Boolean Query Results:")
	printResults(searchResults)

	// Create a new conjunction query
	conjunctionQuery := bleve.NewConjunctionQuery(
		bleve.NewMatchQuery("Go"),
		bleve.NewMatchQuery("programming"),
	)
	searchRequest = bleve.NewSearchRequest(conjunctionQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conjunction Query Results:")
	printResults(searchResults)

	// Create a new disjunction query
	disjunctionQuery := bleve.NewDisjunctionQuery(
		bleve.NewMatchQuery("Go"),
		bleve.NewMatchQuery("Python"),
	)
	searchRequest = bleve.NewSearchRequest(disjunctionQuery)
	searchResults, err = index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disjunction Query Results:")
	printResults(searchResults)
}

func printResults(searchResults *bleve.SearchResult) {
	if len(searchResults.Hits) > 0 {
		for _, hit := range searchResults.Hits {
			fmt.Printf("Document: %s, Score: %f\n", hit.ID, hit.Score)
		}
	} else {
		fmt.Println("No results found")
	}
	fmt.Println()
}
