package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/MohammadLashkari/snippetbox/internal/models"
	"github.com/go-playground/form"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

func main() {

	host := flag.String("host", "localhost", "HTTP network host")
	port := flag.String("port", "8080", "HTTP network port")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	app := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   formDecoder,
	}

	srv := http.Server{
		Addr:     net.JoinHostPort(*host, *port),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("startig server on %s\n", net.JoinHostPort(*host, *port))
	errorLog.Fatal(srv.ListenAndServe())
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
