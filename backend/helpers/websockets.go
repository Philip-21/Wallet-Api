package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

// sends messages from the server to all connected clients
var broadcast = make(chan []byte)

var upgrader = websocket.Upgrader{} 


func HandleWebsocket(c*gin.Context){
	//upgrades the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error in Upgrading websocket connection", err)
		return
	}
	clients[conn] = true 

	defer func() {
		conn.Close()
		delete(clients, conn)
	}()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Websocket error", err)
			break
		}
		broadcast <- message
	}

}



// A function that continiousy listens to the broadcst channel
func HandleBroadcast() {
	for {
		message := <-broadcast

		for client := range clients {
			// iterates over all connected clients and sends the message
			//to each client using the client.WriteMessage method
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Writer websocket error", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
