// Code generated by go-bindata.
// sources:
// assets/robots.txt
// assets/style.css
// views/blog.tmpl
// views/error.tmpl
// views/index.tmpl
// DO NOT EDIT!

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len
	return b, nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsRobotsTxt = "\x55\x73\x65\x72\x2d\x61\x67\x65\x6e\x74\x3a\x20\x2a\x0a\x44\x69\x73\x61\x6c\x6c\x6f\x77\x3a\x20\x2f\x0a"

func assetsRobotsTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsRobotsTxt,
		"assets/robots.txt",
	)
}

func assetsRobotsTxt() (*asset, error) {
	bytes, err := assetsRobotsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/robots.txt", size: 26, mode: os.FileMode(436), modTime: time.Unix(1467431681, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsStyleCss = "\x2f\x2a\x21\x20\x68\x74\x74\x70\x3a\x2f\x2f\x62\x65\x74\x74\x65\x72\x6d\x6f\x74\x68\x65\x72\x66\x75\x63\x6b\x69\x6e\x67\x77\x65\x62\x73\x69\x74\x65\x2e\x63\x6f\x6d\x2f\x20\x2a\x2f\x0a\x62\x6f\x64\x79\x20\x7b\x0a\x09\x6d\x61\x72\x67\x69\x6e\x3a\x20\x34\x30\x70\x78\x20\x61\x75\x74\x6f\x3b\x0a\x09\x70\x61\x64\x64\x69\x6e\x67\x3a\x30\x20\x31\x30\x70\x78\x3b\x0a\x0a\x09\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x37\x30\x30\x70\x78\x3b\x0a\x0a\x09\x6c\x69\x6e\x65\x2d\x68\x65\x69\x67\x68\x74\x3a\x20\x31\x2e\x36\x3b\x0a\x09\x66\x6f\x6e\x74\x2d\x73\x69\x7a\x65\x3a\x20\x31\x38\x70\x78\x3b\x0a\x0a\x09\x63\x6f\x6c\x6f\x72\x3a\x20\x23\x34\x34\x34\x3b\x0a\x7d\x0a\x0a\x68\x31\x2c\x20\x68\x32\x2c\x20\x68\x33\x20\x7b\x0a\x09\x6c\x69\x6e\x65\x2d\x68\x65\x69\x67\x68\x74\x3a\x20\x31\x2e\x32\x3b\x0a\x7d\x0a\x0a\x2f\x2a\x21\x20\x68\x74\x74\x70\x73\x3a\x2f\x2f\x72\x61\x77\x67\x69\x74\x2e\x63\x6f\x6d\x2f\x20\x2a\x2f\x0a\x2e\x6f\x66\x66\x73\x63\x72\x65\x65\x6e\x20\x7b\x0a\x09\x70\x6f\x73\x69\x74\x69\x6f\x6e\x3a\x20\x61\x62\x73\x6f\x6c\x75\x74\x65\x3b\x0a\x09\x6c\x65\x66\x74\x3a\x20\x2d\x39\x39\x39\x39\x70\x78\x3b\x0a\x7d\x0a\x0a\x69\x6e\x70\x75\x74\x2e\x75\x72\x6c\x20\x7b\x0a\x09\x77\x69\x64\x74\x68\x3a\x20\x31\x30\x30\x25\x3b\x0a\x0a\x09\x70\x61\x64\x64\x69\x6e\x67\x3a\x20\x30\x20\x36\x70\x78\x3b\x0a\x0a\x09\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x23\x66\x62\x66\x62\x66\x62\x3b\x0a\x09\x62\x6f\x72\x64\x65\x72\x3a\x20\x74\x68\x69\x6e\x20\x73\x6f\x6c\x69\x64\x20\x23\x64\x66\x64\x66\x64\x66\x3b\x0a\x0a\x09\x66\x6f\x6e\x74\x2d\x73\x69\x7a\x65\x3a\x20\x31\x2e\x33\x65\x6d\x3b\x0a\x09\x74\x65\x78\x74\x2d\x61\x6c\x69\x67\x6e\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x0a\x09\x63\x75\x72\x73\x6f\x72\x3a\x20\x74\x65\x78\x74\x3b\x0a\x7d\x0a\x0a\x2f\x2a\x21\x0a\x20\x2a\x20\x43\x6f\x70\x79\x72\x69\x67\x68\x74\x20\x32\x30\x31\x36\x20\x54\x6f\x6d\x20\x54\x68\x6f\x72\x6f\x67\x6f\x6f\x64\x2e\x20\x41\x6c\x6c\x20\x72\x69\x67\x68\x74\x73\x20\x72\x65\x73\x65\x72\x76\x65\x64\x2e\x0a\x20\x2a\x20\x55\x73\x65\x20\x6f\x66\x20\x74\x68\x69\x73\x20\x73\x6f\x75\x72\x63\x65\x20\x63\x6f\x64\x65\x20\x69\x73\x20\x67\x6f\x76\x65\x72\x6e\x65\x64\x20\x62\x79\x20\x61\x0a\x20\x2a\x20\x4d\x6f\x64\x69\x66\x69\x65\x64\x20\x42\x53\x44\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x6c\x69\x63\x65\x6e\x73\x65\x20\x74\x68\x61\x74\x20\x63\x61\x6e\x20\x62\x65\x20\x66\x6f\x75\x6e\x64\x20\x69\x6e\x0a\x20\x2a\x20\x74\x68\x65\x20\x4c\x49\x43\x45\x4e\x53\x45\x20\x66\x69\x6c\x65\x2e\x0a\x20\x2a\x2f\x0a\x75\x6c\x20\x7b\x0a\x09\x6c\x69\x73\x74\x2d\x73\x74\x79\x6c\x65\x2d\x74\x79\x70\x65\x3a\x20\x73\x71\x75\x61\x72\x65\x3b\x0a\x7d\x0a\x0a\x70\x72\x65\x20\x7b\x0a\x09\x77\x6f\x72\x64\x2d\x77\x72\x61\x70\x3a\x20\x62\x72\x65\x61\x6b\x2d\x77\x6f\x72\x64\x3b\x0a\x09\x77\x68\x69\x74\x65\x2d\x73\x70\x61\x63\x65\x3a\x20\x70\x72\x65\x2d\x77\x72\x61\x70\x3b\x0a\x7d\x0a\x0a\x0a\x2e\x70\x6f\x73\x74\x20\x7b\x0a\x09\x62\x6f\x72\x64\x65\x72\x2d\x62\x6f\x74\x74\x6f\x6d\x3a\x20\x74\x68\x69\x6e\x20\x73\x6f\x6c\x69\x64\x20\x23\x65\x65\x65\x3b\x0a\x09\x70\x61\x64\x64\x69\x6e\x67\x3a\x20\x35\x70\x78\x20\x30\x3b\x0a\x09\x6d\x61\x72\x67\x69\x6e\x3a\x20\x31\x30\x70\x78\x20\x30\x3b\x0a\x7d\x0a\x0a\x2e\x70\x6f\x73\x74\x2e\x70\x68\x6f\x74\x6f\x20\x69\x6d\x67\x20\x7b\x0a\x09\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x31\x30\x30\x25\x3b\x0a\x7d\x0a"

func assetsStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsStyleCss,
		"assets/style.css",
	)
}

