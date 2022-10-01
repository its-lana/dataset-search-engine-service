package main

import (
	"be-dse/middleware"
	"be-dse/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {

	r := router.Router()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handler := cors.Handler(middleware.NewAuthMiddleware(r))
	port := os.Getenv("PORT")

	fmt.Println("Server dijalankan pada port ", port)
	log.Fatal(http.ListenAndServe(port, handler))

}
