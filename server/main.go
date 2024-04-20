package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	go func() {

		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		infolog.Println("Shutting down server...")

		if err := srv.Shutdown(nil); err != nil {
			errorlog.Fatalf("Server forced to shutdown: %v", err)
		}

	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		errorlog.Fatalf("Server startup failed: %v", err)
	}
}
