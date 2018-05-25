package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}

}
