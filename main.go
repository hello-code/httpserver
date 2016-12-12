package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(dir)))

	var port string
	arg := os.Args
	if len(arg) == 1 {
		port = ":8000"
	} else {
		port = ":" + arg[1]
	}

	fmt.Printf("From: %s\nListening localhost%s", dir, port)
	server := http.Server{
		Addr: port,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
