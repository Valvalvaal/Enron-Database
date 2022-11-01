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
	dirname := os.Args[1]
	f, err := os.Open("./" + dirname)
	fmt.Println(f.Name(), err)

	files, err := f.Readdir(0)
	fmt.Println(files, err)
}
