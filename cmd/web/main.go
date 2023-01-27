package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.Handle("/snippet/download/", http.StripPrefix("/snippet/download", http.HandlerFunc(downloadHandler)))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	var err error
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	infoFile, err := os.OpenFile(filepath.Join(mydir, "/info.log"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()
	infoLog := log.New(infoFile, "INFO\t", log.Ldate|log.Ltime)
	errFile, err := os.OpenFile(filepath.Join(mydir, "/error.log"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer errFile.Close()
	errorLog := log.New(errFile, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
