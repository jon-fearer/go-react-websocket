package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var port = 4837
var addr = fmt.Sprintf(":%d", port)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		clientPort := 3721
		return origin == fmt.Sprintf("localhost:%d", clientPort) ||
			strings.HasPrefix(origin, "chrome-extension://")
	},
}
var db *pg.DB
var channel <- chan pg.Notification

func handleMessages(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	for {
		log.Println("checking channel")
		message := <-channel
		log.Println(message)
		err = c.WriteMessage(websocket.TextMessage, []byte(message.Payload))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func connectAndListen() {
	db = pg.Connect(&pg.Options{
		Addr:     "database:5432",
		User:     "admin",
		Password: "admin",
		Database: "message",
	})
	context := db.Context()
	listener := db.Listen(context,"message_channel")
	channel = listener.Channel()
}

func main() {
	log.SetFlags(0)
	connectAndListen()
	http.HandleFunc("/messages", handleMessages)
	log.Printf("starting webserver on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
