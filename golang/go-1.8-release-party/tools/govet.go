package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://golang.xxx")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Resp: %s", res.Status)
}
