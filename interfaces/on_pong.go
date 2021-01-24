package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandlePong(pong string, socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('pong', now());")
	fmt.Println("PONG")
}
