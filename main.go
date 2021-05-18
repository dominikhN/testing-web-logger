package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getHostname() (hostname string, err error) {
	hostname, err = os.Hostname()
	return
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ %s ] [ %s ] : %s\n", r.RemoteAddr, r.Method, r.URL)
	hostname, err := getHostname()
	if err != nil {
		fmt.Println("Cannot get hostname: ", err)
		fmt.Fprint(w, "OK\n", r.URL.Path[1:])
		return
	}
	fmt.Fprintf(w, "( %s ) OK\n", hostname)
}
func handlerBase(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ %s ] [ %s ] : %s\n", r.RemoteAddr, r.Method, r.URL)
	http.Redirect(w, r, "/health", http.StatusFound)
}

func main() {
	portPtr := flag.Int("port", 8443, "port to listen on")
	flag.Parse()
	port := ":" + strconv.Itoa(*portPtr)

	http.HandleFunc("/health", handlerHealth)
	http.HandleFunc("/", handlerBase)
	log.Println("registered handlers")
	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
