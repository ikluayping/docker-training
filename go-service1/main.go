package main

import (
	"github.com/codegangsta/negroni"
	"go-service1/pkg/approuter"
	"log"
	"net/http"
	"os"
)

func main() {
	appRouter := approuter.NewCreateAppRouter()
	n := negroni.Classic()
	n.UseHandler(appRouter)
	url := os.Getenv("PORT")
	log.Printf("Url: %s", url)
	err := http.ListenAndServe(url, n)
	if err != nil {
		log.Printf("cannot listen to url %s with error %s", url, err.Error())
	}
}
