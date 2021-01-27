package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"
	"time"

	gowebsocket "github.com/Satssuki/GoWebsocket"
)

func HandlePong(pong string, socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('pong', now());")
	fmt.Println("PONG", time.Now())
}
