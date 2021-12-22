package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var port = 4837
var addr = fmt.Sprintf(":%d", port)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		clientPort := 3721
		return origin == fmt.Sprintf("http://localhost:%d", clientPort)
	},
}
var db *pg.DB

func handleMessages(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err = c.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	channel := getChannel()
	for {
		message := <-channel
		err = c.WriteMessage(websocket.TextMessage, []byte(message.Payload))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func connect() {
	db = pg.Connect(&pg.Options{
		Addr:     "database:5432",
		User:     "admin",
		Password: "admin",
		Database: "message",
	})
}

func getChannel() <-chan pg.Notification {
	context := db.Context()
	listener := db.Listen(context, "message_channel")
	return listener.Channel()
}

func main() {
	log.SetFlags(0)
	connect()
	http.HandleFunc("/messages", handleMessages)
	log.Printf("starting webserver on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