func assetsStyleCss() (*asset, error) {
	bytes, err := assetsStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/style.css", size: 825, mode: os.FileMode(436), modTime: time.Unix(1467465676, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsBlogTmpl = "\x3c\x21\x64\x6f\x63\x74\x79\x70\x65\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x20\x6c\x61\x6e\x67\x3d\x65\x6e\x3e\x0a\x3c\x68\x65\x61\x64\x3e\x0a\x09\x3c\x6d\x65\x74\x61\x20\x63\x68\x61\x72\x73\x65\x74\x3d\x75\x74\x66\x2d\x38\x3e\x0a\x09\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x76\x69\x65\x77\x70\x6f\x72\x74\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x77\x69\x64\x74\x68\x3d\x64\x65\x76\x69\x63\x65\x2d\x77\x69\x64\x74\x68\x2c\x69\x6e\x69\x74\x69\x61\x6c\x2d\x73\x63\x61\x6c\x65\x3d\x31\x22\x3e\x0a\x09\x3c\x74\x69\x74\x6c\x65\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x20\xc2\xb7\x20\x6c\x6f\x63\x61\x6c\x2d\x74\x75\x6d\x62\x6c\x72\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x09\x3c\x6c\x69\x6e\x6b\x20\x72\x65\x6c\x3d\x73\x74\x79\x6c\x65\x73\x68\x65\x65\x74\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x61\x73\x73\x65\x74\x5f\x70\x61\x74\x68\x20\x22\x73\x74\x79\x6c\x65\x2e\x63\x73\x73\x22\x7d\x7d\x22\x3e\x0a\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x3c\x62\x6f\x64\x79\x3e\x0a\x09\x3c\x68\x65\x61\x64\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x73\x69\x74\x65\x2d\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x3c\x68\x31\x3e\x3c\x61\x20\x68\x72\x65\x66\x3d\x2f\x3e\x6c\x6f\x63\x61\x6c\x2d\x74\x75\x6d\x62\x6c\x72\x3c\x2f\x61\x3e\x3c\x2f\x68\x31\x3e\x0a\x09\x09\x3c\x68\x32\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x2e\x55\x72\x6c\x2e\x53\x74\x72\x69\x6e\x67\x7d\x7d\x22\x20\x74\x69\x74\x6c\x65\x3d\x22\x56\x69\x65\x77\x20\x6f\x6e\x20\x54\x75\x6d\x62\x6c\x72\x22\x3e\xe2\xa4\xb4\x3c\x2f\x61\x3e\x3c\x2f\x68\x32\x3e\x0a\x09\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x0a\x09\x3c\x6d\x61\x69\x6e\x3e\x0a\x09\x7b\x7b\x2d\x20\x72\x61\x6e\x67\x65\x20\x2e\x50\x6f\x73\x74\x73\x7d\x7d\x0a\x09\x09\x3c\x61\x72\x74\x69\x63\x6c\x65\x20\x63\x6c\x61\x73\x73\x3d\x22\x70\x6f\x73\x74\x20\x7b\x7b\x2e\x54\x79\x70\x65\x7d\x7d\x22\x3e\x0a\x09\x09\x7b\x7b\x2d\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x74\x65\x78\x74\x22\x7d\x7d\x0a\x09\x09\x09\x3c\x68\x65\x61\x64\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x09\x09\x3c\x68\x32\x20\x63\x6c\x61\x73\x73\x3d\x74\x69\x74\x6c\x65\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x3c\x2f\x68\x32\x3e\x0a\x09\x09\x09\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x2e\x42\x6f\x64\x79\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x70\x68\x6f\x74\x6f\x22\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x7b\x7b\x2d\x20\x69\x66\x20\x65\x71\x20\x28\x6c\x65\x6e\x20\x2e\x50\x68\x6f\x74\x6f\x73\x29\x20\x31\x7d\x7d\x0a\x09\x09\x09\x7b\x7b\x77\x69\x74\x68\x20\x24\x69\x6d\x67\x20\x3a\x3d\x20\x28\x69\x6e\x64\x65\x78\x20\x2e\x50\x68\x6f\x74\x6f\x73\x20\x30\x29\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x09\x3c\x69\x6d\x67\x20\x73\x72\x63\x3d\x22\x7b\x7b\x24\x69\x6d\x67\x2e\x4f\x72\x69\x67\x69\x6e\x61\x6c\x53\x69\x7a\x65\x2e\x55\x72\x6c\x20\x7c\x20\x69\x6d\x61\x67\x65\x5f\x70\x61\x74\x68\x7d\x7d\x22\x20\x61\x6c\x74\x3d\x22\x7b\x7b\x2e\x43\x61\x70\x74\x69\x6f\x6e\x7d\x7d\x22\x3e\x0a\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x63\x61\x70\x74\x69\x6f\x6e\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x2e\x43\x61\x70\x74\x69\x6f\x6e\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6c\x73\x65\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x72\x61\x6e\x67\x65\x20\x2e\x50\x68\x6f\x74\x6f\x73\x7d\x7d\x0a\x09\x09\x09\x09\x09\x3c\x69\x6d\x67\x20\x73\x72\x63\x3d\x22\x7b\x7b\x2e\x4f\x72\x69\x67\x69\x6e\x61\x6c\x53\x69\x7a\x65\x2e\x55\x72\x6c\x20\x7c\x20\x69\x6d\x61\x67\x65\x5f\x70\x61\x74\x68\x7d\x7d\x22\x20\x74\x69\x74\x6c\x65\x3d\x22\x7b\x7b\x2e\x43\x61\x70\x74\x69\x6f\x6e\x7d\x7d\x22\x20\x61\x6c\x74\x3d\x22\x7b\x7b\x2e\x43\x61\x70\x74\x69\x6f\x6e\x7d\x7d\x22\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x71\x75\x6f\x74\x65\x22\x7d\x7d\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x3e\x0a\x09\x09\x09\x09\x3c\x62\x6c\x6f\x63\x6b\x71\x75\x6f\x74\x65\x3e\x7b\x7b\x2e\x54\x65\x78\x74\x7d\x7d\x3c\x2f\x62\x6c\x6f\x63\x6b\x71\x75\x6f\x74\x65\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x69\x66\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x09\x09\x3c\x70\x20\x63\x6c\x61\x73\x73\x3d\x73\x6f\x75\x72\x63\x65\x3e\x7b\x7b\x2e\x53\x6f\x75\x72\x63\x65\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x3c\x2f\x70\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x6c\x69\x6e\x6b\x22\x7d\x7d\x0a\x09\x09\x09\x3c\x68\x65\x61\x64\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x09\x09\x3c\x68\x32\x20\x63\x6c\x61\x73\x73\x3d\x6c\x69\x6e\x6b\x3e\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x2e\x55\x72\x6c\x2e\x53\x74\x72\x69\x6e\x67\x7d\x7d\x22\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x3c\x2f\x61\x3e\x3c\x2f\x68\x32\x3e\x0a\x09\x09\x09\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x2e\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x63\x68\x61\x74\x22\x7d\x7d\x0a\x09\x09\x09\x7b\x7b\x69\x66\x20\x2e\x54\x69\x74\x6c\x65\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x3c\x68\x65\x61\x64\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x09\x09\x3c\x68\x32\x20\x63\x6c\x61\x73\x73\x3d\x74\x69\x74\x6c\x65\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x3c\x2f\x68\x32\x3e\x0a\x09\x09\x09\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x0a\x09\x09\x09\x3c\x75\x6c\x20\x63\x6c\x61\x73\x73\x3d\x22\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x20\x64\x69\x61\x6c\x6f\x67\x75\x65\x22\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x72\x61\x6e\x67\x65\x20\x2e\x44\x69\x61\x6c\x6f\x67\x75\x65\x7d\x7d\x0a\x09\x09\x09\x09\x3c\x6c\x69\x3e\x0a\x09\x09\x09\x09\x09\x7b\x7b\x69\x66\x20\x2e\x4c\x61\x62\x65\x6c\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x09\x09\x09\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x63\x68\x61\x74\x2d\x6c\x61\x62\x65\x6c\x3e\x7b\x7b\x2e\x4c\x61\x62\x65\x6c\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x3c\x2f\x73\x70\x61\x6e\x3e\x0a\x09\x09\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x09\x09\x09\x3c\x71\x20\x63\x6c\x61\x73\x73\x3d\x70\x68\x72\x61\x73\x65\x3e\x7b\x7b\x2e\x50\x68\x72\x61\x73\x65\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x3c\x2f\x71\x3e\x0a\x09\x09\x09\x09\x3c\x2f\x6c\x69\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x09\x3c\x2f\x75\x6c\x3e\x0a\x09\x09\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x61\x75\x64\x69\x6f\x22\x7d\x7d\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x2e\x50\x6c\x61\x79\x65\x72\x7d\x7d\x0a\x0a\x09\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x63\x61\x70\x74\x69\x6f\x6e\x3e\x0a\x09\x09\x09\x09\x09\x7b\x7b\x2e\x43\x61\x70\x74\x69\x6f\x6e\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x0a\x09\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x76\x69\x64\x65\x6f\x22\x7d\x7d\x0a\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x70\x6f\x73\x74\x2d\x62\x6f\x64\x79\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x28\x69\x6e\x64\x65\x78\x20\x2e\x50\x6c\x61\x79\x65\x72\x20\x28\x73\x75\x62\x20\x28\x6c\x65\x6e\x20\x2e\x50\x6c\x61\x79\x65\x72\x29\x20\x31\x29\x29\x2e\x45\x6d\x62\x65\x64\x43\x6f\x64\x65\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x0a\x0a\x09\x09\x09\x09\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x63\x61\x70\x74\x69\x6f\x6e\x3e\x0a\x09\x09\x09\x09\x09\x7b\x7b\x2e\x43\x61\x70\x74\x69\x6f\x6e\x20\x7c\x20\x72\x61\x77\x5f\x68\x74\x6d\x6c\x7d\x7d\x0a\x09\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x09\x3c\x2f\x64\x69\x76\x3e\x0a\x09\x09\x7b\x7b\x65\x6e\x64\x20\x2d\x7d\x7d\x0a\x09\x09\x3c\x2f\x61\x72\x74\x69\x63\x6c\x65\x3e\x0a\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x0a\x09\x7b\x7b\x2d\x20\x69\x66\x20\x6e\x65\x20\x28\x6c\x65\x6e\x20\x2e\x50\x6f\x73\x74\x73\x29\x20\x30\x7d\x7d\x0a\x0a\x09\x09\x3c\x66\x6f\x6f\x74\x65\x72\x3e\x0a\x09\x09\x09\x3c\x70\x3e\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x2f\x62\x2f\x7b\x7b\x2e\x4e\x61\x6d\x65\x7d\x7d\x2f\x61\x2f\x7b\x7b\x28\x69\x6e\x64\x65\x78\x20\x2e\x50\x6f\x73\x74\x73\x20\x30\x29\x2e\x49\x64\x7d\x7d\x2f\x22\x3e\xe2\x86\x90\x20\x50\x72\x65\x76\x20\x70\x61\x67\x65\x3c\x2f\x61\x3e\x20\xc2\xb7\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x2f\x62\x2f\x7b\x7b\x2e\x4e\x61\x6d\x65\x7d\x7d\x2f\x62\x2f\x7b\x7b\x28\x69\x6e\x64\x65\x78\x20\x2e\x50\x6f\x73\x74\x73\x20\x28\x73\x75\x62\x20\x28\x6c\x65\x6e\x20\x2e\x50\x6f\x73\x74\x73\x29\x20\x31\x29\x29\x2e\x49\x64\x7d\x7d\x2f\x22\x3e\x4e\x65\x78\x74\x20\x70\x61\x67\x65\x20\xe2\x86\x92\x3c\x2f\x61\x3e\x3c\x2f\x70\x3e\x0a\x09\x09\x3c\x2f\x66\x6f\x6f\x74\x65\x72\x3e\x0a\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x3c\x2f\x6d\x61\x69\x6e\x3e\x0a\x0a\x09\x3c\x66\x6f\x6f\x74\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x73\x69\x74\x65\x2d\x66\x6f\x6f\x74\x65\x72\x3e\x0a\x09\x09\x3c\x70\x3e\xc2\xa9\x20\x32\x30\x31\x36\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x68\x74\x74\x70\x73\x3a\x2f\x2f\x74\x6f\x6d\x74\x68\x6f\x72\x6f\x67\x6f\x6f\x64\x2e\x63\x6f\x2e\x75\x6b\x2f\x3e\x54\x6f\x6d\x20\x54\x68\x6f\x72\x6f\x67\x6f\x6f\x64\x3c\x2f\x61\x3e\x2e\x3c\x2f\x70\x3e\x0a\x09\x3c\x2f\x66\x6f\x6f\x74\x65\x72\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x7b\x7b\x2d\x20\x2f\x2a\x20\x2d\x2a\x2d\x20\x6d\x6f\x64\x65\x3a\x20\x68\x74\x6d\x6c\x3b\x2d\x2a\x2d\x20\x2a\x2f\x20\x2d\x7d\x7d\x0a"

func viewsBlogTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsBlogTmpl,
		"views/blog.tmpl",
	)
}

