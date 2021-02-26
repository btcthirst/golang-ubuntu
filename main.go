package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getPots(n string) {

	conn, err := http.Get(n)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(conn.Body)
	fmt.Println(string(body))
}

func main() {
	for i := 1; i < 100; i++ {
		text := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/", i)
		go getPots(text)
	}

	time.Sleep(time.Second)

}
