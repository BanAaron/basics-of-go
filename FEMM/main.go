package main

import (
	"fmt"
	"html/template"
	"net/http"

	"aaronbarratt.dev/go/femm/api"
	"aaronbarratt.dev/go/femm/data"
)

// handleHello
func handleHello(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println(err)
	}
}

func handleTemplate(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("internal server error"))
		return
	}
	html.Execute(writer, data.GetAll())
}

func main() {
	// first crete a server multiplexer
	server := http.NewServeMux()

	// we can create routes this
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	// we can also handle non-function objects like this
	fileServer := http.FileServer(http.Dir("./public"))
	server.Handle("/", fileServer)

	// then start the server like this
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		errorMessage := fmt.Errorf("error while opening the server: %s", err)
		panic(errorMessage)
	}
}
