package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	conn, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(conn.Body)
	fmt.Println(string(body))
}
