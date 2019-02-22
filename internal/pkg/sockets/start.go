package sockets

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/alanleite/hex-game-server/internal/pkg/controllers/echo"
)

func Start() {

	var addr = flag.String("addr", os.Getenv("URL"), "http service address")

	flag.Parse()
	log.SetFlags(0)
	routes()
	log.Fatal(http.ListenAndServe(*addr, nil))

}

func routes() {
	http.HandleFunc("/echo", echo.Index)
}