func viewsBlogTmpl() (*asset, error) {
	bytes, err := viewsBlogTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/blog.tmpl", size: 2624, mode: os.FileMode(436), modTime: time.Unix(1467467705, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsErrorTmpl = "\x3c\x21\x64\x6f\x63\x74\x79\x70\x65\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x20\x6c\x61\x6e\x67\x3d\x65\x6e\x3e\x0a\x3c\x68\x65\x61\x64\x3e\x0a\x09\x3c\x6d\x65\x74\x61\x20\x63\x68\x61\x72\x73\x65\x74\x3d\x75\x74\x66\x2d\x38\x3e\x0a\x09\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x76\x69\x65\x77\x70\x6f\x72\x74\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x77\x69\x64\x74\x68\x3d\x64\x65\x76\x69\x63\x65\x2d\x77\x69\x64\x74\x68\x2c\x69\x6e\x69\x74\x69\x61\x6c\x2d\x73\x63\x61\x6c\x65\x3d\x31\x22\x3e\x0a\x09\x3c\x74\x69\x74\x6c\x65\x3e\x7b\x7b\x2e\x43\x6f\x64\x65\x7d\x7d\x3a\x20\x7b\x7b\x2e\x4e\x61\x6d\x65\x7d\x7d\x20\xc2\xb7\x20\x6c\x6f\x63\x61\x6c\x2d\x74\x75\x6d\x62\x6c\x72\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x09\x3c\x6c\x69\x6e\x6b\x20\x72\x65\x6c\x3d\x73\x74\x79\x6c\x65\x73\x68\x65\x65\x74\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x61\x73\x73\x65\x74\x5f\x70\x61\x74\x68\x20\x22\x73\x74\x79\x6c\x65\x2e\x63\x73\x73\x22\x7d\x7d\x22\x3e\x0a\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x3c\x62\x6f\x64\x79\x3e\x0a\x09\x3c\x68\x65\x61\x64\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x73\x69\x74\x65\x2d\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x3c\x68\x31\x3e\x3c\x61\x20\x68\x72\x65\x66\x3d\x2f\x3e\x6c\x6f\x63\x61\x6c\x2d\x74\x75\x6d\x62\x6c\x72\x3c\x2f\x61\x3e\x3c\x2f\x68\x31\x3e\x0a\x09\x09\x3c\x68\x32\x3e\x7b\x7b\x2e\x43\x6f\x64\x65\x7d\x7d\x3a\x20\x7b\x7b\x2e\x4e\x61\x6d\x65\x7d\x7d\x3c\x2f\x68\x32\x3e\x0a\x09\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x0a\x09\x3c\x6d\x61\x69\x6e\x3e\x0a\x09\x09\x7b\x7b\x69\x66\x20\x2e\x4d\x65\x73\x73\x61\x67\x65\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x3c\x70\x3e\x7b\x7b\x2e\x4d\x65\x73\x73\x61\x67\x65\x7d\x7d\x3c\x2f\x70\x3e\x0a\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x7b\x7b\x69\x66\x20\x2e\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x20\x2d\x7d\x7d\x0a\x09\x09\x09\x3c\x70\x3e\x7b\x7b\x2e\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x7d\x7d\x3c\x2f\x70\x3e\x0a\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x3c\x2f\x6d\x61\x69\x6e\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x7b\x7b\x2d\x20\x2e\x50\x61\x64\x64\x69\x6e\x67\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x2f\x2a\x20\x2d\x2a\x2d\x20\x6d\x6f\x64\x65\x3a\x20\x68\x74\x6d\x6c\x3b\x2d\x2a\x2d\x20\x2a\x2f\x20\x2d\x7d\x7d\x0a"

func viewsErrorTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsErrorTmpl,
		"views/error.tmpl",
	)
}

