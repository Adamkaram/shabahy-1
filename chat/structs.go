package chat

type roomMessage struct {
	Message  string `json:"message"`
	RoomHash string `json:"room_hash"`
}
