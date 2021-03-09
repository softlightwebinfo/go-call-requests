package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func callXhr(server string, ch chan<- string, i int) {
	start := time.Now()
	server = fmt.Sprintf("%s?call=%d", server, i)
	resp, err := http.Get(server)
	if err != nil {
		fmt.Println("ERROR Call", err)
		return
	}
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s and last routine %d", secs, len(body), server, i)
}

func run(server string, numRequests int) {
	start := time.Now()
	ch := make(chan string)

	for i := 0; i < numRequests; i++ {
		go callXhr(server, ch, i)
	}

	for i := 0; i < numRequests; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	if len(os.Args) < 3 {
		_ = fmt.Errorf("Error numbers arguments")
		return
	}

	requests, err := strconv.Atoi(os.Args[1])
	if err != nil {
		_ = fmt.Errorf("Error: first parameter is not number")
		return
	}

	run(os.Args[2], requests)
}
