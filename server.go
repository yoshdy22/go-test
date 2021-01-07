package main

import (
	"fmt"
	"net/http"
	)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello"))
    })

    fmt.Printf("Start web server\n")
    http.ListenAndServe(":8080", nil)
}
