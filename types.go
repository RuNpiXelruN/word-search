package grpcserver

import (
	"context"

	wordsearch "github.com/RuNpiXelruN/word-search/proto"
)

//go:generate moq -out mock_wordsearchservice.go . WordSearchService

// WordSearchService type
type WordSearchService interface {
	SingleWordSearch(ctx context.Context, req *wordsearch.SingleWordSearchRequest) (*wordsearch.SingleWordSearchResponse, error)
	UpdateSearchList(ctx context.Context, req *wordsearch.UpdateSearchListRequest) (*wordsearch.UpdateSearchListResponse, error)
	TopFiveSearchResults(ctx context.Context, req *wordsearch.TopFiveRequest) (*wordsearch.TopFiveResponse, error)
	IncrementCount(item *wordsearch.SearchItem)
	WordExists(searchTerm string, items []*wordsearch.SearchItem, increment bool) (exists bool)
	StartGRPC() error
	StartRest() error
}

// WordSearchClient type
type WordSearchClient struct{}

// NewWordSearchClient func
func NewWordSearchClient() *WordSearchClient {
	return &WordSearchClient{}
}
