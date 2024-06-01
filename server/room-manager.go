package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

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
	var result sql.Result
	var err error

	if !exists {

		result, err = db.Exec("INSERT INTO rooms (name, created_at) VALUES (?, ?);", name, time.Now())

		if err != nil {
			fmt.Println("Error inserting room into database: ", err)
			return nil
		}

		room_id, err := result.LastInsertId()

		if err != nil {
			fmt.Println("Error getting room id: ", err)
			return nil
		}

		id := int(room_id)

		room = newRoom(id)
		rm.rooms[name] = room
		go room.runRoom()
	}

	room.getMessages()

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
