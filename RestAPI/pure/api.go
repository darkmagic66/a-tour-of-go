package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "testing/iotest"
)

type Movie struct {
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

func movieHandler(res http.ResponseWriter, req *http.Request) {
	method := req.Method
	fmt.Fprintf(res, "%s Movie", method)
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(res, "error : %v", err)
			return
		}
		fmt.Println(string(body))
		obj := Movie{}
		err = json.Unmarshal(body, &obj)
		return
	}
}

func main() {
	fmt.Println("server start soon")
	http.HandleFunc("/movie", movieHandler)

	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}
