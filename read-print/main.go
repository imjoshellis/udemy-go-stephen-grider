package main

import (
	"io"
	"os"
)

func main() {
	dat, err := os.Open(os.Args[1])
	check(err)
	io.Copy(os.Stdout, dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
