package interfaces

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandlePong(pong string, socket gowebsocket.Socket) {
	fmt.Println("Pong")
}
