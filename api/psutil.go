package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/ghostsck/tools/monitor"
)

func main() {
	port := flag.String("port", ":58000", "HTTP listen port")
	flag.Parse()
	http.HandleFunc("/", getInfo)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, monitor.InfoMiniJSON())
}
