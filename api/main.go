package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	game "github.com/nbelkaid/battlesnakes/api/game"
)

func main() {

	user := os.Getenv("APP_DB_USER")
	pwd := os.Getenv("APP_DB_PASS")
	dbName := os.Getenv("APP_DB_NAME")
	host := os.Getenv("CONTAINER_DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pwd,
		host,
		dbPort,
		dbName)

	db, err := gorm.Open("postgres", dbConnString)
	if err != nil {
		fmt.Println("Not enable to connect to Database - ", dbConnString, err)
		panic(err)
	}

	r := mux.NewRouter()

	game.ConfigRoute(db, r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
