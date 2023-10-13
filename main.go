package main

import (
	"log"
	"net/http"
)

func main() {
	// Criando um FileServer
	fileServer := http.FileServer(http.Dir("./public"))

	// Criando um mux
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog"))
	})

	// Subindo o Server
	// log.Fatal vai mostrar um possível erro
	log.Fatal(http.ListenAndServe(":8080", mux))
}
