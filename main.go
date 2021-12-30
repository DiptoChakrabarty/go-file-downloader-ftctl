package main

import (
	"fmt"

	"ytctl/download"
)

func main() {
	fmt.Println("This is Start")
	vd := download.Downloader{
		Link:        "https://www.youtube.com/watch?v=vdhSk8vCx-k&t=328s",
		Path:        "proj.mp4",
		Connections: 10,
	}

	fmt.Println(vd)

	err := vd.Connect()
	if err != nil {
		panic(err)
	}

}
