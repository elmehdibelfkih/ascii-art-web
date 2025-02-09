package main

import (
	"fmt"
	"net/http"
	"os"

	"ascii-art-web/internal"
)

func main() {
	internal.CreateTemplates()
	http.HandleFunc("/static/", internal.StaticHandler)
	http.HandleFunc("/", internal.RootHandler)
	http.HandleFunc("/ascii-art", internal.AsciiArtHandler)
	http.HandleFunc("/export", internal.ExportHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
