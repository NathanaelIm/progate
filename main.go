package main

import (
	"log"
	"net/http"
	"progate/handler"
)

func main() {
	mux:=http.NewServeMux() 

	mux.HandleFunc("/",handler.Homehandler)
	mux.HandleFunc("/hello",handler.Hellohandler)
	mux.HandleFunc("/mario",handler.Mariohandler)
	mux.HandleFunc("/product",handler.Producthandler)
	mux.HandleFunc("/post-get",handler.PostGet)
	mux.HandleFunc("/form",handler.Form)
	mux.HandleFunc("/process",handler.Process)

	fileServer:=http.FileServer(http.Dir("assets"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	log.Println("Starting web on port 8080")

	err:=http.ListenAndServe(":8080",mux)
	log.Fatal(err)
}

