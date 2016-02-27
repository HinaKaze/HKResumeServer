package main

import (
	"html/template"
	"io"
	"os"

	"github.com/russross/blackfriday"

	"net/http"
)

func main() {
	StartHttpServer()
}

func StartHttpServer() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/resume", resume)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err.Error())
	}
}

func welcome(rw http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	temp.Execute(rw, nil)
}

func resume(rw http.ResponseWriter, req *http.Request) {
	files, err := os.Open("resume.md")
	if err != nil {
		panic(err.Error())
	}

	var bytes []byte = make([]byte, 1024)
	var mdbytes []byte = make([]byte, 0)
	n, err := files.Read(bytes)
	if err != nil && err != io.EOF {
		panic(err.Error())
	}
	for n > 0 {
		mdbytes = append(mdbytes, bytes[:n]...)
		n, err = files.Read(bytes)
		if err != nil && err != io.EOF {
			panic(err.Error())
		}
	}
	httpcontent := blackfriday.MarkdownBasic(mdbytes)
	rw.Write(httpcontent)
}
