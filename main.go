package main

import (
	"fmt"

	root "github.com/gary23w/metasearch_api/cmd/root"
)

func main() {
	fmt.Println("[*] Search API CLI interface")
	root.Execute()
}
