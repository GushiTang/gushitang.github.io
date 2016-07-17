// dev web server
package main

import (
    "fmt"
    "log"
    "net/http"
)

const (
    // INDEX = "index.html"
    // INDEX = "kapu.html"
    INDEX = "return.html"
)

func main() {
    fmt.Println("starting web on http://127.0.0.1:8000")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, INDEX)
    })
    log.Fatal(http.ListenAndServe(":8000", nil))
}
