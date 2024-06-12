package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Printf("Server is listening on port %s...\n")
}

func homeHandler() {

}
