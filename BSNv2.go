package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var text string
	textArg := flag.String("nummer", "*", "Gebruik -nummer en vul BSN in")
	flag.Parse()
	// als flag niet is gebruikt
	if fmt.Sprintf("%s", *textArg) == "" {
		fmt.Print("Voer nummer in: ")
		fmt.Scanln(&text)
	} else {
		text = fmt.Sprintf("%s", *textArg)
	}
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	fmt.Println(strings.Contains(s, *textArg))
}
