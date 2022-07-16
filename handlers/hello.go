package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
 l *log.Logger 
}
// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}
// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		h.l.Println("Hey ! Nice to meet you.")
		d,err := ioutil.ReadAll(r.Body) 
		if err!=nil{
			http.Error(rw,"OOPS",http.StatusBadRequest)
        // rw.WriteHeader(http.StatusBadRequest)  this writerHeader allows to return status code and other responses to the users .
		 return
		}
		//WRITE the response
		fmt.Fprintf(rw,"Seni Seviorum !! %s",d)
	})
	
}