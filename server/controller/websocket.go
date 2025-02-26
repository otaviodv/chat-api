package controller

import (
	"chat-api/helper"
	"chat-api/model"
	"context"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type socketHandler struct {
	upgrader websocket.Upgrader
	rooms    map[primitive.ObjectID]map[*websocket.Conn]context.CancelFunc
	mu       *sync.Mutex
}

var sh socketHandler
var receivedMsgs chan model.Message

func init() {
	sh = socketHandler{
		upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		rooms:    map[primitive.ObjectID]map[*websocket.Conn]context.CancelFunc{},
		mu:       &sync.Mutex{},
	}

	receivedMsgs = make(chan model.Message, 10)
	go handleReceivedMsg(receivedMsgs)
}

func handleReceivedMsg(msgs chan model.Message) {
	for msg := range msgs {
		sh.mu.Lock()

		connections, ok := sh.rooms[msg.RoomId]
		if !ok || len(connections) == 0 {
			sh.mu.Unlock()
			continue
		}

		for conn, cancel := range connections {
			if err := conn.WriteJSON(msg); err != nil {
				cancel()
			}
		}
		sh.mu.Unlock()
	}
}

func HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	var roomID primitive.ObjectID
	{
		room := model.Room{}
		err := room.Get(rawID)
		if err != nil {
			helper.ErrorJSON(w, err, http.StatusNotFound)
			return
		}
		roomID = room.Id
	}
	con, err := sh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer con.Close()

	ctx, cancel := context.WithCancel(r.Context())

	sh.mu.Lock()
	if _, ok := sh.rooms[roomID]; !ok {
		sh.rooms[roomID] = make(map[*websocket.Conn]context.CancelFunc)
	}
	slog.Info("new client connected", "roomID", roomID, "clientIP", r.RemoteAddr)
	sh.rooms[roomID][con] = cancel

	sh.mu.Unlock()

	<-ctx.Done()

	sh.mu.Lock()
	delete(sh.rooms[roomID], con)
	sh.mu.Unlock()
}
