package main

import (
	"go-mongo/db"
	"go-mongo/handler"
	"go-mongo/router"
	"go-mongo/store"
	"log"
	"os"

	_ "go-mongo/docs" // docs is generated by Swag CLI, you have to import it.

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	mongoClient, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	g := r.Group("")
	projectDb := db.SetupProjectDb(mongoClient)
	project := store.NewProjectStore(projectDb)
	h := handler.NewHandler(project)
	h.Register(g)
	r.Logger.Fatal(r.Start("0.0.0.0:" + port))
}
