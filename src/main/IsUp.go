package main

import (
	"net"
	"strconv"
	"net/http"
	"strings"
	"fmt"
	"github.com/ogier/pflag"
	"time"
)

func main() {

	var port *int = pflag.IntP("port", "p", 8888, "The port to listen on")
	var timeout *int = pflag.IntP("timeout", "t", 10, "Seconds to wait before giving up")
	pflag.Parse()

	fmt.Printf("Listening on port %d...\n", *port)
	s := service{}
	s.port = *port
	s.timeout = *timeout
	s.listen()
}

type service struct {
	port int
	timeout int
}

func (s service) listen() {
	http.HandleFunc("/", s.handleRequest)
	http.ListenAndServe(":" + strconv.Itoa(s.port), nil)
}

func (s service) handleRequest(w http.ResponseWriter, r *http.Request) {

	if len(r.RequestURI[1:]) == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	pathComponents := strings.Split(r.RequestURI[1:], "/")
	if len(pathComponents) != 2 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	host := pathComponents[0]
	port, err := strconv.Atoi(pathComponents[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isUp := s.testConnection(host, port)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"success\": %t}", isUp)
}

func (s service) testConnection(host string, port int) (bool) {

	timeout := time.Duration(s.timeout) * time.Second

	conn, err := net.DialTimeout("tcp", host + ":" + strconv.Itoa(port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
