package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

func _1_7_webSev() {
	http.HandleFunc("/", handler_normal)
	http.HandleFunc("/GetWebSite", handler)
	http.HandleFunc("/count", handler_counter)
	http.HandleFunc("/gif", handler_gif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler_gif(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("handler", count)
	if len(os.Args) == 2 {
		url := os.Args[1]
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			fmt.Fprintf(w, "url.path = %q \n", r.URL.Path)
			return
		}
		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s:%v\n", url, err)
			fmt.Fprintf(w, "url.path = %q \n", r.URL.Path)
			return
		}
		fmt.Fprintf(w, "%s", b)
	} else {
		fmt.Fprintf(w, "url.path = %q \n", r.URL.Path)
	}
}

func handler_normal(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	//fmt.Fprintf(w, "url.path = %q \n", r.URL.Path)
	fmt.Println("handler_normal", count)
}

func handler_counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("handler_counter", count)
	fmt.Fprintf(w, "count = %d \n", count)
}
