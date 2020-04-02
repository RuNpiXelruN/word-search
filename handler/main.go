package main

import (
	grpcserver "github.com/RuNpiXelruN/word-search"
)

func main() {
	grpcSearchClient := grpcserver.NewWordSearchClient()
	grpcserver.Start(grpcSearchClient)
}
