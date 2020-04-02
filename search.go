package wordsearch

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/y0ssar1an/q"
)

var searchList []*SearchItem

func init() {
	searchList = []*SearchItem{
		&SearchItem{
			Name:        "hello",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "goodbye",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "simple",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "list",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "search",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "filter",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "yes",
			SearchCount: 0,
		},
		&SearchItem{
			Name:        "no",
			SearchCount: 0,
		},
	}
}

func incrementCount(item *SearchItem) {
	item.SearchCount++
}

func wordExists(searchTerm string, items []*SearchItem, increment bool) (exists bool) {
	exists = false
	for _, item := range items {
		if item.Name == searchTerm {
			if increment {
				incrementCount(item)
			}
			exists = true
			break
		}
	}

	return exists
}

// SingleWordSearch func
func (sc *SearchClient) SingleWordSearch(ctx context.Context, req *SingleWordSearchRequest) (*SingleWordSearchResponse, error) {
	wordRaw := req.Word
	word := strings.ToLower(wordRaw)
	message := fmt.Sprintf("Sorry, '%v' cannot be found.", wordRaw)
	exists := wordExists(word, searchList, true)

	if exists {
		message = fmt.Sprintf("Yay, '%v' is one of our words.", wordRaw)
	}
	return &SingleWordSearchResponse{
		Message: message,
	}, nil
}

// UpdateSearchList func
func (sc *SearchClient) UpdateSearchList(ctx context.Context, req *UpdateSearchListRequest) (*UpdateSearchListResponse, error) {
	wordRaw := req.Word
	word := strings.ToLower(wordRaw)

	exists := wordExists(word, searchList, false)
	if exists {
		return &UpdateSearchListResponse{
			Message:  fmt.Sprintf("Search term '%v', is already on the list.", wordRaw),
			WordList: searchList,
		}, nil
	}

	newSearchItem := &SearchItem{
		Name:        word,
		SearchCount: 0,
	}

	searchList = append(searchList, newSearchItem)

	return &UpdateSearchListResponse{
		Message:  fmt.Sprintf("New search term '%v', has been added to the list :)", wordRaw),
		WordList: searchList,
	}, nil
}

// TopFiveSearchResults func
func (sc *SearchClient) TopFiveSearchResults(ctx context.Context, req *TopFiveRequest) (*TopFiveResponse, error) {
	sort.Slice(searchList, func(i, j int) bool {
		return searchList[i].SearchCount > searchList[j].SearchCount
	})

	q.Q(searchList)
	return &TopFiveResponse{
		TopFive: searchList,
	}, nil
}
