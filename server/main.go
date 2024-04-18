package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP Network address")
	flag.Parse()
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorlog: errorlog,
		infolog:  infolog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}

	infolog.Printf("Starting server in %s", *addr)

	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
