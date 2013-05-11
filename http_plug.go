package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

const addrTpl = ":%d"

var respTpl *template.Template
var listenAddr string

type resp struct {
	http.Request
	Timestamp string
}

func init() {
	flag.Parse()

	port64, err := strconv.ParseUint(flag.Arg(0), 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	if port64 > 65535 {
		log.Fatal("Port nummber must ne in the [1..65535] range.")
	}

	listenAddr = fmt.Sprintf(addrTpl, uint16(port64))

	respTpl, err = template.New("foo").Parse(`Time = {{.Timestamp}}
Host = {{.Host}}
RemoteAddr = {{.RemoteAddr}}
RequestURI = {{.RequestURI}}
`)
}

func PlugServer(w http.ResponseWriter, req *http.Request) {
	r := resp{*req, time.Now().Format("20060102150405")}
	err := respTpl.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", PlugServer)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
