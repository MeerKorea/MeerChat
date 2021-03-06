package server

import (
	"log"

	"github.com/jjwow73/MeerChat/pkg/chat"
)

type hub struct {
	connInfos  map[*connInfo]bool
	broadcast  chan *chat.MessageProtocol
	register   chan *connInfo
	unregister chan *connInfo
	done       chan interface{}
}

func newHub() *hub {
	hub := &hub{
		connInfos:  make(map[*connInfo]bool),
		broadcast:  make(chan *chat.MessageProtocol),
		register:   make(chan *connInfo),
		unregister: make(chan *connInfo),
		done:       make(chan interface{}),
	}
	go hub.run()
	return hub
}

func (hub *hub) run() {
	for {
		select {
		case connInfo := <-hub.register:
			hub.addConn(connInfo)
		case connInfo := <-hub.unregister:
			hub.removeConn(connInfo)
		case message := <-hub.broadcast:
			log.Println("Broadcasting")
			hub.sendMessageToEachConn(message)
		case <-hub.done:
			return
		}
	}
}

func (hub *hub) addConn(connInfo *connInfo) {
	log.Println("register conn")
	hub.connInfos[connInfo] = true
}

func (hub *hub) removeConn(connInfo *connInfo) {
	log.Println("unregister conn")
	if _, exist := hub.connInfos[connInfo]; exist {
		delete(hub.connInfos, connInfo)
		close(connInfo.channel)
		connInfo = nil
	}
	// If hub has no connection then remove hub
	if len(hub.connInfos) == 0 {
		close(hub.done)
	}
}

func (hub *hub) sendMessageToEachConn(message *chat.MessageProtocol) {
	for connInfo := range hub.connInfos {
		select {
		case connInfo.channel <- message:
		default:
			log.Println("error may occurred")
			hub.removeConn(connInfo)
		}
	}
}
