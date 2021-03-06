package chat

import (
	"encoding/json"
	"github.com/ElegantSoft/shabahy/middlewares"
	"github.com/ElegantSoft/shabahy/rooms"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var (
	roomService = rooms.RoomService
)

func SetupSocket(server *socketio.Server) {
	server.OnConnect("/", func(s socketio.Conn) error {
		return nil
	})

	server.OnEvent(nameSpaces.CHAT, events.JoinRoom, func(s socketio.Conn, roomHash string) {
		s.Join(roomsPrefix.SimpleRoom + roomHash)
	})

	server.OnEvent(nameSpaces.CHAT, events.SendMessage, func(s socketio.Conn, data string) {
		err, id := middlewares.AuthorizeSocket(s)
		if err != nil {
			log.Println(err)
			return
		}
		var message roomMessage
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			log.Fatal(err)
		}
		err = roomService.AppendMessage(message.RoomHash, message.Message, id)
		if err != nil {
			log.Println(err)
		}
		server.BroadcastToRoom(nameSpaces.CHAT, roomsPrefix.SimpleRoom+message.RoomHash, events.ReceiveMessage, message.Message)

	})
}
