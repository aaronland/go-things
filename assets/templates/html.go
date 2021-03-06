// Code generated by go-bindata.
// sources:
// www/templates/html/things.html
// www/templates/html/things.html~
// DO NOT EDIT!

package templates

import (
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

var _wwwTemplatesHtmlThingsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\x41\x6f\x9c\x3e\x10\xc5\xcf\xcb\xa7\x98\xff\xdc\x37\xbe\xff\x8b\x91\xa2\xb4\x87\x4a\x95\xda\x43\x2e\x3d\x0e\x78\x36\x38\x31\x36\xb5\x87\x4d\x57\xab\x7c\xf7\xca\x18\x12\x56\x69\xaa\xe5\x82\xc1\xcf\xbf\x19\xbf\x07\x3e\x9f\xc1\xf0\xc1\x7a\x06\x94\xde\xfa\x87\x84\xf0\xf2\x52\xd5\xff\x7d\xfe\x7e\x77\xff\xf3\xc7\x17\xe8\x65\x70\x4d\x55\xe7\x1b\x38\xf2\x0f\x1a\xd9\x63\x53\x01\xd4\x3d\x93\xc9\x83\x7c\xd5\x62\xc5\x71\x73\x3f\x13\xe0\x2b\x7c\xb3\x4f\x0c\xb7\x6d\x98\xa4\x56\x65\x6a\x15\xa6\x2e\xda\x51\x20\xc5\x4e\xa3\x7a\xa4\x23\x95\x17\xaa\x14\xbf\x19\x1d\x75\xdc\x07\x67\x38\xde\x74\xce\xb2\x97\x9b\xc7\x84\x4d\xad\x8a\xec\x5a\x0c\x8d\xf6\x72\x1d\x6c\xae\x2b\x19\xd6\xdb\x0f\x8a\x3b\xeb\x9f\x40\x4e\x23\x6b\x14\xfe\x2d\xaa\x4b\x09\x21\xb2\xd3\x98\xe4\xe4\x38\xf5\xcc\x82\xd0\x47\x3e\x68\xcc\x93\x2b\x71\xd6\xa9\xd9\x3c\xb5\xba\x57\xb7\xc1\x9c\x9a\x6a\x45\x1b\x7b\x84\xce\x51\x4a\x1a\xbb\xe0\x85\xac\xe7\x88\x97\xdd\x57\xbb\x4b\x5d\x0c\xcf\x08\xd6\x68\xfc\x35\x71\x3c\x61\x93\xe7\x67\xd6\x21\xc4\xa1\xa9\x76\x59\x6e\xfd\x38\xc9\xa6\x65\x04\x4f\x03\x6b\x9c\xed\x2e\xab\x97\xe1\x91\xdc\xc4\x1a\x11\x36\x49\x68\xec\xfa\x10\x12\x03\xc1\xa2\x52\x6f\x55\xd4\x52\x26\x0f\x8d\x3d\x36\x1f\xf7\x17\x39\x4d\x4e\x12\xbe\x13\xb7\x11\x3a\xc7\x14\x35\x92\x73\x2b\xfc\x43\xcc\x6c\xe6\x3e\x57\x45\x98\xfd\xd6\x68\x6c\x1a\x1d\x9d\xfe\xf7\xc1\xf3\x27\x2c\xd4\x12\xf1\x48\x7e\xb3\xc8\x51\xcb\x6e\x0e\x74\x24\xff\xba\x85\xbf\x19\x96\x4d\xa2\xc8\xb4\xf8\x54\x96\xe7\xa8\xb6\x3d\x94\x67\x43\x42\xfb\xe7\x3e\xa4\xe0\x0f\x36\x26\xd9\x67\x41\x2e\xb2\x32\x16\xe4\xbc\xdf\xdd\x5a\xab\x9d\x44\xc2\xb6\xb7\x34\xb5\x83\x15\x5c\x42\x5a\x9f\x96\xcd\xb7\xe2\xa1\x15\xbf\x4f\xc3\x7c\x1b\xa3\x1d\x28\x67\x7d\x6b\x4c\xad\x0a\xea\x5f\xec\x8e\x7c\xc7\xee\x2a\x36\x36\x77\xb3\xf8\x12\xbb\xa4\xf5\x3e\xf3\xd5\xbf\x55\xb0\x7b\xfb\xbd\xca\x9b\xd7\x5f\xae\x56\xe5\x4b\xaf\x55\x39\x52\xce\x67\x60\x6f\xf2\x61\xf3\x27\x00\x00\xff\xff\x02\xe0\x79\x44\x83\x04\x00\x00")

func wwwTemplatesHtmlThingsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlThingsHtml,
		"www/templates/html/things.html",
	)
}

func wwwTemplatesHtmlThingsHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlThingsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/things.html", size: 1155, mode: os.FileMode(420), modTime: time.Unix(1580227464, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlThingsHtml2 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xbd\x6e\xc3\x30\x0c\x84\xe7\xfa\x29\x58\x2d\xdd\xac\x17\x60\xb4\xf4\x67\x4d\x87\x02\x45\xa7\x40\x95\x98\x5a\x80\x44\x15\x16\x11\xb4\x10\xfc\xee\x05\xe3\x18\xc8\x14\xf4\x16\x13\x87\xf3\x7d\x14\x7b\x87\x48\xc7\xc4\x04\x26\x71\xa4\x1f\x03\xcb\x32\xe0\xfd\xd3\xfe\xf1\xed\xe3\xf5\x19\x26\x29\xd9\x0d\xa8\x1f\xc8\x9e\xbf\x76\x86\xd8\xb8\x01\x00\x27\xf2\x51\x07\x15\x4a\x92\x4c\xee\x7d\xaa\x0f\x0d\xf6\x0c\x2f\x69\x6e\x82\x76\x75\x2f\x99\xde\x85\xca\x77\xf6\x72\x26\x85\x43\xa8\xa5\x54\x3e\x14\x12\x6f\x96\xe5\x76\x28\xb4\xa6\x19\xc5\xda\x8d\x8b\x9f\x35\xfe\xba\x61\xdb\x20\xa6\x13\x84\xec\x5b\xdb\x99\x50\x59\x7c\x62\x9a\x8d\x83\x2b\xdd\x46\x68\x2d\xcd\x06\x46\xe5\xdc\x01\x60\xf1\x89\x9d\x4e\xaa\x71\x1c\xcf\xa6\x5d\xdd\xff\x14\x1e\x6b\x95\xad\xf0\xb2\xa3\x8d\xe9\x74\xf5\x33\xda\xf5\x09\x68\xd7\x2b\xf7\x0e\xc4\x51\xef\xff\x17\x00\x00\xff\xff\xe0\x53\x6d\xb7\x95\x01\x00\x00")

func wwwTemplatesHtmlThingsHtml2Bytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlThingsHtml2,
		"www/templates/html/things.html~",
	)
}

func wwwTemplatesHtmlThingsHtml2() (*asset, error) {
	bytes, err := wwwTemplatesHtmlThingsHtml2Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/things.html~", size: 405, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
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
	"www/templates/html/things.html": wwwTemplatesHtmlThingsHtml,
	"www/templates/html/things.html~": wwwTemplatesHtmlThingsHtml2,
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
	"www": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"html": &bintree{nil, map[string]*bintree{
				"things.html": &bintree{wwwTemplatesHtmlThingsHtml, map[string]*bintree{}},
				"things.html~": &bintree{wwwTemplatesHtmlThingsHtml2, map[string]*bintree{}},
			}},
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

