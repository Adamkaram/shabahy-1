package chat

var nameSpaces = struct {
	CHAT string
}{
	CHAT: "/chat",
}

var events = struct {
	JoinRoom       string
	SendMessage    string
	ReceiveMessage string
}{
	JoinRoom:       "join-room",
	SendMessage:    "send-message",
	ReceiveMessage: "receive-message",
}

var roomsPrefix = struct {
	SimpleRoom string
}{
	SimpleRoom: "room-",
}
