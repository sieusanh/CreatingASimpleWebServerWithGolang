package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	//"html"
)

var counter int
var mutex = &sync.Mutex{}
//var mutex sync.Mutex

func echoString(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	*/

	/*
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	*/

	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})
	*/

	// cách làm này sẽ bỏ qua "/static" trong request link của user
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//log.Fatal(http.ListenAndServe(":8081", nil))
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt","server.key", nil))

}