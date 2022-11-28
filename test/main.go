package main

import (
	"flag"
	"fmt"
)

var flagStr string

func main() {
	flag.StringVar(&flagStr, "str", "default", "this is an important value")
	flag.Parse()
	fmt.Println(flagStr)
}
