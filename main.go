package main

import (
	"fmt"

	"ytctl/download"
)

func main() {
	fmt.Println("This is Start")
	vd := download.Downloader{
		Link:        "https://raw.githubusercontent.com/DiptoChakrabarty/go-gin-movies/main/docker-compose.yml",
		Path:        "docker-compose.yml",
		Connections: 10,
	}

	fmt.Println(vd)

	size, err := vd.Connect()
	if err != nil {
		panic(err)
	}

	var parts = make([][2]int, vd.Connections)
	partSize := size / vd.Connections

	fmt.Printf("Size of each part is %d bytes\n", partSize)

	for i := 0; i < vd.Connections; i++ {
		if i > 0 {
			parts[i][0] = parts[i-1][1] + 1
		}

		if i < vd.Connections-1 {
			parts[i][1] = parts[i][0] + partSize
		} else {
			parts[i][1] = size - 1
		}
	}

	fmt.Println(parts)

	err = vd.Download(parts[0])
	if err != nil {
		panic(err)
	}

}
