// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/tmthrgd/local-tumblr/internal/tumblr"
)

var (
	bucketBlogs = []byte("blogs")
	bucketPosts = []byte("posts")
)

const bucketPostsPrefix = "posts-"

var errNotFound = errors.New("not found")

func openDB(path string, ro bool) (*bolt.DB, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{
		ReadOnly: ro,
	})
	if err != nil {
		return nil, err
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucket(bucketBlogs); err != nil && err != bolt.ErrBucketExists {
			return err
		}

		if _, err := tx.CreateBucket(bucketPosts); err != nil && err != bolt.ErrBucketExists {
			return err
		}

		return nil
	}); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

type blog struct {
	tumblr.Blog

	LastPost int
}

func getBlogs(db *bolt.DB) ([]blog, error) {
	blogs := make([]blog, 0, 10)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketBlogs)
		c := b.Cursor()

		pb := tx.Bucket(bucketPosts)

		for k, v := c.First(); k != nil; k, v = c.Next() {
			blogs = append(blogs, blog{})
			blog := &blogs[len(blogs)-1]

			if err := json.Unmarshal(v, &blog.Blog); err != nil {
				return err
			}

			bb := pb.Bucket([]byte(blog.Name))
			bc := bb.Cursor()
			_, vv := bc.Last()
			if vv == nil {
				continue
			}

			var post tumblr.BasePost
			if err := json.Unmarshal(vv, &post); err != nil {
				return err
			}

			blog.LastPost = post.Id
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func getBlog(db *bolt.DB, blog string, last, count int) (*tumblr.Blog, []interface{}, error) {
	var info tumblr.Blog
	posts := make([]interface{}, count)

	var gotPosts int

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketBlogs)
		bv := b.Get([]byte(blog))
		if bv == nil {
			return errNotFound
		}

		err := json.Unmarshal(bv, &info)
		if err != nil {
			return err
		}

		b = tx.Bucket(bucketPosts)
		b = b.Bucket([]byte(blog))
		c := b.Cursor()

		var k, v []byte
		var next func() ([]byte, []byte)

		switch {
		case last > 0:
			var id [8]byte
			binary.BigEndian.PutUint64(id[:], uint64(last))

			k, v = c.Seek(id[:])

			if bytes.Equal(k, id[:]) {
				k, v = c.Prev()
			}

			next = c.Prev
		case last < 0:
			var id [8]byte
			binary.BigEndian.PutUint64(id[:], uint64(-last))

			k, v = c.Seek(id[:])

			if bytes.Equal(k, id[:]) {
				k, v = c.Next()
			}

			next = c.Next
		case last == 0:
			k, v = c.Last()
			next = c.Prev
		}

		for ; k != nil && gotPosts < count; k, v = next() {
			var post *interface{}

			if last < 0 {
				post = &posts[count-1-gotPosts]
			} else {
				post = &posts[gotPosts]
			}

			if err := tumblr.UnmarshalPost(v, post); err != nil {
				return err
			}

			gotPosts++
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	if last < 0 {
		return &info, posts[count-gotPosts:], nil
	} else {
		return &info, posts[:gotPosts], nil
	}
}

func storeBlog(db *bolt.DB, blog *tumblr.Blog) error {
	return db.Update(func(tx *bolt.Tx) error {
		v, err := json.Marshal(blog)
		if err != nil {
			return err
		}

		b := tx.Bucket(bucketBlogs)
		if err := b.Put([]byte(blog.Name), v); err != nil {
			return err
		}

		b = tx.Bucket(bucketPosts)
		_, err = b.CreateBucketIfNotExists([]byte(blog.Name))
		return err
	})
}

func storePost(db *bolt.DB, blog string, post interface{}) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketPosts)
		b = b.Bucket([]byte(blog))

		v, err := json.Marshal(post)
		if err != nil {
			return err
		}

		base := tumblr.ToBasePost(post)

		var id [8]byte
		binary.BigEndian.PutUint64(id[:], uint64(base.Id))

		return b.Put(id[:], v)
	})
}
