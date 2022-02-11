package main

import (
	"os"
)

func main() {

	chPort := make(chan string)

	if len(os.Args) == 2 && os.Args[1] == "--webview" {
		go webserverStart(chPort)
		webviewStart(chPort)
	} else {
		webserverStart(nil)
	}
}
