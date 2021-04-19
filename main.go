package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"stock-checker/checker"
	"time"
)

func main() {
	c := flag.String("c", "products.json", "Specify the products configuration file.")
	flag.Parse()

	f, err := ioutil.ReadFile(*c)
	if err != nil {
		panic(err)
	}

	var products []checker.Product
	err = json.Unmarshal(f, &products)
	if err != nil {
		panic(err)
	}

	ch := checker.Checker{
		Products:          &products,
		HttpClient:        http.DefaultClient,
		WaitBetweenChecks: time.Minute,
	}
	ch.Run()
}
