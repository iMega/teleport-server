// Code generated by go-bindata.
// sources:
// schema.graphql
// DO NOT EDIT!

package schema

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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x50\x4d\x4e\x32\x41\x10\xdd\xf7\x29\x6a\xc2\x86\x2f\xf9\x4e\x30\x4b\xc3\x86\x85\x46\xa3\x1e\x60\xd2\x53\xca\x04\xa6\x1b\x7b\x1a\x12\x62\x48\x70\xe2\x4a\x3d\x83\x0b\xbd\xc0\x88\x92\x20\x04\xce\xf0\xfa\x46\xa6\xe7\x07\x41\xdd\x74\x55\x52\xef\xaf\x5f\x26\x7b\x9c\x46\x74\x2b\x88\x6e\x46\x6c\x26\x21\x9d\xf9\x21\x88\xd2\x91\x8d\x6c\xa2\x55\x48\xc7\xf5\x26\xa6\x42\xd8\xc9\x90\x2b\x48\xc9\x51\x3a\xe6\x76\x12\x87\xd4\xed\x04\xff\x42\x3a\xd1\x31\x07\x82\x48\xf6\x58\xf6\x2f\x74\x9f\x55\xdb\xfa\x37\xa4\x73\x6b\x12\x75\xed\x31\x97\x19\x9b\xc0\x4b\xb5\x08\xcf\x78\x73\x0f\x58\x60\xe5\x72\x72\x77\xd4\xed\x88\x44\x59\x36\x57\x91\xe4\x52\xab\xf4\xa8\xe5\x6b\xca\x0b\xb6\x98\xa3\x70\x33\xc2\x9c\xb0\x42\xe1\x72\x14\x58\x63\x8b\x77\x2c\xaa\x78\xa7\x46\xc7\x23\x69\x29\x49\x87\x03\x4e\x59\xd9\xec\x5b\xac\x45\x78\xc5\x06\x4b\xcf\xc4\xda\x3d\x61\xe3\x1e\xf1\x49\x58\xe2\x03\x0b\x6c\x5c\x8e\xa5\xbb\xaf\xce\x2e\xc7\xd6\xcd\xc8\x8f\xca\x11\xc5\x61\x9a\xd2\xac\x29\xa7\x54\x97\x86\x23\xcb\xb5\x7f\xd8\x04\xf1\x8d\x18\x4e\xf5\xb8\xb9\xec\x35\x76\xa4\xf5\x80\x23\x15\x88\x1d\xdb\xf7\xd3\x1e\x46\x59\xf6\xbb\xb4\x06\x52\x35\x5b\x8b\xfc\xa7\x9f\xe0\x7a\xdb\x45\xf4\xe4\x3f\xcb\x68\xfe\x42\x87\x12\x82\x28\x92\x36\x19\xf3\x5e\xbc\xa9\xf8\x0a\x00\x00\xff\xff\xa8\xd1\x53\x4e\x2a\x02\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 554, mode: os.FileMode(420), modTime: time.Unix(1529304574, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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