func viewsErrorTmpl() (*asset, error) {
	bytes, err := viewsErrorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/error.tmpl", size: 555, mode: os.FileMode(436), modTime: time.Unix(1467435803, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexTmpl = "\x3c\x21\x64\x6f\x63\x74\x79\x70\x65\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x20\x6c\x61\x6e\x67\x3d\x65\x6e\x3e\x0a\x3c\x68\x65\x61\x64\x3e\x0a\x09\x3c\x6d\x65\x74\x61\x20\x63\x68\x61\x72\x73\x65\x74\x3d\x75\x74\x66\x2d\x38\x3e\x0a\x09\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x76\x69\x65\x77\x70\x6f\x72\x74\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x77\x69\x64\x74\x68\x3d\x64\x65\x76\x69\x63\x65\x2d\x77\x69\x64\x74\x68\x2c\x69\x6e\x69\x74\x69\x61\x6c\x2d\x73\x63\x61\x6c\x65\x3d\x31\x22\x3e\x0a\x09\x3c\x74\x69\x74\x6c\x65\x3e\x6c\x6f\x63\x61\x6c\x2d\x74\x75\x6d\x62\x6c\x72\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x09\x3c\x6c\x69\x6e\x6b\x20\x72\x65\x6c\x3d\x73\x74\x79\x6c\x65\x73\x68\x65\x65\x74\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x61\x73\x73\x65\x74\x5f\x70\x61\x74\x68\x20\x22\x73\x74\x79\x6c\x65\x2e\x63\x73\x73\x22\x7d\x7d\x22\x3e\x0a\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x3c\x62\x6f\x64\x79\x3e\x0a\x09\x3c\x68\x65\x61\x64\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x73\x69\x74\x65\x2d\x68\x65\x61\x64\x65\x72\x3e\x0a\x09\x09\x3c\x68\x31\x3e\x6c\x6f\x63\x61\x6c\x2d\x74\x75\x6d\x62\x6c\x72\x3c\x2f\x68\x31\x3e\x0a\x09\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x0a\x09\x3c\x6d\x61\x69\x6e\x3e\x0a\x09\x09\x3c\x66\x6f\x72\x6d\x20\x6d\x65\x74\x68\x6f\x64\x3d\x67\x65\x74\x20\x61\x63\x74\x69\x6f\x6e\x3d\x2f\x67\x6f\x74\x6f\x2f\x3e\x0a\x09\x09\x09\x3c\x6c\x61\x62\x65\x6c\x20\x66\x6f\x72\x3d\x75\x72\x6c\x20\x63\x6c\x61\x73\x73\x3d\x6f\x66\x66\x73\x63\x72\x65\x65\x6e\x3e\x54\x75\x6d\x62\x6c\x72\x20\x55\x52\x4c\x3a\x3c\x2f\x6c\x61\x62\x65\x6c\x3e\x0a\x09\x09\x09\x3c\x69\x6e\x70\x75\x74\x20\x63\x6c\x61\x73\x73\x3d\x75\x72\x6c\x20\x6e\x61\x6d\x65\x3d\x75\x72\x6c\x20\x74\x79\x70\x65\x3d\x75\x72\x6c\x20\x70\x6c\x61\x63\x65\x68\x6f\x6c\x64\x65\x72\x3d\x22\x50\x61\x73\x74\x65\x20\x61\x20\x54\x75\x6d\x62\x6c\x72\x20\x62\x6c\x6f\x67\x20\x6f\x72\x20\x70\x6f\x73\x74\x20\x55\x52\x4c\x20\x68\x65\x72\x65\x2e\x22\x20\x61\x75\x74\x6f\x66\x6f\x63\x75\x73\x3e\x0a\x09\x09\x3c\x2f\x66\x6f\x72\x6d\x3e\x0a\x0a\x09\x09\x3c\x68\x72\x3e\x0a\x0a\x09\x09\x3c\x75\x6c\x3e\x0a\x09\x09\x7b\x7b\x2d\x20\x72\x61\x6e\x67\x65\x20\x2e\x7d\x7d\x0a\x09\x09\x09\x3c\x6c\x69\x3e\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x2f\x62\x2f\x7b\x7b\x2e\x4e\x61\x6d\x65\x7d\x7d\x2f\x22\x3e\x7b\x7b\x2e\x4e\x61\x6d\x65\x7d\x7d\x3c\x2f\x61\x3e\x3a\x20\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x0a\x09\x09\x09\x09\x7b\x7b\x2d\x20\x69\x66\x20\x6e\x65\x20\x2e\x4c\x61\x73\x74\x50\x6f\x73\x74\x20\x30\x7d\x7d\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x2f\x64\x2f\x7b\x7b\x2e\x55\x72\x6c\x2e\x48\x6f\x73\x74\x7d\x7d\x2f\x75\x2f\x7b\x7b\x2e\x4c\x61\x73\x74\x50\x6f\x73\x74\x7d\x7d\x2f\x22\x20\x74\x69\x74\x6c\x65\x3d\x22\x55\x70\x64\x61\x74\x65\x22\x3e\xe2\x86\xba\x3c\x2f\x61\x3e\x0a\x09\x09\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x2e\x55\x72\x6c\x2e\x53\x74\x72\x69\x6e\x67\x7d\x7d\x22\x20\x74\x69\x74\x6c\x65\x3d\x22\x56\x69\x65\x77\x20\x6f\x6e\x20\x54\x75\x6d\x62\x6c\x72\x22\x3e\xe2\xa4\xb4\x3c\x2f\x61\x3e\x3c\x2f\x6c\x69\x3e\x0a\x09\x09\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x09\x09\x3c\x2f\x75\x6c\x3e\x0a\x09\x3c\x2f\x6d\x61\x69\x6e\x3e\x0a\x0a\x09\x3c\x66\x6f\x6f\x74\x65\x72\x20\x63\x6c\x61\x73\x73\x3d\x73\x69\x74\x65\x2d\x66\x6f\x6f\x74\x65\x72\x3e\x0a\x09\x09\x3c\x70\x3e\xc2\xa9\x20\x32\x30\x31\x36\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x68\x74\x74\x70\x73\x3a\x2f\x2f\x74\x6f\x6d\x74\x68\x6f\x72\x6f\x67\x6f\x6f\x64\x2e\x63\x6f\x2e\x75\x6b\x2f\x3e\x54\x6f\x6d\x20\x54\x68\x6f\x72\x6f\x67\x6f\x6f\x64\x3c\x2f\x61\x3e\x2e\x3c\x2f\x70\x3e\x0a\x09\x3c\x2f\x66\x6f\x6f\x74\x65\x72\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x7b\x7b\x2d\x20\x2f\x2a\x20\x2d\x2a\x2d\x20\x6d\x6f\x64\x65\x3a\x20\x68\x74\x6d\x6c\x3b\x2d\x2a\x2d\x20\x2a\x2f\x20\x2d\x7d\x7d\x0a"

func viewsIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexTmpl,
		"views/index.tmpl",
	)
}

func viewsIndexTmpl() (*asset, error) {
	bytes, err := viewsIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.tmpl", size: 941, mode: os.FileMode(436), modTime: time.Unix(1467475123, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/robots.txt": assetsRobotsTxt,
	"assets/style.css":  assetsStyleCss,
	"views/blog.tmpl":   viewsBlogTmpl,
	"views/error.tmpl":  viewsErrorTmpl,
	"views/index.tmpl":  viewsIndexTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"robots.txt": &bintree{assetsRobotsTxt, map[string]*bintree{}},
		"style.css":  &bintree{assetsStyleCss, map[string]*bintree{}},
	}},
	"views": &bintree{nil, map[string]*bintree{
		"blog.tmpl":  &bintree{viewsBlogTmpl, map[string]*bintree{}},
		"error.tmpl": &bintree{viewsErrorTmpl, map[string]*bintree{}},
		"index.tmpl": &bintree{viewsIndexTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
