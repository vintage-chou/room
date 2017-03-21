package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world from Go in minimal Docker container.")
}

func main() {

	http.HandleFunc("/", helloHandler)

	fmt.Println("started, serving at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("LisenAndServe:" + err.Error())
	}
}
