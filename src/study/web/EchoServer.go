package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var cnt int
var lock sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	lock.Lock()
	cnt++
	lock.Unlock()
	fmt.Fprintf(writer, request.URL.Path)
}

func counter(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Counter %d\n", cnt)
}
