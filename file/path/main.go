package main

import (
	"fmt"
	"log"
	"os"
)

/*
	How to get the directory of the currently running file?
	In nodejs I use __dirname . What is the equivalent of this in Golang?
	ref: https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
*/
func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
