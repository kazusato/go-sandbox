package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	url := "http://www.yahoo.co.jp/"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
