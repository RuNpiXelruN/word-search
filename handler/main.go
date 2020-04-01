package main

import (
	"fmt"

	ws "github.com/RuNpiXelruN/word-search"
)

func main() {
	fmt.Println("Hellllooo")

	sc := ws.NewSearchClient()
	sc.Start()
}
