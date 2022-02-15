package cmd

import (
	"fmt"
	"ftctl/download"
	"sync"

	"github.com/spf13/cobra"
)

//var dl = &download.Downloader{}
//var s string

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "provide file link to download",
	//Args:  cobra.ExactArgs(1),
	RunE: downloadmanager,
}

func init() {
	RootCmd.AddCommand(DownloadCmd)
	DownloadCmd.PersistentFlags().StringP("link", "l", "", "Link of file to download")
	DownloadCmd.PersistentFlags().StringP("path", "p", "download.yml", "path to set download")
	DownloadCmd.PersistentFlags().IntP("connection", "c", 10, "No of connections to make")
	//DownloadCmd.Flags().StringVarP(&s, "kcheck", "k", s, "just check")
}

func downloadmanager(cmd *cobra.Command, args []string) error {
	fmt.Println("This is Start")
	//fmt.Println("This is value of kcheck", s)
	//input := strings.Join(args, " ")
	link, _ := cmd.Flags().GetString("link")
	path, _ := cmd.Flags().GetString("path")
	connection, _ := cmd.Flags().GetInt("connection")
	//fmt.Println(link, path, connection, input)
	var dl = download.Downloader{
		Link:        link,
		Path:        path,
		Connections: connection,
	}

	size, err := dl.Connect()
	if err != nil {
		return err
	}

	var parts = make([][2]int, dl.Connections)
	partSize := size / dl.Connections

	fmt.Printf("Size of each part is %d bytes\n", partSize)

	for i := 0; i < dl.Connections; i++ {
		if i > 0 {
			parts[i][0] = parts[i-1][1] + 1
		}

		if i < dl.Connections-1 {
			parts[i][1] = parts[i][0] + partSize
		} else {
			parts[i][1] = size - 1
		}
	}

	//fmt.Println(parts)
	var wg sync.WaitGroup

	for index, part := range parts {
		index := index
		part := part
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = dl.Download(part, index)
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
	dl.MergeFiles()
	return nil
}
