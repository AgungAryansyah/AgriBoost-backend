package ws

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	ID   string
	Name string
	Conn *websocket.Conn
	Room string
}

type Room struct {
	Clients map[string]*Client
	Mutex   sync.Mutex
}

var Rooms = make(map[string]*Room)
var RoomsMutex sync.Mutex
