package interfaces

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

func HandleError(err error, socket gowebsocket.Socket) {
	fmt.Println("Error", err)
}
