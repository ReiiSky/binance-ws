package listener

import (
	"binance/listener/interfaces"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/sacOO7/gowebsocket"
)

// RunListener for binance websocket
func RunListener() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	client := gowebsocket.New(os.Getenv("WSURL"))
	client.OnConnected = interfaces.HandleOpen
	client.OnTextMessage = interfaces.HandleMessage
	client.OnConnectError = interfaces.HandleError
	client.OnDisconnected = interfaces.HandleClose
	client.OnPingReceived = interfaces.HandlePing
	client.OnPongReceived = interfaces.HandlePong
	fmt.Println("Connected")
	client.Connect()
	
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			client.Close()
			return
		}
	}
}
