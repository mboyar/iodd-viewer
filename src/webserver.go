package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmplFile = "tmpl/template.html"
var port string

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// thx to https://github.com/TutorialEdge/go-file-upload-tutorial
	// for the file upload example

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	t, err := template.New("template.html").ParseFiles(tmplFile)
	check(err)

	err = t.Execute(w, generateHtmlData(file))
	check(err)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.New("template.html").ParseFiles(tmplFile)
	check(err)

	var dataHtml Table
	err = t.Execute(w, dataHtml)
	check(err)
}

func webserverStart(c chan string) {

	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/upload", uploadHandler)

	//Thx to https://inblog.in/Deploy-a-Go-Golang-Web-App-to-Heroku-EAx7ifbLzD
	port := os.Getenv("PORT")

  	if len(port) == 0 {
		port = "8080"
  	}

	if c != nil {
		c <- port
	}
	
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
