package interfaces

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandleOpen(socket gowebsocket.Socket) {
	fmt.Println("Connected")
}
