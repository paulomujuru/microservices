package main

import(
	"fmt"
	"log"
)


func main() {


	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request)){
		log.Println("Hello world")
	}
	http.ListenAndServe(":9000", nil)
}


