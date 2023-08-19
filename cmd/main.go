package main

import (
	"log"
	"net/http"

	"example/backend/pkg/config"
	"example/backend/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Setup Database

	//  "<username>:<password>@tcp(<url>:<port>)/<database_name>?charset=utf8mb4&parseTime=True"
	dataSourceName := "root:1230@tcp(localhost:3306)/example?charset=utf8mb4&parseTime=True"
	config.InitDB(dataSourceName)
	// Setup Router
	router := mux.NewRouter()

	// Setup Routes
	userRouter := router.PathPrefix("/users").Subrouter()
	routes.InitUserRoute(userRouter)

	serverAddr := "localhost:8080"

	// Use the router as the default HTTP handler
	http.Handle("/", router)

	// Start the HTTP server
	log.Printf("Server is running on %s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatal(err)
	}
}
