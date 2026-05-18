package store

import (
	"github.com/YaelDev-HS/redsocial-go/internal/data"
	"github.com/gorilla/websocket"
)

const (
	NEW_MESSAGE = "NEW_MESSAGE"
)

type StoreMessage struct {
	Type    string        `json:"type"`
	Message *data.Message `json:"message"`
}

type Store interface {
	AddConn(conn *websocket.Conn)
	RemoveConn(conn *websocket.Conn)
	NotifyAll(message *StoreMessage)
}

func New() Store {
	return &WebsocketStore{
		connectios: make([]*websocket.Conn, 0),
	}
}
