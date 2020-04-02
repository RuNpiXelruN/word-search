// !/bin/sh main

package main

import (
	grpcserver "github.com/RuNpiXelruN/word-search"
)

func main() {

	// create a new reference to a grpcserver.WordSearchClient
	// which implements the  grpcserver.WordSearchService interface
	// and passed to the Start() function.
	grpcSearchClient := grpcserver.NewWordSearchClient()
	grpcserver.Start(grpcSearchClient)
}
