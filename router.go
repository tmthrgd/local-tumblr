// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/julienschmidt/httprouter"
)

func getRouter(db *bolt.DB, api *http.Client, images string) http.Handler {
	router := httprouter.New()

	indexHandler := getIndexHandler(db)
	router.HEAD("/", indexHandler)
	router.GET("/", indexHandler)
	router.GET("/goto/", getGotoHandler())
	downloadHandler := getDownloadHandler(db, api, images)
	router.GET("/d/:blog/p/:page/", downloadHandler)
	router.GET("/d/:blog/u/:until/", downloadHandler)
	blogHandler := getBlogHandler(db)
	router.GET("/b/:blog/", blogHandler)
	router.GET("/b/:blog/b/:before/", blogHandler)
	router.GET("/b/:blog/a/:after/", blogHandler)
	router.GET("/b/:blog/p/:post/", getPostHandler(db))

	assetsRouter := http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,

		Prefix: "assets",
	})
	router.Handler(http.MethodHead, "/favicon.ico", assetsRouter)
	router.Handler(http.MethodGet, "/favicon.ico", assetsRouter)
	router.Handler(http.MethodHead, "/robots.txt", assetsRouter)
	router.Handler(http.MethodGet, "/robots.txt", assetsRouter)

	router.ServeFiles("/assets/*filepath", &assetfs.AssetFS{
		Asset:     AssetFromNameHash,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfoFromNameHash,

		Prefix: "assets",
	})

	router.ServeFiles("/images/*filepath", http.Dir(images))

	errRouter := errorHandler{router}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", fullVersionStr)
		errRouter.ServeHTTP(w, r)
	})
}
