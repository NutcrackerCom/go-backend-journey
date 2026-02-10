package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
	read, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Println(string(read))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9091", nil)
}
