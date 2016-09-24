package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!</h1>\n")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	fmt.Println("Example app listening at http://localhost:8888")
	fmt.Println("Temp print")
	http.ListenAndServe(":8888", nil)
}
