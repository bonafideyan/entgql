// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/node_sdl.tmpl
package internal

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateNode_sdlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\xcd\x6e\xe3\x36\x10\xbe\xfb\x29\x46\x82\x0f\xdd\x45\xd6\x0b\xe4\x28\xa0\x07\xaf\x9d\x00\x06\x6a\xab\x4d\xbc\xe8\xa1\x08\x1a\xae\x35\x92\x09\x50\xa4\x4a\x52\xcd\xba\x8a\xde\xbd\x18\x52\x12\x65\x25\xd9\x04\x8b\x9c\x4c\x6b\xa8\xf9\x7e\xe6\xc3\xa8\x69\x3e\xc1\x5c\xaa\x0c\x21\xf9\x15\x16\xf0\xa9\x6d\x67\xfe\x11\x2b\xfd\xa3\x1d\x1d\xe8\xf1\xac\x69\x3e\x7f\x84\x2b\x59\x97\x06\x3e\x7e\x76\x8f\x50\xd6\x25\x34\x8d\xbb\xdc\xb6\xa9\xce\x50\x7f\x39\x41\x33\x03\xd8\x2d\xf7\x5f\x6f\x96\xbf\xcd\x00\x36\xeb\xbf\x97\xb7\x2b\x7f\x58\x5f\xdd\xae\x5c\x7b\xcd\x64\x81\x30\xe7\x84\xe0\xd0\x17\x1b\x99\xe1\x77\x34\x6d\x3b\x03\xea\xc8\x3d\xee\x23\xd4\x55\x85\x7a\xc5\x0c\xb6\x6d\xd7\xe7\xa5\xea\xd0\x1c\x65\xd6\xb6\xb3\x9e\xf0\xfe\x54\xe1\x40\xd8\x9e\x2a\x0c\x84\x81\x97\x95\xc0\x12\xa5\x35\xb0\x23\x0b\x88\x39\xcf\x12\xd8\xac\xa3\x31\xcd\x3c\xd0\xbc\xe6\x28\xb2\x81\x65\xde\xf3\x38\xb0\x12\x05\xff\x8f\x8e\x42\x3d\x78\x4a\xd7\x5c\x1b\xdb\xb6\x89\xbf\x48\x34\xe0\x11\xcc\x81\x09\xa6\x77\x0e\xbe\x69\x78\x0e\x52\x59\xa0\x3e\x5c\x08\xf6\x4d\x60\xdb\x46\x4d\xe3\x05\x04\x29\x23\x2a\x18\xa8\x5c\x65\x85\xf3\xcb\x8d\x0b\x77\xdd\xbc\xe6\xe8\xa0\x16\x1e\xc2\x15\x79\x0e\x4a\x53\x61\x7b\x99\xd2\x4f\x7a\x99\x42\xaf\x00\xdf\xaa\x00\xa7\x9c\x71\x91\x56\x96\x2b\xc9\xc4\x94\xb3\x30\x18\x30\xd3\xcb\xad\x87\xde\x3a\xc8\x38\x8e\x6f\x90\x65\x06\x98\xcc\x00\x25\x49\x36\x50\xb1\x82\x4b\x46\xcd\xc0\x1e\xb5\xaa\x8b\x23\x30\x30\x68\x41\xe5\x70\x1f\xb0\xef\x17\x71\x1c\xbf\x9d\xf6\x2f\x2c\xb7\xa8\x13\xb8\xb5\x9a\xcb\xe2\x02\xbe\x61\xae\x34\x86\xff\x39\x5d\x4b\x60\x23\xed\x05\x08\x36\x1c\x1f\x8e\x48\xb7\x02\xec\x9f\xf4\x60\x23\xab\xda\x5e\x80\xf2\x11\x4f\xe0\xaf\x50\xef\x62\x1f\xdd\x7d\x18\xbf\xb5\x52\x52\xe2\x81\x34\x4d\x06\x39\xa4\x33\x8e\xe3\xa5\x93\xe7\xc3\x78\x0f\x98\x15\x08\x9c\x3c\x40\x38\x0c\xaf\x3b\xd1\xe7\xb9\xa5\xc9\xbb\xac\xba\x16\x87\x5a\x1b\xa5\x21\x57\x1a\x6a\xe3\x1a\x04\x3f\x3b\xc7\xfc\x95\x5e\x7a\xe4\xdf\xdc\x1f\xf1\x0c\x9e\x59\x87\x8c\x32\x23\xdf\xdd\x31\x2b\xb0\xeb\x40\x89\x4b\x02\x83\x68\x10\x10\x88\x82\x55\xc0\x40\x70\x33\xcc\xad\xef\xfc\x2f\x13\x35\x9a\x73\x21\x94\xe0\xa9\x53\x41\x53\xdf\x85\x18\x18\x78\x38\xf2\xc3\x91\x90\x2c\xe3\xd2\x38\x6a\x67\xc4\x65\xd6\x9b\x40\x14\x78\xf6\xac\x07\xae\x95\x1f\x5c\x70\x31\xba\xeb\xcc\xd8\xc8\x5c\xe9\x92\x0d\x3a\x5e\x68\x52\xb1\x02\xe9\x6a\x02\xbf\x77\xa7\x91\x99\x07\x55\x4b\xc7\x9a\x09\x71\x46\xf0\xa4\x6a\x2a\x8a\x0c\x0a\xb4\x90\x6b\x55\x3e\x37\x63\x00\xab\x2c\x13\x2b\xea\xe2\xb2\x18\x0d\x3b\xcc\xa5\x2f\x6c\x31\x4e\x7f\xc3\x30\x42\x40\x87\x05\xe6\x07\x7d\xcd\x85\x45\xfd\xda\x26\xeb\x56\xc4\x3c\x5f\x7c\x95\xfc\x9f\x1a\xdf\x67\xb9\x8d\xb0\xa7\xd9\x07\x58\xee\xd6\xe3\x49\x04\x01\xd1\xdd\x0c\x20\xbd\xf9\x41\x71\x97\xee\x5f\xac\xf6\x76\xfd\x51\xa3\x3e\x0d\xdf\xa8\xef\x96\x22\xed\x82\xe7\x0b\x8d\xd7\x27\xbd\xba\x27\x6b\xa3\xfb\x00\x7c\x18\xc5\xfd\x67\x2d\x94\xaf\x1a\xf8\xe5\x34\x76\x7a\xf8\x9c\xf5\x5c\xc6\xc5\xb7\x7b\xff\x94\xfb\xd3\x11\xfc\xd4\x22\x96\xd3\x3d\xfc\x9a\x40\x78\x84\x4a\xd4\x9a\x51\xf1\x1d\x36\xf2\x74\xe6\x93\x85\x2c\x9f\xdb\xc7\x72\xba\x8e\xfb\x94\x6c\x6b\xeb\xa5\x76\x41\xf9\x3f\x00\x00\xff\xff\x3c\xbb\x7f\x2d\x0a\x09\x00\x00")

func templateNode_sdlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateNode_sdlTmpl,
		"template/node_sdl.tmpl",
	)
}

func templateNode_sdlTmpl() (*asset, error) {
	bytes, err := templateNode_sdlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/node_sdl.tmpl", size: 2314, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/node_sdl.tmpl": templateNode_sdlTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"node_sdl.tmpl": &bintree{templateNode_sdlTmpl, map[string]*bintree{}},
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
