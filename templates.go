// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tmthrgd/local-tumblr/internal/tumblr"
)

var (
	templateFuncs = map[string]interface{}{
		"asset_path": assetPath,
		"truncate":   truncate,
		"sub":        sub,
		"raw_html":   rawHTML,
		"image_path": imagePath,
	}

	indexTemplate = template.Must(template.New("index.tmpl").Funcs(templateFuncs).Parse(string(MustAsset("views/index.tmpl"))))
	errorTemplate = template.Must(template.New("error.tmpl").Funcs(templateFuncs).Parse(string(MustAsset("views/error.tmpl"))))
	blogTemplate  = template.Must(template.New("blog.tmpl").Funcs(templateFuncs).Parse(string(MustAsset("views/blog.tmpl"))))
)

func assetPath(name string) (string, error) {
	if strings.HasPrefix(name, "http://") || strings.HasPrefix(name, "https://") || strings.HasPrefix(name, "//") {
		return name, nil
	}

	name, err := AssetHashName(filepath.Join("assets/", name))
	if err != nil {
		return "", err
	}

	return "/" + name, nil
}

func truncate(value string, length int) string {
	numRunes := 0

	for index := range value {
		numRunes++

		if numRunes > length {
			return value[:index]
		}
	}

	return value
}

func sub(a, b int) int {
	return a - b
}

func rawHTML(v string) template.HTML {
	return template.HTML(v)
}

func imagePath(u tumblr.URL) string {
	_, name := path.Split(u.Path)

	if _, err := os.Stat(filepath.Join(images, name)); os.IsNotExist(err) {
		return u.String()
	}

	return path.Join("/images", name)
}

var indexModTime time.Time

func init() {
	if stat, err := AssetInfo("views/index.tmpl"); err == nil {
		indexModTime = stat.ModTime()
	} else {
		panic(err)
	}
}

func executeTemplate(tmpl *template.Template, data interface{}, w http.ResponseWriter) (didWrite bool, err error) {
	return executeTemplateWithCode(tmpl, data, w, http.StatusOK)
}

func executeTemplateWithCode(tmpl *template.Template, data interface{}, w http.ResponseWriter, code int) (didWrite bool, err error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)
	buf.Reset()

	if err = tmpl.Execute(buf, data); err != nil {
		return false, err
	}

	h := w.Header()
	h.Set("Content-Length", strconv.FormatInt(int64(buf.Len()), 10))
	h.Set("Content-Type", "text/html; charset=utf-8")

	w.WriteHeader(code)

	_, err = buf.WriteTo(w)
	return true, err
}
