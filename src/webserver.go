package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmplFile = "tmpl/template.html"

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

func startWebserver() {

	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/upload", uploadHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
