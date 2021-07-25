package main

import (
	"fmt"
	"net/http"

	"github.com/mvrilo/go-particles/static"
)

func main() {
	fmt.Println("Started http server at http://127.0.0.1:8000/")
	http.Handle("/", http.FileServer(http.FS(static.Files)))
	http.ListenAndServe(":8000", nil)
}
