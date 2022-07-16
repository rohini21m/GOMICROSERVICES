package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rohini21m/GOMICROSERVES/handlers"

	"github.com/nicholasjackson/env"
)


func main(){
	var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
	env.Parse()

	l := log.New(os.Stdout,"product-api",log.LstdFlags)
	//creating handlers 
	hh := handlers.NewHello(l)
	gb := handlers.NewBye(l)

	// register handler with server .

	sm:= http.NewServeMux() 
	sm.Handle("/",hh)
	sm.Handle("/goodbye",gb)
//	http.ListenAndServe(":9090", sm)
	//creating new server 
	s:= &http.Server{
		Addr: *bindAddress, // set the address 
		ErrorLog: l, // set logger for server 
		Handler: sm, // set default handler 
		IdleTimeout: 120*time.Second,  // max time for connections using TCP Keep-Alive
		ReadTimeout: 1*time.Second, // max time to read and request the client
		WriteTimeout: 1*time.Second, //max time to write the response to client 
	}
// start the server 
l.Println("starting the server on  port 9090")
	go func() {
	err:= s.ListenAndServe()
	
	if err!=nil{
		l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
	}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
c:= make(chan os.Signal,1)
signal.Notify(c,os.Interrupt)
signal.Notify(c,os.Kill)
// Block until a signal is received.
sig := <-c
log.Println(" received signal, gracefull shutdown", sig)
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
    defer cancel()
	s.Shutdown(ctx)

	s.ListenAndServe()
} 
