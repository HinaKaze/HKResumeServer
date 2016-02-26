package main

import (
	"html/template"
	"log"
	"os"

	"github.com/russross/blackfriday"

	"net/http"
)

var httpcontent []byte

func main() {
	files, err := os.Open("resume.md")
	if err != nil {
		panic(err.Error())
	}
	var bytes []byte = make([]byte, 1024)
	n, err := files.Read(bytes)
	if err != nil {
		panic(err.Error())
	}
	log.Println(n)
	log.Println(string(bytes))

	//for n, err := files.Read(bytes); n > 0 && err != nil; n, err = files.Read(bytes) {
	//	log.Println(string(bytes))
	//}
	httpcontent = blackfriday.MarkdownBasic(bytes)
	log.Println(string(httpcontent))

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
	rw.Write(httpcontent)
}
