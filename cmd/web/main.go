package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/IfedayoAwe/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	var err error
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	infoFile, err := os.OpenFile(filepath.Join(mydir, "/logs", "/info.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer infoFile.Close()
	errFile, err := os.OpenFile(filepath.Join(mydir, "/logs/error.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer errFile.Close()

	dbUser, dbPass, dbName := os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME")
	conf := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbUser, dbPass, dbName)

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", conf, "MySQL data source name")
	flag.Parse()

	infoLog := log.New(infoFile, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(errFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err.Error())

}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
