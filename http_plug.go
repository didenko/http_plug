// Copyright 2013 Vlad Didenko
//
// http_plug is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// http_plug is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

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

	respTpl, err = template.New("foo").Parse("Client = {{.RemoteAddr}}\nServer = {{.Host}}\nTime   = {{.Timestamp}}\n\n")
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
