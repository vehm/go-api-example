package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// fs := http.FileServer(http.Dir("./frontend/dist"))
	// http.Handle("/", http.StripPrefix("/", fs))

	th := newTodoHandler()
	http.Handle("/api/todos", th)

	port := ":8080"
	fmt.Println("Started and listening on port", port)
	log.Panic(
		http.ListenAndServe(port, nil),
	)
}
