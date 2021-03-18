package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	printFile(os.Args[1])
}

func printFile(name string) {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
