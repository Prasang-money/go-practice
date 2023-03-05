package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":8085", "HTTP network address")
	flag.Parse()

	// creating custom infoLog and errorLog
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileserver := http.FileServer(http.Dir("ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	sv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("starting server on address: %s", *addr)
	if err := sv.ListenAndServe(); err != nil {
		errorLog.Panic(err)
	}

}
