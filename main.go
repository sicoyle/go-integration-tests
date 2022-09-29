package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloGophers)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Printf("http.ListenAndServe err: %v", err)
		os.Exit(1)
	}
}

func HelloGophers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello gophers!\n")
	log.Print("/ visited!")
}
