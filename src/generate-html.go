package main

import (
	"html/template"
	"os"
	"path/filepath"
)

func generateHtml(data Table, filename string) {

	t, err := template.New("template.html").ParseFiles(filename)
	check(err)

	f, err := os.Create(filepath.Dir(filename) + "/index.html")
	check(err)

	err = t.Execute(f, data)
	check(err)
}
