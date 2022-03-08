/*
 * hello-world
 * Thomas vanBommel
 * 2022-03-06
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
