package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getPots(n string, i int) {

	conn, err := http.Get(n)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(conn.Body)

	nameFile := fmt.Sprintf("test%d.txt", i)

	wrFile(nameFile, body)
}

func wrFile(nf string, dt []byte) {
	ioutil.WriteFile(nf, dt, 0777)

}

func main() {
	for i := 1; i < 60; i += 10 {
		text := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/", i)
		go getPots(text, i)
	}

	time.Sleep(time.Second)

}
