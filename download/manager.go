package download

import (
	"fmt"
	"net/http"
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

func (d Downloader) Connect() error {
	fmt.Println("Establishing new Connection")
	req, err := d.checkLink("HEAD")
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("Got response %v\n", resp.StatusCode)
	return nil
}
