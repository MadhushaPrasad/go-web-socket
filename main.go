package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ // configure the upgrader
	ReadBufferSize:  1024, // read buffer size in bytes
	WriteBufferSize: 1024, // write buffer size in bytes
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page") // write data to response
}

func reader(conn *websocket.Conn) { // read data from client
	for { // infinite loop
		messageType, p, err := conn.ReadMessage() // read message from client
		if err != nil {                           // check if error
			log.Println(err) // log error / print error
			return           // return from function
		}
		log.Println(string(p))                      // log message
		if err = conn.WriteMessage(messageType, p); // write message to client
		err != nil {                                // check if error
			log.Println(err) // log error / print error
			return           // return from function
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) { //websocket endpoint
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // allow all connections ( cross-origin access )

	ws, err := upgrader.Upgrade(w, r, nil) // upgrade connection to websocket

	if err != nil { // check if error
		log.Println(err) // log error / print error
		return
	}

	log.Println("Client Successfully Connected...") // log message

	reader(ws) // call reader function
}

func setupRoutes() { // setup routes
	http.HandleFunc("/", homePage)     // handle home page
	http.HandleFunc("/ws", wsEndpoint) // handle websocket endpoint
}

func main() {
	fmt.Println("Go WebSockets")                 // print message
	setupRoutes()                                //call setupRoutes function
	log.Fatal(http.ListenAndServe(":8080", nil)) // listen on port 8080
}
