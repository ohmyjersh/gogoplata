package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get(`http://localhost:8080`)
	if err != nil {
		fmt.Println(`error`)
	}
	defer resp.Body.Close()
	body, parseErr := ioutil.ReadAll(resp.Body)
	if parseErr != nil {
		fmt.Println(`error`)
	}
	fmt.Printf(`%s`, body)
}
