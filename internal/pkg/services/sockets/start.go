package sockets

import (
	"flag"
	"log"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{} // use default options

func Start() {

	flag.Parse()
	log.SetFlags(0)
}
