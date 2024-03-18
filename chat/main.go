package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var upgrader = websocket.Upgrader{}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleWebSocket)
	go handleMessages()
	log.Println("service is running, port8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("service start fail: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection upgrade fail", err)
		return
	}

	clients[conn] = true
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message fail: ", err)
			break
		}

		broadcast <- string(message)
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("send message fail: ", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
