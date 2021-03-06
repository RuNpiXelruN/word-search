package grpcserver

import (
	"context"
	"fmt"
	"sort"
	"strings"

	wordsearch "github.com/RuNpiXelruN/word-search/proto"
)

var searchList []*wordsearch.SearchItem

func init() {

	// Initialising our default
	// set of words to be searched
	searchList = []*wordsearch.SearchItem{
		&wordsearch.SearchItem{
			Name:        "hello",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "goodbye",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "simple",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "list",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "search",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "filter",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "yes",
			SearchCount: 0,
		},
		&wordsearch.SearchItem{
			Name:        "no",
			SearchCount: 0,
		},
	}
}

// IncrementCount func increments the
// SearchCount value of a wordsearch.SearchItem
func (wsc *WordSearchClient) IncrementCount(item *wordsearch.SearchItem) {
	item.SearchCount++
}

// WordExists func determines if a word
// exists in our slice of wordsearch.SearchItem
func (wsc *WordSearchClient) WordExists(searchTerm string, items []*wordsearch.SearchItem, increment bool) (exists bool) {
	exists = false
	for _, item := range items {
		if item.Name == searchTerm {
			if increment {
				wsc.IncrementCount(item)
			}
			exists = true
			break
		}
	}

	return exists
}

// SingleWordSearch func searches for a given
// word and returns a message detailing whether
// it was found within our slice of wordsearch.SearchItem
func (wsc *WordSearchClient) SingleWordSearch(ctx context.Context, req *wordsearch.SingleWordSearchRequest) (*wordsearch.SingleWordSearchResponse, error) {
	wordRaw := req.Word
	word := strings.ToLower(wordRaw)
	message := fmt.Sprintf("Sorry, '%v' cannot be found.", wordRaw)
	exists := wsc.WordExists(word, searchList, true)

	if exists {
		message = fmt.Sprintf("Yay, '%v' is one of our words.", wordRaw)
	}
	return &wordsearch.SingleWordSearchResponse{
		Message: message,
	}, nil
}

// UpdateSearchList func updates our slice
// of wordsearch.SearchItem if the given
// word doens't already exist.
func (wsc *WordSearchClient) UpdateSearchList(ctx context.Context, req *wordsearch.UpdateSearchListRequest) (*wordsearch.UpdateSearchListResponse, error) {
	wordRaw := req.Word
	word := strings.ToLower(wordRaw)

	exists := wsc.WordExists(word, searchList, false)
	if exists {
		return &wordsearch.UpdateSearchListResponse{
			Message:  fmt.Sprintf("Search term '%v', is already on the list.", wordRaw),
			WordList: searchList,
		}, nil
	}

	newSearchItem := &wordsearch.SearchItem{
		Name:        word,
		SearchCount: 0,
	}

	searchList = append(searchList, newSearchItem)

	return &wordsearch.UpdateSearchListResponse{
		Message:  fmt.Sprintf("New search term '%v', has been added to the list :)", wordRaw),
		WordList: searchList,
	}, nil
}

// TopFiveSearchResults func returns the
// top 5 searched words from our slice of
// wordsearch.SearchItem, ordered by highest
// SearchCount number.
func (wsc *WordSearchClient) TopFiveSearchResults(ctx context.Context, req *wordsearch.TopFiveRequest) (*wordsearch.TopFiveResponse, error) {
	sort.Slice(searchList, func(i, j int) bool {
		return searchList[i].SearchCount > searchList[j].SearchCount
	})

	topList := searchList[:5]
	return &wordsearch.TopFiveResponse{
		TopFive: topList,
	}, nil
}
