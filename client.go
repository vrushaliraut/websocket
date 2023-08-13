package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	/*- The make built-in function allocates and initializes an object of type slice, map, or chan (only)
	- crated interrupt variable with parameter in make as a channel with some buffered size */
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "ws://localhost:8080/ws"

	log.Printf("connecting to %s", url)

	connection, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("dial:", err)
	}

	defer connection.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("Received message: %s", message)
		}
	}()

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter message: ")
			input, _ := reader.ReadString('\n')
			err := connection.WriteMessage(websocket.TextMessage, []byte(input))
			if err != nil {
				log.Println("write: ", err)
				return
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Send periodic messages if needed
			/*err := connection.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write: ", err)
				return
			}*/
		case <-interrupt:
			log.Println("interrupt")
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
				websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close : ", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
