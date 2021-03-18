package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "字符分隔")
var line = flag.Bool("n", true, "换行")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *line {
		fmt.Println()
	}
}
