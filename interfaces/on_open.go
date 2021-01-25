package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"

	gowebsocket "github.com/Satssuki/GoWebsocket"
)

func HandleOpen(socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('open', now());")
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('ping', now());")
	fmt.Println("OPEN")
}
