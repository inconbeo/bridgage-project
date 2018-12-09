package main

import (
	"fmt"
	"time"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":3001", nil))
// }

func main() {
	for true {
		fmt.Println("Infinite Loop 2")
		time.Sleep(time.Second)
	}
}
