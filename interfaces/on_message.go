package interfaces

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandleMessage(message string, socket gowebsocket.Socket) {
	fmt.Println(message)
}
