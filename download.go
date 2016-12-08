// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"github.com/tmthrgd/local-tumblr/internal/tumblr"
)

var (
	errHttpNotFound = errors.New("not found")
	errOutOfPosts   = errors.New("out of posts")
)

func getDownloadHandler(db *bolt.DB, api *http.Client, images string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	download := func(w http.ResponseWriter, r *http.Request, blog string, num, until int) error {
		if verbose {
			log.Printf("downloading page %d of %s", num, blog)
		}

		params := make(url.Values)
		params.Add("offset", strconv.Itoa((num-1)*20))
		params.Add("limit", "20")

		resp, err := api.Get(tumblrPostsURL(blog, params))
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()

		var tresp tumblr.Response
		if err := decoder.Decode(&tresp); err != nil {
			return err
		}

		switch tresp.Meta.Status {
		case http.StatusOK:
		case http.StatusNotFound:
			return errHttpNotFound
		default:
			return fmt.Errorf("%d: %s", tresp.Meta.Status, tresp.Meta.Msg)
		}

		var presp tumblr.PostsResponse
		if err := json.Unmarshal(tresp.Response, &presp); err != nil {
			return err
		}

		if err := storeBlog(db, &presp.Blog); err != nil {
			return err
		}

		if len(presp.Posts) == 0 {
			return errOutOfPosts
		}

		for _, rawPost := range presp.Posts {
			var post interface{}
			if err := tumblr.UnmarshalPost(rawPost, &post); err != nil {
				return err
			}

			switch post := post.(type) {
			case *tumblr.PhotoPost:
				var wg sync.WaitGroup
				wg.Add(len(post.Photos))

				for i := range post.Photos {
					photo := &post.Photos[i]
					go downloadPhoto(api, images, photo, &wg)
				}

				wg.Wait()
			}

			if err := storePost(db, presp.Blog.Name, post); err != nil {
				return err
			}

			if until != 0 {
				if base := tumblr.ToBasePost(post); base.Id <= until {
					return errOutOfPosts
				}
			}
		}

		return nil
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		blog, page, until := ps.ByName("blog"), ps.ByName("page"), ps.ByName("until")
		if len(blog) == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if len(until) != 0 {
			id, err := strconv.Atoi(until)
			if err != nil {
				log.Printf("%[1]T %[1]v", err)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

		pageLoop:
			for page := 1; ; page++ {
				switch err := download(w, r, blog, page, id); err {
				case nil:
				case errOutOfPosts:
					break pageLoop
				case errHttpNotFound:
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
					return
				default:
					log.Printf("%[1]T %[1]v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}

			return
		}

		if idx := strings.Index(page, ".."); idx != -1 {
			start, _, err := parsePageString(page[:idx])
			if err != nil {
				log.Printf("%[1]T %[1]v", err)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			end, _, err := parsePageString(page[idx+2:])
			if err != nil {
				log.Printf("%[1]T %[1]v", err)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			errors := make([]error, end-start+1)
			var wg sync.WaitGroup
			wg.Add(len(errors))

			for page := start; page <= end; page++ {
				go func(page int) {
					errors[page-start] = download(w, r, blog, page, 0)
					wg.Done()
				}(page)
			}

			wg.Wait()

			var did int

			for _, err := range errors {
				switch err {
				case nil:
					did++
				case errOutOfPosts:
					break
				case errHttpNotFound:
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
					return
				default:
					log.Printf("%[1]T %[1]v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}

			if did < 5 {
				return
			}

			switch err := download(w, r, blog, start, 0); err {
			case nil:
			case errHttpNotFound:
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				log.Printf("%[1]T %[1]v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

			return
		}

		num, _, err := parsePageString(page)
		if err != nil {
			log.Printf("%[1]T %[1]v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		switch err := download(w, r, blog, num, 0); err {
		case nil:
		case errHttpNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			log.Printf("%[1]T %[1]v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func downloadPhoto(api *http.Client, images string, photo *tumblr.PostPhoto, wg *sync.WaitGroup) {
	defer wg.Done()

	_, name := path.Split(photo.OriginalSize.Url.Path)
	path := filepath.Join(images, name)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		if !os.IsExist(err) {
			log.Printf("%[1]T %[1]v", err)
		}

		return
	}
	defer f.Close()

	if verbose {
		log.Printf("downloading image %s (%s)", name, photo.OriginalSize.Url)
	}

	for tries := 3; tries > 0; tries-- {
		resp, err := api.Do(&http.Request{
			Method: http.MethodGet,

			URL: &photo.OriginalSize.Url.URL,
		})
		if err != nil {
			log.Printf("%[1]T %[1]v", err)
			continue
		}

		if err := f.Truncate(0); err != nil {
			log.Printf("%[1]T %[1]v", err)
			break
		}

		if _, err := copyBuffer(f, resp.Body); err != nil {
			log.Printf("%[1]T %[1]v", err)
			continue
		}

		if debug {
			log.Printf("download of image %s succeeded", name)
		}

		return
	}

	log.Printf("download of image %s failed", name)

	f.Close()
	os.Remove(path)
}
