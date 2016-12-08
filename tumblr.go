// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go-bindata -nomemcopy -nocompress assets/... views/...
//go:generate gofmt -w bindata.go
//go:generate ./asset-hashes assets/*

var debug bool
var verbose bool
var images string

func main() {
	flag.BoolVar(&debug, "debug", false, "do not delete temporary files")
	flag.BoolVar(&verbose, "verbose", false, "log more information than normal")

	var addr string
	flag.StringVar(&addr, "addr", ":8080", "the address to listen on")

	var proxy string
	flag.StringVar(&proxy, "proxy", "", "the proxy address to use")

	var dbPath string
	flag.StringVar(&dbPath, "db", "tumblr.db", "path to the database")

	flag.StringVar(&images, "images", "", "the directory to store images in")

	var videos string
	flag.StringVar(&videos, "videos", "", "the directory to store videos in")

	var ro bool
	flag.BoolVar(&ro, "ro", false, "operate in read only mode")

	var apiKey string
	flag.StringVar(&apiKey, "api-key", "XR1x0c5pzhJBZ3yI31GRA0Q3mroKM7g8lP80DZ0EgTmOBmWLdm", "the tumblr api key")

	flag.Parse()

	fmt.Printf("%s (Go runtime %s).\n", fullVersionStr, runtime.Version())
	fmt.Println("Copyright 2016 Tom Thorogood. All rights reserved.")

	api, err := getTumblrAPI(apiKey, proxy)
	if err != nil {
		panic(err)
	}

	db, err := openDB(dbPath, ro)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := getRouter(db, api, images)

	if verbose {
		log.Printf("listening on %s", addr)
	}

	log.Fatal(http.ListenAndServe(addr, router))
}
