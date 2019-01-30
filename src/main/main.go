package main

import "net/http"

type demo struct{}

func (*demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		w.WriteHeader(404)

		return
	}

	w.Write([]byte("this is a demo http handler" + r.RequestURI))
}

func main() {
	httpServer := &http.Server{
		Addr: ":5000",
		Handler:&demo{},
	}

	httpServer.ListenAndServe()
}