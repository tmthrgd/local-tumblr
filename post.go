// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
)

func getPostHandler(db *bolt.DB) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var cacheControl = fmt.Sprintf("public, max-age=%d", time.Minute/time.Second)

	var _ = cacheControl
	var _ = log.Printf

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
