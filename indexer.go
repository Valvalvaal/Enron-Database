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

// enron_mail_20110402

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./indexer directory")
	}
	dirname := os.Args[1]

	f, err := os.Open("./" + dirname)
	check(err)
	files1, err := f.ReadDir(0)
	check(err)
	fmt.Println(files1[0].Name())
	for _, v := range files1 {
		fmt.Println(v.Name(), v.IsDir())
	}

	r, err := os.Open("./" + dirname + "/" + f.Name())

	files, err := r.Readdir(0)
	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}

}
