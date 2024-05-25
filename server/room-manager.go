package main

type RoomManager struct {
	rooms map[string]*Room
}

func newRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),
	}
}

func (rm *RoomManager) getRoom(name string) *Room {
	room, exists := rm.rooms[name]
	if !exists {
		room = newRoom()
		rm.rooms[name] = room
		go room.runRoom()
	}
	return room
}
