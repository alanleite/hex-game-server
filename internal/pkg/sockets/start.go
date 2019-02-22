package sockets

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/alanleite/hex-game-server/internal/pkg/routes"
	"github.com/rs/cors"
)

//Start socket server
func Start() {

	flag.Parse()    // Don't know what it does... I know, shame on me...
	log.SetFlags(0) // It ether

	// Cors Configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("CORS") + ":" + os.Getenv("PORT")},
		AllowCredentials: true,
	})

	// Get sockets router
	router := routes.Socket()

	// Make connections handler with cors stuffs
	handler := c.Handler(router)

	// Start Server
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))

}
