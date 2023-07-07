package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	serv := NewGetServer()
	mux.HandleFunc("/generate_preview", serv.generateHandler)
	port := os.Getenv("GP_PORT")
	if port == "" {
		port = "8380"
	}
	listen := os.Getenv("GP_LISTEN")
	fmt.Printf("LISTEN ON %s:%s\n", listen, port)
	log.Fatalln(http.ListenAndServe(listen+":"+port, mux))

}
