package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var err error
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	infoFile, err := os.OpenFile(filepath.Join(mydir, "/logs", "/info.log"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer infoFile.Close()
	errFile, err := os.OpenFile(filepath.Join(mydir, "/logs/error.log"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer errFile.Close()
	infoLog := log.New(infoFile, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(errFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err.Error())
}
