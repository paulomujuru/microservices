package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Marcello")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oooops", http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(rw, "Hello %s Marcello \n", data)
}
