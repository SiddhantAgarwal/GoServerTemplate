package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func testRoute(route string) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()
	res, err := http.Get(ts.URL + route)
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
}

func ExampleIndex() {
	testRoute("")
	// Output:{"isServerDown":false}
}

func ExampleRouteNotFound() {
	testRoute("/unknown")
	// Output:{"Error": "Not Found"}
}
