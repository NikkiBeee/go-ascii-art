package main

import (
    "fmt"
    // "os"
	"log"
	"io/ioutil"
	// "strconv"
	// "reflect"
)



func main() {
	i, err := ioutil.ReadFile("../AsciiArt/RubberDuck.txt")
	sz := len(i)
	
	
	
	if err != nil {
		log.Fatal(err)
	}

	l := 0 
	for  l < sz {
		fmt.Print(string(i[l]))
		l++
	}
	
}
