package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = 8080

func indexHandler(w http.ResponseWriter, r *http.Request) {
    for _, envvar := range os.Environ() {
        w.Write([]byte(fmt.Sprintf("%s\n", envvar)))
    }
}

func main() {
    for _, e := range os.Environ() {
        fmt.Println(e)
    }
    http.HandleFunc("/", indexHandler)
    log.Printf("Starting HTTP server at port: %d\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
