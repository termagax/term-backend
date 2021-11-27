package api

import (
	"fmt"
	"net/http"

	"github.com/apex/log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
}

func noMaam(w http.ResponseWriter, r *http.Request) {
	fmt.Println("no")
}

func Serve() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/yttranscriber", noMaam)

	if http.ListenAndServe(":8080", nil) != nil {
		log.Fatal("yes")
	}
}
