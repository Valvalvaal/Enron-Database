package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./indexer directory")
	}
	dat, err := os.ReadFile("")
	fmt.Println(dat, err)
}
