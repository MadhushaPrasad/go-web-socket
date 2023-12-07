package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Greeting string `json:"greeting"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConnections *websocket.Conn
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		//check http.Request origin
		//make sure it's ok to access this resource
		return true
	}

	wsConnections, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Could not upgrade connection to websocket", err.Error())
		return
	}

	defer wsConnections.Close()

	// event loop
	for {
		var message Message

		err := wsConnections.ReadJSON(&message)
		if err != nil {
			fmt.Println("Error reading json.", err.Error())
			continue
		}

		fmt.Println("Message received: ", message.Greeting)
	}
}

func main() {
	fmt.Println("Hello, world.")

	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

}
