package main

import (
	"ascii-art-web/internal"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", internal.RootHandler)
	http.HandleFunc("/ascii-art", internal.AsciiArtHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
