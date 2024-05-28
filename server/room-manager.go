package main

import "encoding/json"

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

func (rm *RoomManager) listActiveRooms() []byte {
	//create map where key is room name and value is number of clients in room

	type roomData struct {
		RoomName   string `json:"roomName"`
		NumClients int    `json:"numClients"`
	}

	var roomsData []roomData

	for roomName, room := range rm.rooms {
		roomsData = append(roomsData, roomData{
			RoomName:   roomName,
			NumClients: len(room.clients),
		})
	}

	roomsDataBytes, err := json.Marshal(roomsData)
	if err != nil {
		return []byte{}
	}

	return roomsDataBytes
}
