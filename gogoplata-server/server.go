//Build a web server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(write http.ResponseWriter, req *http.Request) {
	fmt.Fprint(write, "<h1>hi</h1>")
}
