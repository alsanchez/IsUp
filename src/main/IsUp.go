package main

import (
	"net"
	"strconv"
	"net/http"
	"strings"
	"fmt"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

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

	isUp := testConnection(host, port)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"success\": %t}", isUp)
}

func testConnection(host string, port int) (bool) {
	conn, err := net.Dial("tcp", host + ":" + strconv.Itoa(port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
