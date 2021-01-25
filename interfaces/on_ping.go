package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"

	gowebsocket "github.com/Satssuki/GoWebsocket"
)

func HandlePing(ping string, socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('ping', now());")
	fmt.Println("PING")
}
