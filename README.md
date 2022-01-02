# ftctl
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

This is a file downloader cli written in Golang which uses the concurrent feature of go to download files. The cli is built using cobra.

## How to Run Project
```sh
- Clone Repo git@github.com:DiptoChakrabarty/go-file-downloader-ftctl.git
- Build cli using go build
- Use the execuable to run
```

## Commands present

- download : the main command used to perform actions under which other flags are present

```sh
- link : takes link of what to download
- path: location where you want to save in host system
- connections: number of concurrent requests to make
```

## Example Commands

* ./ftctl download --link=https://raw.githubusercontent.com/DiptoChakrabarty/go-gin-movies/main/docker-compose.yml --connection=10
* ./ftctl download --link=https://raw.githubusercontent.com/DiptoChakrabarty/go-gin-movies/main/docker-compose.yml --connection=10 --path=compose.yaml


The project has a few bugs you can check the issues for it [![GitHub issues](https://img.shields.io/github/issues/Naereen/StrapDown.js.svg)](https://GitHub.com/DiptoChakrabarty/go-file-downloader-ftctl/issues/)

![Download File](https://github.com/DiptoChakrabarty/go-file-downloader-ftctl/blob/main/images/download.PNG?raw=true)

Downloading a file

![Download Result](https://github.com/DiptoChakrabarty/go-file-downloader-ftctl/blob/main/images/download1_result.PNG?raw=true)

Downloaded File

![Download Video](https://github.com/DiptoChakrabarty/go-file-downloader-ftctl/blob/main/images/videodownload.PNG?raw=true)

Downloading a Video
