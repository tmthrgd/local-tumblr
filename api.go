// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/gregjones/httpcache"
	"golang.org/x/net/proxy"
)

func tumblrBlogURL(blog string) string {
	return "https://api.tumblr.com/v2/blog/" + blog
}

func tumblrPostsURL(blog string, params url.Values) string {
	query := params.Encode()
	if len(query) != 0 {
		query = "?" + query
	}

	return tumblrBlogURL(blog) + "/posts" + query
}

type tumblrApiKeyRoundTripper struct {
	Key       string
	Transport http.RoundTripper
}

func (t *tumblrApiKeyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	tr := t.Transport
	if tr == nil {
		tr = http.DefaultTransport
	}

	if req.URL.Host != "api.tumblr.com" {
		return tr.RoundTrip(req)
	}

	r := *req

	u := *r.URL
	q := u.Query()
	q.Set("api_key", t.Key)
	u.RawQuery = q.Encode()
	r.URL = &u

	return tr.RoundTrip(&r)
}

func getTumblrAPI(apiKey, proxyAddr string) (*http.Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("invalid api key")
	}

	if verbose {
		log.Printf("using api key: %s", apiKey)
	}

	var proxyTr http.RoundTripper

	if len(proxyAddr) != 0 {
		log.Printf("using proxy: socks5://%s", proxyAddr)

		dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			return nil, err
		}

		proxyTr = &http.Transport{
			Dial: dialer.Dial,
		}
	}

	transport := httpcache.NewMemoryCacheTransport()
	transport.MarkCachedResponses = true
	transport.Transport = &tumblrApiKeyRoundTripper{
		Key:       apiKey,
		Transport: proxyTr,
	}

	return &http.Client{
		Transport: transport,
	}, nil
}
