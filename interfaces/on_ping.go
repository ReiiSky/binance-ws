package interfaces

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandlePing(ping string, socket gowebsocket.Socket) {
	fmt.Println("Ping")
}
