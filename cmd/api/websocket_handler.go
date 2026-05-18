package main

import (
	"errors"
	"fmt"
	"log"
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
	token, ok := app.GetParam(r, "token")
	_, err := app.getUserAuthByToken(token)

	if err != nil {
		switch {
		case errors.Is(err, unauthorizedError):
			app.unauthorized(w, err)
		default:
			app.internalServerError(w, err)
		}

		return
	}

	if !ok {
		app.unauthorized(w, fmt.Errorf("missing token"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	app.store.AddConn(conn)

	defer func() {
		log.Println("Close connection")
		app.store.RemoveConn(conn)
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
