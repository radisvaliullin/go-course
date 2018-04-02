package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://localhost:7373")
	if err != nil {
		log.Fatal("http: get, err - ", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("http: get body, err - ; body - %s", err, string(b))
	}

	log.Print("http: res body - ", string(b))
}
