package store

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketStore struct {
	connectios []*websocket.Conn
	mu         sync.RWMutex
}

func (w *WebsocketStore) AddConn(conn *websocket.Conn) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.connectios = append(w.connectios, conn)
}

func (w *WebsocketStore) RemoveConn(conn *websocket.Conn) {
	w.mu.Lock()
	defer w.mu.Unlock()
}

func (w *WebsocketStore) NotifyAll(message *StoreMessage) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	for _, v := range w.connectios {
		v.WriteJSON(message)
	}
}
