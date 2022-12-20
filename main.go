package main

import (
	handlers "ascii-art-web/docs/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainPageHandler)
	http.HandleFunc("/ascii", handlers.AsciiPage)
	FileServer := http.FileServer(http.Dir("docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", FileServer))
	fmt.Println("Server is listening to port #8080 ... ")
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
