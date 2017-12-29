package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleServer() {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	res.Body.Close()
	fmt.Printf("%s", resp)
	// Output:{"isServerDown":false}
}
