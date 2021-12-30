package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func (d Downloader) Download(part [2]int, index int) error {
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

	err = d.writeFile(d.Path, index, resp)
	if err != nil {
		return err
	}

	return nil
}

func (d Downloader) writeFile(path string, index int, resp *http.Response) error {
	bf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("part-%v.tmp", index)
	err = ioutil.WriteFile(filename, bf, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (d Downloader) MergeFiles() error {
	fil, err := os.OpenFile(d.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fil.Close()

	for i := 0; i < d.Connections; i++ {
		filename := fmt.Sprintf("part-%v.tmp", i)
		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		byt, err := fil.Write(buf)
		if err != nil {
			return err
		}
		fmt.Printf("File %v merged , Bytes : %v\n", i, byt)
	}
	return nil
}
