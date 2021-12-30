package download

import (
	"fmt"
	"net/http"
	"strconv"
)

func (d Downloader) sendReq(method string) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		d.Link,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (d Downloader) Connect() (int, error) {
	fmt.Println("Establishing new Connection")
	req, err := d.sendReq("HEAD")
	if err != nil {
		return -1, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	fmt.Printf("Got response %v\n", resp.StatusCode)

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return -1, err
	}
	return size, nil
}

func (d Downloader) Download(part [2]int) error {
	req, err := d.sendReq("GET")
	if err != nil {
		return err
	}

	partBytes := fmt.Sprintf("bytes=%v-%v", part[0], part[1])
	req.Header.Set("Range", partBytes)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("Part of File Downloaded")
	fmt.Printf("Downloaded %v bytes", resp.Header.Get("Content-Length"))
	return nil
}
