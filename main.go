package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ip4, err := http.Get("http://check-host.net/ip")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(ip4.Body)
	text := string(body)
	fmt.Println(text)
}
