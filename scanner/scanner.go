package scanner

import "github.com/gorilla/websocket"

// Scanner should implement ...
// Add your url-based scanner here.
type Scanner interface {
	Run(fuzzableURLs []string, conn *websocket.Conn)
	Report() interface{}
}
