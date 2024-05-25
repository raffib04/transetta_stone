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

func (rm *RoomManager) listActiveRooms() []string {
	var rooms []string
	for roomName := range rm.rooms {
		rooms = append(rooms, roomName)
	}
	return rooms
}
