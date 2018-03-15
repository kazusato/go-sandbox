package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fp, err := os.Open("input_ascii.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(">> " + scanner.Text())
	}
	fp.Close()
}
