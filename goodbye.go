package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
 l *log.Logger 
}

func NewBye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r* http.Request){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		g.l.Println("Will see you on screen.")
	
		rw.Write([]byte("Rijaye Derim !!"))
	})
	
}