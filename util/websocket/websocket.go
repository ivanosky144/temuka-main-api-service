package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	rest "github.com/temuka-api-service/util/rest"
)

var (
	Upgrader  = websocket.Upgrader{}
	GlobalHub = NewHub()
)

type Hub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan interface{}
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*websocket.Conn]bool),
		Broadcast:  make(chan interface{}),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.Register:
			h.Clients[conn] = true
		case conn := <-h.Unregister:
			if _, ok := h.Clients[conn]; ok {
				delete(h.Clients, conn)
				conn.Close()
			}
		case message := <-h.Broadcast:
			for conn := range h.Clients {
				err := conn.WriteJSON(message)
				if err != nil {
					h.Unregister <- conn
					conn.Close()
				}
			}
		}
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to upgrade connection"})
		return
	}

	GlobalHub.Register <- conn
	defer func() {
		GlobalHub.Unregister <- conn
	}()

	for {
		var message interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			break
		}
		GlobalHub.Broadcast <- message
	}
}
