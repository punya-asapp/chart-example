package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	message := flag.String("message", "Hello, World!", "Message served by this server")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, *message)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
