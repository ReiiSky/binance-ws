package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"

	gowebsocket "github.com/Satssuki/GoWebsocket"
)

func HandleError(err error, socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('error', now());")
	fmt.Println("Close")
	ReConnect(socket)
}
