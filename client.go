package grpcserver

import "sync"

// SearchClient type
type SearchClient struct {
	wg sync.WaitGroup
}

// NewSearchClient func
func NewSearchClient() *SearchClient {
	return &SearchClient{}
}
