package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Something pinged me!")
	io.WriteString(w, "this is yash!\n")
}

func handler2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Something pinged me!")
	io.WriteString(w, "this is yash 2!\n")
}

func main() {

	http.HandleFunc("/hello", handler)
	http.HandleFunc("/", handler2)

	http.ListenAndServe(":9900", nil)

}
