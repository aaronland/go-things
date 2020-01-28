// Code generated by go-bindata.
// sources:
// www/static/javascript/things.api.js
// www/static/javascript/things.init.js
// www/static/javascript/things.placeholder.client.js
// www/static/css/things.css
// DO NOT EDIT!

package http

import (
	"github.com/whosonfirst/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

var _staticJavascriptThingsApiJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdb\x6a\xdc\x30\x10\x7d\xb6\xbe\x62\xea\x17\xcb\xc4\x38\x21\xf4\x25\xd9\xba\x7d\x0a\xbd\x90\x5e\xc8\x05\x0a\xa5\x2c\x8a\x35\xbb\x11\x38\x92\x33\x92\x73\x61\xe3\x7f\x2f\x92\x6c\xef\x3a\x49\x4b\xdf\x64\xe9\xcc\xcc\x99\x33\x67\x7c\x27\x08\xdc\xb5\xd2\x6b\x0b\xd5\x78\x78\x7a\x82\x4d\xbf\x60\x2c\x7e\x96\xa2\x55\x50\x01\x5f\x75\xba\x76\xca\x68\x9e\x6f\x18\x03\x00\xf0\xa1\xa8\x65\x6b\x94\x76\x50\x41\x63\x6a\xe1\xdf\x4b\x43\x6a\xad\xf4\x22\x60\x26\xa0\xc5\x66\x05\x15\x6c\x18\x4b\x32\x21\x65\x76\x0c\x53\xbe\x50\xa6\x80\x9a\xba\x9b\xab\x02\x8c\x5e\xda\xae\xae\xd1\xda\x70\x46\x22\x43\xbe\x62\x32\x65\x72\xb4\x0c\x21\x50\xc1\x97\xf3\xef\xdf\x4a\xeb\x48\xe9\xb5\x5a\x3d\xc6\x4c\xf9\x22\x62\xb7\x11\x46\x2f\x1b\x23\x24\x54\xdb\x9a\x64\xdb\x90\x34\x09\xfd\x0b\x5a\xa3\x6f\x81\x6c\x5b\xc6\x8f\x05\x4b\x12\x96\x24\x6a\x05\x3c\x5e\x94\x84\x42\x3e\x9e\x3b\xe1\x10\xde\x54\xf0\x36\xdf\xb0\x24\x14\x20\x74\x1d\x69\x8f\xef\x63\x4c\xe4\x28\x5c\x67\x97\xb5\x91\xe8\x65\x0d\x19\x7e\x65\xf1\x36\xfb\xbd\x98\xa3\x1c\x3e\xb8\x17\xa8\x0b\x7c\x70\x11\x39\xd0\xe0\xbb\x39\xdf\xc1\xe1\xc1\x41\xee\x07\x35\xbb\x7e\x0f\x87\x47\x47\xf9\x44\x6d\x94\x8f\x6f\x32\xff\x9c\x1d\xef\xf2\x2a\x20\xbb\x41\x6b\xc5\x7a\xe7\xde\x33\xe9\xbd\x7e\xaf\x77\xe6\xe8\x11\xc6\xdc\x9e\x3f\x89\xfb\x1d\xde\x84\xb6\x35\xda\xe2\xc0\x3c\x19\x86\x30\xa1\xa5\x70\x62\x1c\x59\x2b\xc8\x22\x27\x71\x3f\x15\xdb\x8e\x9d\x7b\x60\xb8\xef\xa1\x16\xae\xbe\x06\x8e\x39\xbc\xe8\x09\xf3\x2d\x33\xff\xd0\xbf\x36\xf6\x95\x50\x0d\xce\x06\x8f\x41\x9d\x79\x96\xbf\x86\x8b\x2b\x43\x6e\x37\xfa\xbf\x83\x3b\xf2\x3b\x33\x6d\xc7\x1e\xa4\xfb\x42\xca\xfd\x0f\xc1\xe4\x55\x0a\x7b\xd1\xee\x43\x5c\x6d\xb4\x35\x0d\x96\x8d\x59\xf3\xf4\xf2\xec\x73\x5a\x74\xa4\xe6\x3e\x8e\xda\xb3\x24\x99\x61\x3f\x9e\x5c\x4c\xd8\x60\x29\xc2\x5b\xa8\x40\xe3\x3d\xfc\xfc\x7a\xfa\xc9\xb9\xf6\x0c\x6f\x3b\xb4\x8e\xe7\x83\x95\x08\x6f\x4b\x21\xe5\xc9\x1d\x6a\x77\xaa\xac\x43\x8d\xc4\x53\xbf\x1e\x69\x31\x2e\x4a\xc0\xbe\x0e\x0c\x8d\x47\x64\xd4\xf6\x1f\xd8\xa0\x5e\xc4\x86\xe3\x2e\x05\xd3\xa2\xe6\xe9\x8f\xcb\x8b\xb4\xf0\x5a\x15\xe0\xa8\xc3\x29\x97\x45\x2d\xf9\xb4\xe7\x43\x5c\x10\xfa\x99\x23\x9e\xb9\x21\x3a\x16\x56\xa2\xb1\x38\x8e\x66\x14\xb1\x2f\xd8\x30\x2a\xb6\x35\x77\xf8\x2d\x0d\x7f\xaa\x3e\xf7\x22\xfd\x09\x00\x00\xff\xff\x2e\xe6\x71\xb6\x10\x05\x00\x00")

func staticJavascriptThingsApiJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptThingsApiJs,
		"static/javascript/things.api.js",
	)
}

