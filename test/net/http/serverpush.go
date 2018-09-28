package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

// preparing image
func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// Send HTML and push image
func handlerHtml(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		fmt.Println("Push /image")
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

// Send image as usual HTTP request
func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}
func main() {
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "cert.pem", "key.pem", nil)
	err = http.ListenAndServe(":18443", nil)
	fmt.Println(err)
}
