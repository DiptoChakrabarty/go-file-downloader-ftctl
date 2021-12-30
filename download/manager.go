package download

import (
	"fmt"
	"net/http"
	"strconv"
)

func (d Downloader) checkLink(method string) (*http.Request, error) {
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
	req, err := d.checkLink("HEAD")
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
