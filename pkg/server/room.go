package server

import (
	"../client"
	"context"
	"log"
)

var rooms roomList

func init() {
	rooms = roomList{rooms: map[string]*Room{}}
}

type Room struct {
	id       string
	password string
	hub      *Hub
}

type roomList struct {
	rooms map[string]*Room
}

func getRoom(id string, password string, ctx context.Context) (room *Room, auth bool) {
	room, exist := rooms.rooms[id]
	if exist {
		if password == room.password {
			return room, true
		}
		return room, false
	}

	hub := newHub()
	go hub.run(ctx)

	room = &Room{id: id, password: password, hub: hub}
	go room.deactivateRoom(ctx)
	rooms.rooms[id] = room

	return room, true
}

func removeRoom(id string) (exist bool) {
	room, exist := rooms.rooms[id]
	if !exist {
		return exist
	}
	log.Println("delete room :", room.id)
	room = nil
	delete(rooms.rooms, id)
	return exist
}

func (room *Room) broadcast(message []byte) {
	room.hub.broadcast <- message
}

func (room *Room) register(client *client.Client) {
	room.hub.register <- client
}

func (room *Room) unregister(client *client.Client) {
	room.hub.unregister <- client
}

func (room *Room) deactivateRoom(ctx context.Context) {
	select {
	case active := <-room.hub.active:
		if !active {
			removeRoom(room.id)
		}
	case <-ctx.Done():
		return
	}

}
