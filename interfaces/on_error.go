package interfaces

import (
	"binance/listener/infrastructure"

	"github.com/sacOO7/gowebsocket"
)

func HandleError(err error, socket gowebsocket.Socket) {
	infrastructure.InsertQuery(0, "INSERT INTO `log_binance` (`stat`, `timestamp`)VALUES ('error', now());")
}
