package main

import (
	"fmt"
	"net"
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

	ip := getIPAddress()
	fmt.Printf("From: %s\nListening: %s%s\n", dir, ip, port)

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

// get ip address
func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "Cann't get IP address!"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
