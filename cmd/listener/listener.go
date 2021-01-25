package listener

import (
	"binance/listener/infrastructure"
	"binance/listener/interfaces"
	"log"
	"os"
	"os/signal"
	"strings"

	gowebsocket "github.com/Satssuki/GoWebsocket"
)

// RunListener for binance websocket
func RunListener() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	urlList := strings.Split(
		os.Getenv("DATABASE_URL"), ",",
	)

	dbNameList := strings.Split(
		os.Getenv("DATABASE_NAME"), ",",
	)

	userList := strings.Split(
		os.Getenv("DATABASE_USER"), ",",
	)

	passwordList := strings.Split(
		os.Getenv("DATABASE_PASSWORD"), ",",
	)

	infrastructure.InitBinanceDB(urlList, dbNameList, userList, passwordList)

	client := gowebsocket.New(os.Getenv("WSURL"))
	client.OnConnected = interfaces.HandleOpen
	client.OnTextMessage = interfaces.HandleMessage
	client.OnConnectError = interfaces.HandleError
	client.OnDisconnected = interfaces.HandleClose
	client.OnPingReceived = interfaces.HandlePing
	client.OnPongReceived = interfaces.HandlePong

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
