package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

// Docs
// https://golang.org/pkg/net/http
// https://golang.org/pkg/io/#Writer

// This is our function we are going to use to handle the request
// All handlers need to accept two arguments
// 1. Is the ResponseWriter interface, we use this to write a reponse back to the client
// 2. Is the Reponse struct which holds useful information about the request headers, method, url etc
func hello(w http.ResponseWriter, r *http.Request) {
	// log-it
	fmt.Printf("[%s] from %s \n", r.Method, r.RemoteAddr)

	// - Create some simple validation for the name parameter with the following rules
	//    1. Check the name is present. Return an error message informing the client if not
	//    2. The name must greater the one character long
	var requestName = r.FormValue("name")

	switch {
	case requestName == "":
		issueResponse(w, http.StatusBadRequest, "The 'name' parameter is required")
	case len(requestName) < 2:
		issueResponse(w, http.StatusBadRequest, "The 'name' parameter must be greater than one character")
	default:
		issueResponse(w, http.StatusOK, fmt.Sprintf("%s %s", "hello", requestName))
	}
}

func issueResponse(w http.ResponseWriter, responseCode int, responseResult string) {
	res := Response{
		Code:   responseCode,
		Result: responseResult,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// port to listen on
var port string

func main() {
	fmt.Println("--------------------------------------")
	fmt.Println("Running server on localhost:" + port)
	fmt.Println("--------------------------------------")

	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	fmt.Println("Waiting for connections ....")
	http.ListenAndServe(":"+port, nil)
}

func init() {
	flag.StringVar(&port, "port", "8000", "port to run on")
	flag.Parse()
}
