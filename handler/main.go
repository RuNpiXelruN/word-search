package main

import (
	"fmt"

	grpcserver "github.com/RuNpiXelruN/word-search"
)

func main() {
	fmt.Println("Hellllooo")

	sc := grpcserver.NewSearchClient()
	sc.Start()
}
