package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *application) ConnectWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	defer func() {
		fmt.Println("Close connection")
		conn.Close()
	}()

	for {
		_, _, err := conn.ReadMessage()

		if err != nil {
			fmt.Printf("err: %s\n", err)
			break
		}
	}
}
