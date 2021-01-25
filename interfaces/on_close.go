package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"
	"time"

	gowebsocket "github.com/Satssuki/GoWebsocket"
)

func HandleClose(err error, socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('close', now());")
	fmt.Println(err, "close")

	ReConnect(socket)
}

var reconnectWorker = false

func ReConnect(sock gowebsocket.Socket) {
	if reconnectWorker {
		return
	}
	reconnectWorker = true
	go func() {
		for {
			if sock.IsConnected {
				reconnectWorker = false
				return
			}
			<-time.After(time.Second * 2)
			fmt.Println("Reconnecting")
			sock.Connect()
		}
	}()
}
