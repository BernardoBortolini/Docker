package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "full Cycle Rocks\n")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
