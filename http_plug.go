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
	"os"
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

// PlugServer is a handler which writes the timestamp and some
// utility info back to the requester. It also logs the request.
func PlugServer(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		return
	}

	r := resp{*req, time.Now().UTC().Format("15:04:05 MST")}
	err := respTpl.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}

	respTpl.Execute(os.Stdout, r)
}

func main() {
	http.HandleFunc("/", PlugServer)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
