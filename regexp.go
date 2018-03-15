package main

import (
	"regexp"
	"fmt"
)

func main()  {

	str := "Hello Japanese World!"

	ptn := regexp.MustCompile(`[A-Z][a-z]*`)
	results := ptn.FindAllStringSubmatch(str, -1)

	for _, value := range results {
		fmt.Println(value)
	}

}