func staticJavascriptThingsApiJs() (*asset, error) {
	bytes, err := staticJavascriptThingsApiJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/things.api.js", size: 1296, mode: os.FileMode(420), modTime: time.Unix(1580229808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptThingsInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\xcd\x6e\xe3\x36\x10\x3e\xcb\x4f\x41\xf0\x12\x09\x51\x84\x05\x7a\x5b\x6f\x0e\x69\x9b\xa2\x41\xb7\x5d\x20\x0d\x50\x14\xc1\x22\xa0\xa9\xb1\xcd\x86\x26\x55\x92\x4a\x6a\x2c\xf2\xee\xc5\x0c\xf5\x67\xc9\x56\xbc\x97\x5d\x8b\x9c\x6f\x38\xf3\xcd\x6f\x5e\x95\x29\xed\x6b\x21\xca\xf2\xf6\x05\x4c\xf8\xac\x7c\x00\x03\x2e\xe5\xda\x8a\x92\xe7\x6c\x5d\x1b\x19\x94\x35\x0c\xbf\x53\x40\x99\xec\xdb\x62\xc1\x18\x63\x2f\xc2\xb1\x4a\xb3\x6b\x56\x5a\x59\xef\xc0\x84\x62\x03\xe1\x56\x03\xfe\xfc\x71\x7f\x57\xa6\xbc\xd2\x42\x02\xcf\x96\xbd\xfc\xbf\x35\xb8\xfd\x1c\x84\x04\x78\xb6\x44\x40\x87\x72\xe0\x6b\x1d\xfc\x1c\xae\x11\xc1\xc7\x5a\x54\xd8\x2a\xb3\x79\x5a\x5b\xb7\x9b\x03\x92\xd4\x15\x4a\x4d\xb1\x5a\xac\x60\xd6\xc1\x08\x26\xb1\x29\x7a\x65\xcb\x59\x57\x23\x18\xa5\xa6\x58\x5f\xaf\x76\x2a\xbc\x8f\x8e\x72\x0d\x5f\x07\x9c\x45\x3d\x52\x18\x79\x8e\x0b\x51\xee\x20\x54\x0e\x3c\x84\x96\xbe\x36\x0d\x52\x0c\x7e\xd2\xfb\x57\x38\xd8\xd9\x17\xb8\x09\xc1\xa9\x55\x1d\x20\xe5\xa5\x08\xe2\xea\x75\x6b\xbd\x35\x6b\xe5\x7c\xb8\x52\x25\xaa\x1d\x62\x5e\x84\xae\x81\x5d\x33\xce\xbb\x73\xa2\xb0\x50\xc6\x80\x7b\x80\xff\x42\x73\xd9\xde\xa2\x11\x85\x0f\x7b\x0d\x45\xa9\x7c\xa5\x05\xf2\xca\x8d\x35\x80\x1a\x16\x49\x13\xfc\x88\xff\xf5\xe1\xf7\xcf\xad\xf2\xf6\x62\x06\x4b\x09\x37\x15\x58\x69\x2b\x9f\x51\xa2\xd2\x43\x7b\x13\xa2\xe7\x2d\x86\x8b\xfe\x19\xf2\x5c\x58\x23\xb5\x92\xcf\x23\xc2\x92\x9e\xca\x34\x23\xab\x42\xed\x0c\x5b\x0b\xed\x21\x6a\x7a\x1b\xeb\x8b\x71\x3d\xa1\x8f\x64\x13\x8c\x51\x09\x95\x92\xc1\x3f\xa9\x92\x5d\x0f\xb2\x0e\x83\x7c\x46\x48\x0e\xd0\x95\x70\x1e\xee\x4c\x48\xfb\xd3\x8c\x08\xa2\x64\x8a\x31\x19\xc7\x10\x23\x84\xd7\xd2\xd5\xbb\xd5\x30\xc9\x4e\x19\x41\x82\x3c\xea\x0d\x6e\xcf\xbe\x2d\x92\xe8\x4c\x9b\x74\xd6\x3c\xf9\x5a\x4a\xf0\x7e\xe8\xb3\xf3\x15\xd2\x98\x48\x6b\xbc\xd5\x50\x68\xbb\x49\xf9\x97\xdf\x6e\xfe\xe6\x39\xc3\xbb\x65\x92\x2c\x92\x31\xcb\x4d\xa0\xa6\x0f\x80\x73\xd6\x0d\xd5\x83\x73\x53\xf5\xb7\xf7\xf7\x5f\xee\x79\xce\xf0\xf2\x88\x7e\x3c\x69\x5e\xe8\x95\x13\x3f\xec\x1a\xdd\x4a\x2e\x7a\x1e\x2f\x3e\x0e\x02\x95\xe3\x9d\xb1\x01\x2e\x3e\x12\xad\xf9\x48\x0f\xe9\xf0\x85\xa8\x14\x36\xe5\x94\x3e\xf3\xc8\x70\x3e\xa0\x27\xef\x3c\xc9\x3a\x1f\xdf\x98\x14\x41\x6e\x59\x0a\x59\x43\xec\x81\x47\x7f\xde\xfc\x8c\xfe\xb4\xf2\x23\xba\xde\x16\x47\x13\x73\x90\xe9\x94\x70\x4e\xbc\x3e\xb5\x6e\x76\x04\x1a\xb1\x83\x9c\xa9\x92\x9a\xc3\x89\x82\x6a\x2b\xee\xdd\x92\x9c\x2b\xf9\xb6\x2a\x67\xda\xc6\x03\xf1\xc7\xee\x98\x56\xcf\xc0\xc4\xca\xd6\x81\x71\x76\xc9\xd0\xc6\x98\x77\x7d\x12\xfb\x77\xcb\x84\xbc\x3a\x52\xf5\x1d\x17\xfd\x68\xea\xb3\x35\x1e\x11\x19\x54\x1d\xb6\x36\x68\x5a\xeb\xb9\x06\xb3\x09\xdb\xae\xb6\xb4\xf2\x07\x8d\x5e\x3a\x10\x01\x9a\x1e\x9d\xf2\x9a\xba\x72\x82\x52\x23\x7b\xa5\x16\xde\xf3\x9c\x71\x0f\x1a\x64\xb8\xda\x81\xa9\x9b\xda\x5a\x5b\xc7\x52\x54\xae\xae\x3f\x2c\x99\x62\x9f\xa2\x11\x4b\xa6\x2e\x2f\xc9\xae\xae\xc9\xf7\x76\x3d\xaa\xaf\x93\x6a\xa1\xce\xe0\x0a\x55\x2e\xfb\x33\x24\x92\x4e\x1b\x46\xfb\x55\x40\x48\x08\xfb\x2a\x5e\x76\x5f\x13\x9d\xed\x50\x7d\xfc\x3a\x2c\x1d\xad\x0c\x88\x4d\xc4\x36\xbf\x1b\x64\x7f\xd3\xfc\x7a\xfc\xd0\x21\xd5\x9a\xa5\xcd\x69\x61\x40\x6d\xb6\x2b\x5b\xbb\xad\xb5\x25\xd5\x73\xcc\x8e\xaa\xf6\xdb\x94\xb3\x14\x93\xe0\xa8\x2c\x39\xc2\x2e\x19\xcf\x78\xd7\x37\xe2\x7f\x03\xcb\x87\x2f\xed\x84\x74\xf6\x8c\x57\x3a\xb9\xd3\x2f\x4c\x95\x6b\x2b\x85\x56\x61\x3f\xd6\x3d\xbe\x27\x9d\xbd\xb6\xa9\x22\x8a\xf8\x44\xcd\xc8\xc4\x28\xf4\x3d\xf6\x39\xd8\x28\x6b\x4e\x59\x17\x6f\x1b\xdb\xba\x1e\x79\xca\x3a\x77\xd2\xcb\xe6\xfa\xa8\x93\x98\x2e\x3e\xb8\x6e\x3b\x8b\xf0\x7f\xac\x32\x29\x56\x43\x36\x4c\x2b\x15\x60\x37\x53\x5c\x5a\x75\x0e\xa3\xe4\xd9\x0d\xe1\x7c\x08\x3a\xc0\xf3\xde\xde\x66\xb3\xed\xf1\xa2\xaa\xc0\x94\x3f\x6d\x95\x2e\xd3\x91\x99\xd8\xd2\xfe\xb0\x25\xa4\x3d\xba\x73\x8e\xb0\x47\xd6\x03\xa0\xf2\xa6\xc6\x42\xe4\x40\x11\x84\xdb\x40\x58\x36\x87\x54\xd1\xa0\xcf\xdc\x10\x92\x41\xc1\xbf\x0f\x22\x5f\x09\xd6\x8f\x89\x7e\x36\x2c\x69\x84\x0e\x07\xcc\x91\x11\x4d\x6d\x6e\x48\x09\xfa\xd9\x8e\xa7\xef\xff\x53\xe0\xe4\x66\xd8\xdf\x0c\x1f\xc3\xd7\xb3\xb9\x19\xd5\xce\x9f\xf1\x38\xa8\x70\xf5\x53\xa6\xaa\xc3\x61\x2c\x58\x3b\x03\xc6\xb1\x18\x2c\x55\xa0\xfb\x65\x0a\x2b\x04\x8f\x9b\x19\xc1\x3e\xb1\x1f\xda\x51\x1e\xa9\x23\x26\x9a\xc1\x31\xbb\x2b\x21\x64\x38\xa0\xe8\x1c\xd1\xed\xce\x36\xb3\x08\x9d\xda\x1c\x9c\xeb\x34\x0c\x96\xb7\x66\x61\xa1\x86\xbf\xb5\xba\x04\x57\x48\xad\x30\x30\x1e\x84\x93\x5b\x72\xe8\xf4\xea\xf2\xce\xd2\xf2\xd7\xc3\x2f\x3c\x87\x98\x01\xed\x9e\xf4\x96\x2d\x17\xff\x07\x00\x00\xff\xff\xb6\x87\x3b\xfd\xb8\x0e\x00\x00")

func staticJavascriptThingsInitJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptThingsInitJs,
		"static/javascript/things.init.js",
	)
}

func staticJavascriptThingsInitJs() (*asset, error) {
	bytes, err := staticJavascriptThingsInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/things.init.js", size: 3768, mode: os.FileMode(420), modTime: time.Unix(1580229879, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptThingsPlaceholderClientJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xc1\x6e\xdc\x36\x10\x3d\x93\x5f\x31\xd5\x45\x12\xbc\x90\x9c\xa0\x87\x64\x15\xb5\x87\xc2\x6d\x50\xa4\x2d\xd0\xf8\x50\xc0\x08\x16\xb4\x34\xbb\x22\xaa\x90\xf2\x70\x14\xc7\xd8\xec\xbf\x17\xa4\xa4\x95\x64\x6f\x8b\x9e\x56\x3b\xf3\xde\xf0\xf1\xcd\x70\xbe\x28\x02\x6e\xb4\x39\x38\x28\xa7\x8f\x6f\xdf\xe0\x78\x2a\xe4\xf0\x2f\xeb\x5a\x55\x61\x63\xdb\x1a\xe9\x0c\x59\x05\x47\xf8\x05\x7c\x56\xb5\x1a\x0d\x43\x09\xc9\xbe\x37\x15\x6b\x6b\x92\xf4\x28\x25\x00\x80\x3f\x18\x4d\xdd\x59\x1d\x00\x71\xc3\xdc\x6d\xf3\xbc\xb5\x95\x6a\x1b\xeb\x78\xfb\xe6\xd5\x9b\x57\xb9\xea\x74\x5c\xcc\x04\x87\xed\x1e\x4a\x38\x4a\x21\x45\xac\x8d\xe6\x78\x0b\xab\xca\xe2\xb4\x09\x39\x87\x8a\xaa\x66\x99\x65\xfc\xca\x1b\xb0\x66\xe7\xfa\xaa\x42\xe7\xc2\x37\x12\x59\x0a\xbc\xe9\x04\x6b\x76\xad\x55\x35\x94\x33\x93\x5c\x17\x20\x22\x78\xa5\xe8\x80\x5e\x30\xb9\x2e\x1b\xfe\x14\x52\x08\x29\x84\xde\x43\x32\x04\x32\x42\x55\x3f\x7d\x64\xc5\x08\xdf\x95\xf0\x7d\x7a\x94\x22\x1c\x40\xc8\x3d\x19\x8f\x3f\x0d\x9c\x70\x27\x56\xdc\xbb\x5d\x65\x6b\xf4\xfe\x86\x0a\x77\xf1\x10\x8d\x3f\x15\x6b\x94\xbf\xc5\x0b\xd4\x2d\x7e\xe5\x01\x39\xca\x48\x96\x35\xdf\xc1\xeb\xeb\xeb\xd4\x77\x69\x15\xfe\x01\x5e\xbf\x7d\x9b\x9e\xa5\x4d\x66\x24\xc7\xd8\xa7\xe3\xed\x52\xd7\x06\xe2\xcf\xe8\x9c\x3a\x2c\xe2\x5e\xc9\x29\x2d\x2e\xdc\x4c\x0a\xc1\xf4\x04\x53\xe5\xf1\xc7\x5f\x82\xd4\xe3\x42\x3c\xa1\xeb\xac\x71\x38\xcb\x9f\x70\xb5\x62\x05\x25\xfc\xfa\xf1\x8f\xdf\xb3\x4e\x91\xc3\x84\xd4\x63\x5a\xc8\x11\x32\xa6\xfd\x30\x64\xbb\xbd\x6e\x19\x69\x47\xe8\xfa\x96\x5d\xe2\x73\x67\x55\x73\xb7\xd7\x71\x2f\x13\x2a\xc5\x55\x03\x09\xa6\xf0\xc2\x03\x4c\xe7\x1e\xf9\xc4\xa9\x90\x23\x6f\x31\x26\x7b\xa5\x5b\x5c\x0d\x0a\x06\x37\xd7\x55\xfe\x95\xae\xee\x2d\xf1\x92\xfd\xbf\xc9\x3d\x69\x28\xe7\xb7\x73\x05\x51\x1e\x4c\xa2\x7c\x98\xfa\xfc\x47\x46\xfa\x5c\x46\x70\x05\xbe\x49\x23\xbf\xb2\xc6\xd9\x16\xb3\xd6\x1e\x92\xe8\xe7\x9b\xdb\x9f\xde\x47\x1b\x5f\x2a\x5d\xd5\x1f\x1a\x27\x85\x58\xc1\x7f\xb9\xb9\x8d\x36\x23\x36\x4c\x23\xe1\x03\x94\x60\xf0\x11\xfe\xfa\xed\xc3\x7b\xe6\xee\x4f\x7c\xe8\xd1\x71\x92\x8e\x53\x48\xf8\x90\xa9\xba\xbe\xf9\x82\x86\x3f\x68\xc7\x68\x90\x92\xc8\xbf\xac\x68\x33\xbd\xb1\x80\xbd\x0c\x0c\x1e\x0c\xc8\xc1\xe6\xff\xc0\x06\x23\x07\x6c\xf8\x5c\x4a\xb0\x1d\x9a\x51\xbe\xbf\xeb\x06\x98\x7a\x3c\xd7\x72\x68\xea\x49\x71\xb0\xfa\xd9\x4c\x3c\x9b\x87\x61\xc6\x61\xaf\x5a\x87\x53\x73\x26\xef\x4e\x1b\x29\x45\xfc\x6c\x18\x97\xeb\x67\x0c\xad\x76\xcd\x80\x0e\x33\x74\xf7\xa9\x58\x24\x2a\xdb\x87\xa5\x38\x92\xb2\x16\xcd\x81\x9b\x09\xb1\xb7\x04\x89\x87\xe9\xf2\xba\xd0\xf0\x6e\x80\x17\xa0\xaf\xae\xe6\x3d\x45\x33\xfd\x4e\x4f\xbb\xc1\x3d\xea\x70\x3f\x1a\x96\x34\x3f\x75\x8b\xe9\xaf\x94\x43\x88\xc2\xfe\xd5\xfc\x14\x6d\xa5\x10\x62\x52\x98\x75\xbd\x6b\x12\x0a\x36\x88\x7b\x42\xf5\xf7\xf4\x94\xf2\x7c\x24\x1a\xd4\x87\xe6\xde\xf6\xd4\x58\x5b\x07\xb6\x4f\x8b\x3c\x87\x8b\x45\xf2\x1c\x56\x75\x6a\xdc\xab\xbe\xe5\x70\x6a\x65\x0d\x6b\xd3\xe3\xf8\x08\x07\xa3\xe5\x62\xcf\x9c\x4b\x16\xd2\x03\xe4\xf8\x4e\xe4\x02\xe1\x97\x43\x11\x02\xf2\x94\xfa\x26\xff\x13\x00\x00\xff\xff\x52\x17\x25\x6e\xe9\x06\x00\x00")

func staticJavascriptThingsPlaceholderClientJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptThingsPlaceholderClientJs,
		"static/javascript/things.placeholder.client.js",
	)
}

func staticJavascriptThingsPlaceholderClientJs() (*asset, error) {
	bytes, err := staticJavascriptThingsPlaceholderClientJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/things.placeholder.client.js", size: 1769, mode: os.FileMode(420), modTime: time.Unix(1580229580, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticCssThingsCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xd1\x6a\x84\x30\x10\x45\x9f\x93\xaf\x08\xf8\x9c\x65\x97\x76\xa1\xc4\xaf\xd1\x64\x34\xa1\x63\xc6\xc6\x71\x55\xca\xfe\x7b\x89\x59\x76\x0b\xd2\x3e\x3a\x97\x99\x7b\x3c\xa9\xbe\x66\x48\x9b\xfa\x96\xa2\x43\x6a\xd8\x20\x74\x5c\xcb\xbb\x94\xa7\x09\x10\x2c\xeb\x01\xe2\x9c\xe3\xb1\x71\x2e\xc4\xde\xa8\xf3\xb8\xd6\x52\x0c\x4d\xea\x43\x7c\x7d\xad\xda\x43\xe8\x3d\x9b\xeb\xf9\xe6\x1f\x93\x25\x38\xf6\x79\xb0\xd4\x52\xd0\x0d\x52\x87\xb4\x98\xc9\x26\x42\xac\x9f\x17\x75\xae\x34\x97\xf7\xfd\x50\x4b\xc9\x41\xd2\x4c\xa3\x51\x97\x71\x55\x13\x61\x70\xaa\xb2\xd6\xbe\xc2\x96\x98\x69\x38\xe4\x42\x1c\xb0\x31\x64\xf2\xff\xd7\x9e\xf9\x8e\xf1\x67\x9a\xf6\x9f\x3b\x74\xfe\xf2\x72\xba\x26\x18\x6a\x29\x30\x4c\xac\x27\xde\x10\x34\x6f\x23\x18\x15\x29\xc2\x51\x29\x06\xe3\xb3\x93\x9d\xb0\xb1\x9f\x7d\xa2\x39\x3a\x6d\x09\x29\x99\xd2\x7d\x97\xb2\x62\x9f\x15\xb5\xe4\xf6\x37\x2a\x46\x3f\x8a\xd1\x87\xf0\xb7\x22\xbc\x70\x9a\x02\x97\xb7\xf2\xfe\x4f\x00\x00\x00\xff\xff\xf8\x57\xc2\xb8\xde\x01\x00\x00")

func staticCssThingsCssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssThingsCss,
		"static/css/things.css",
	)
}

func staticCssThingsCss() (*asset, error) {
	bytes, err := staticCssThingsCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/things.css", size: 478, mode: os.FileMode(420), modTime: time.Unix(1580227772, 0)}
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
	"static/javascript/things.api.js": staticJavascriptThingsApiJs,
	"static/javascript/things.init.js": staticJavascriptThingsInitJs,
	"static/javascript/things.placeholder.client.js": staticJavascriptThingsPlaceholderClientJs,
	"static/css/things.css": staticCssThingsCss,
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
	"static": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"things.css": &bintree{staticCssThingsCss, map[string]*bintree{}},
		}},
		"javascript": &bintree{nil, map[string]*bintree{
			"things.api.js": &bintree{staticJavascriptThingsApiJs, map[string]*bintree{}},
			"things.init.js": &bintree{staticJavascriptThingsInitJs, map[string]*bintree{}},
			"things.placeholder.client.js": &bintree{staticJavascriptThingsPlaceholderClientJs, map[string]*bintree{}},
		}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
