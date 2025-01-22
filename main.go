package main

import (
	// "io"
	"fmt"
	"net/http"
	// "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("is 8080 serv now")
	http.ListenAndServe(":8080", nil)

}
