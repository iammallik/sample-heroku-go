package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/negroni"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/iammallik/sample-heroku-go/app/config"
	route "github.com/iammallik/sample-heroku-go/app/http"
)

func main() {
	log.Printf("Server started")
	err := config.ConfigureEnv()
	if err != nil {
		log.Panic("Configuration setup failed")
	}

	router := route.NewRouter()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	log.Println("Port =", os.Getenv("PORT"))
	port := ":" + os.Getenv("PORT")
	graceful.Run(port, 10*time.Second, n)
}
