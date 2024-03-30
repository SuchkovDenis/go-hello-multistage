package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

var emojis = []string{"🐳", "🦁", "🐶", "🐬"}

func hello(w http.ResponseWriter, req *http.Request) {
	index := rand.Intn(len(emojis))
	fmt.Fprintf(w, "Hello from Go in Docker %s", emojis[index])
}

