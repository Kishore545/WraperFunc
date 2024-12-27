package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile(Filename string) ([]byte, error) {
	x, err := ioutil.ReadFile(Filename)
	if err != nil {
		return nil, fmt.Errorf("error in ReadFile func %s", err)
	}
	return x, nil
}

func main() {
	x, err := readFile("temp.txt")
	if err != nil {
		log.Fatalf("Read File in main : %s", err)
	}
	fmt.Println(x)
	fmt.Println(string(x))
}
