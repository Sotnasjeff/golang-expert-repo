package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	//realizando com função anonima

	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World"))
	//})

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

//Toda vez que implementamos um novo handler com mux, ele implementa uma interface de handler que nos obriga
// a ter um método chamado ServerHTTP
