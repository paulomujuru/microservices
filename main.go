package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooopssss!!", http.StatusBadRequest)
			/*
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("Ooooooooo"))
			*/
			return
		}

		fmt.Fprintf(rw, "Hello %s world \n", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":8080", nil)
}
