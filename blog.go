// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"github.com/tmthrgd/local-tumblr/internal/tumblr"
)

func getBlogHandler(db *bolt.DB) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var cacheControl = fmt.Sprintf("public, max-age=%d", time.Minute/time.Second)

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		blog, before, after := ps.ByName("blog"), ps.ByName("before"), ps.ByName("after")
		if len(blog) == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		h := w.Header()
		h.Set("Cache-Control", cacheControl)

		if checkLastModified(w, r, time.Now(), time.Minute) {
			return
		}

		var last int
		var err error
		if len(before) != 0 {
			last, err = strconv.Atoi(before)
		} else if len(after) != 0 {
			last, err = strconv.Atoi(after)
			last = -last
		}

		if err != nil {
			log.Printf("%[1]T %[1]v", err)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		info, posts, err := getBlog(db, blog, last, 20)
		if err != nil {
			h.Del("Cache-Control")

			if err == errNotFound {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			} else {
				log.Printf("%[1]T %[1]v", err)

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

			return
		}

		if wrote, err := executeTemplate(blogTemplate, struct {
			*tumblr.Blog
			Posts []interface{}
		}{
			Blog:  info,
			Posts: posts,
		}, w); err != nil {
			log.Printf("%[1]T %[1]v", err)

			if !wrote {
				h.Del("Cache-Control")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}
