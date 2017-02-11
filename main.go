package main

import (
	"fmt"
	"net/http"
	"os"
)

// Manning go web programming - page:83
func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	dir, _ := os.Getwd()
	//http.Handle("/", http.FileServer(http.Dir(dir)))

	var port string
	arg := os.Args
	if len(arg) == 1 {
		port = ":8000"
	} else {
		port = ":" + arg[1]
	}

	fmt.Printf("From: %s\nListening localhost%s\n", dir, port)

	server := http.Server{
		Addr: port,
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", log(fs))

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
