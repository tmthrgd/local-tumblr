// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func gotoNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	newURL := *r.URL
	newURL.Path = "/"
	newURL.RawQuery = ""

	http.Redirect(w, r, newURL.String(), http.StatusFound)
}

func gotoBlogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	blog := ps.ByName("blog")

	if len(blog) == 0 {
		gotoNotFoundHandler(w, r)
		return
	}

	newURL := *r.URL
	newURL.Path = "/b/" + url.QueryEscape(blog) + "/"
	newURL.RawQuery = ""

	http.Redirect(w, r, newURL.String(), http.StatusFound)
}

func gotoPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	blog, page := ps.ByName("blog"), ps.ByName("page")

	if len(blog) == 0 || len(page) == 0 {
		gotoNotFoundHandler(w, r)
		return
	}

	newURL := *r.URL
	newURL.Path = "/b/" + url.QueryEscape(blog) + "/pg/" + url.QueryEscape(page) + "/"
	newURL.RawQuery = ""

	http.Redirect(w, r, newURL.String(), http.StatusFound)
}

func gotoPostHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	blog, post := ps.ByName("blog"), ps.ByName("post")

	if len(blog) == 0 || len(post) == 0 {
		gotoNotFoundHandler(w, r)
		return
	}

	newURL := *r.URL
	newURL.Path = "/b/" + url.QueryEscape(blog) + "/p/" + url.QueryEscape(post) + "/"
	newURL.RawQuery = ""

	http.Redirect(w, r, newURL.String(), http.StatusFound)
}

func getGotoHandler() func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	router := new(httprouter.Router)

	router.NotFound = http.HandlerFunc(gotoNotFoundHandler)
	router.GET("/:blog/", gotoBlogHandler)
	router.GET("/:blog/page/:page", gotoPageHandler)
	router.GET("/:blog/page/:page/", gotoPageHandler)
	router.GET("/:blog/post/:post", gotoPostHandler)
	router.GET("/:blog/post/:post/", gotoPostHandler)

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Cache-Control", "max-age=0")

		if err := r.ParseForm(); err != nil {
			log.Printf("%[1]T %[1]v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		urlField := r.Form.Get("url")

		if len(urlField) == 0 {
			gotoNotFoundHandler(w, r)
			return
		}

		if !strings.HasPrefix(urlField, "http:") && !strings.HasPrefix(urlField, "https:") {
			urlField = "https://" + urlField
		}

		parsedURL, err := url.Parse(urlField)
		if err != nil {
			gotoNotFoundHandler(w, r)
			return
		}

		req := *r

		reqURL := *r.URL
		reqURL.Path = parsedURL.Host + parsedURL.Path
		req.URL = &reqURL

		router.ServeHTTP(w, &req)
	}
}
