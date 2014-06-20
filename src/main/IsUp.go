package main

import (
	"net"
	"strconv"
	"net/http"
	"fmt"
	"github.com/ogier/pflag"
	"github.com/gorilla/mux"
	"time"
)

func main() {

	var port *int = pflag.IntP("port", "p", 8888, "The port to listen on")
	var timeout *int = pflag.IntP("defaul-timeout", "t", 10, "Seconds to wait before giving up on the request")
	pflag.Parse()

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
	router := mux.NewRouter()
	router.HandleFunc("/{host}/{port:[0-9]+}", s.handleRequest).Methods("GET")

	fmt.Printf("Listening on port %d with a default timeout of %d seconds...\n", s.port, s.timeout)
	http.Handle("/", router)
	err := http.ListenAndServe(":" + strconv.Itoa(s.port), nil)
	if err == nil {
		fmt.Println(err)
	}
}

func (s service) handleRequest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	host := params["host"]
	timeout := s.timeout

	if r.FormValue("timeout") != "" {
		timeout, _ = strconv.Atoi(r.FormValue("timeout"))
	}

	port, err := strconv.Atoi(params["port"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isUp := s.testConnection(host, port, timeout)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"success\": %t}", isUp)
}

func (s service) testConnection(host string, port int, timeout int) (bool) {

	timeout_duration := time.Duration(timeout) * time.Second

	conn, err := net.DialTimeout("tcp", host + ":" + strconv.Itoa(port), timeout_duration)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
