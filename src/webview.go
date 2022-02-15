// +build desktop

package main

import (
	"fmt"

	"github.com/webview/webview"
)

func webviewStart(c chan string) {

	debug := true

	port := <-c
	url := "http://localhost:" + port

	fmt.Println("Starting webview... URL: " + url)

	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("IODD Viewer")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
