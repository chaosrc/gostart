package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	f, err := os.OpenFile("./log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(f, "geo", 0)
}

func main() {
	http.HandleFunc("/", static)
	http.HandleFunc("/post", post)
	log.Fatalln(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}

func static(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./index.html")
	if err != nil {
		http.Error(w, "Open file error: %s", http.StatusNoContent)
	}
	defer file.Close()
	io.Copy(w, file)
}

func post(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Write file error", http.StatusNoContent)
	}
	defer r.Body.Close()
	defer file.Close()
	io.Copy(file, r.Body)
	logger.Println(r.URL.Path)
}
