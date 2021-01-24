package interfaces

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandleClose(err error, socket gowebsocket.Socket) {
	fmt.Println("Close", err)
}
