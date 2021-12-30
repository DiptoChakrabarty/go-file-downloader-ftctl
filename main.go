package main

import (
	"fmt"

	"github.com/DiptoChakrabarty/ytctl/types"
)

func main() {
	fmt.Println("This is Start")
	vd := types.Downloader{
		Link:        "https://www.youtube.com/watch?v=vdhSk8vCx-k&t=328s",
		Path:        "proj.mp4",
		Connections: 10,
	}

	fmt.Println((vd))

}
