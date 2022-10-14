package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("running server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

//перенаправить потоки из stdout и stderr в файлы на диске при запуске приложения из терминала
//go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log

//для чего нужно многоуровневое логирование?
//горутина это функции?
//app.errorLog.Output(2, trace) что за call depth
