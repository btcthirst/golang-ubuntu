package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func getPots(n string, i int) {

	conn, err := http.Get(n)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(conn.Body)

	nameFile := fmt.Sprintf("./posts/test%d.txt", i)

	wrFile(nameFile, body)
}

func wrFile(nf string, dt []byte) {
	f, err := os.Create(nf)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	writer := bufio.NewWriter(f)
	writer.WriteString(string(dt))
	writer.Flush()

}

func main() {
	for i := 1; i < 6; i++ {
		text := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/", i)
		go getPots(text, i)
	}

	time.Sleep(time.Second)

}
