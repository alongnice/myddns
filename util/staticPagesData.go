// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// static/pages/writing.html
package util

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

// Mode return file modify time
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


var _staticPagesWritingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x5d\x6f\x1b\xc7\xd5\xbe\xe7\xaf\x38\xef\xc0\x40\x24\x20\xbb\x2b\x51\xca\x22\xaf\xcc\x25\x60\x9b\x6e\x23\xd4\x71\x5a\x2b\x32\x7a\x27\xcc\xee\x0e\xb9\x63\xcd\xce\xac\x67\x66\x49\xb1\x04\x01\x35\x68\xda\x22\x48\x00\x07\x48\x01\x23\x75\xea\xa6\x68\x83\xa2\x68\xea\x5e\x1a\x45\x53\xff\x98\x88\xfe\xb8\xea\x5f\x28\x66\x3f\xc8\x25\xb9\xfa\x70\x24\x37\x69\x7c\x61\xce\xd7\x79\x66\xce\x79\xce\x3c\xab\x39\xad\x48\xc7\x0c\x18\xe6\x3d\x0f\x11\x8e\xda\x8d\x46\x2b\x22\x38\x6c\x37\x00\x5a\x31\xd1\x18\x82\x08\x4b\x45\xb4\x87\x52\xdd\xb5\xde\x44\xb3\x89\x48\xeb\xc4\x22\x77\x53\xda\xf7\xd0\x4f\xad\xdd\x2b\xd6\x35\x11\x27\x58\x53\x9f\x11\x04\x81\xe0\x9a\x70\xed\xa1\xed\xeb\x1e\x09\x7b\xa4\x62\xc7\x71\x4c\x3c\xd4\xa7\x64\x90\x08\xa9\x2b\x4b\x07\x34\xd4\x91\x17\x92\x3e\x0d\x88\x95\x75\x5e\x07\xca\xa9\xa6\x98\x59\x2a\xc0\x8c\x78\xeb\x4b\x30\x38\xd5\x91\x90\x15\x90\x3b\xb4\xd8\x4b\x53\xcd\x48\xbb\xd3\xb9\xb9\x63\xfd\xf0\x9d\x96\x93\x77\xcd\xc4\xff\x59\x16\x5c\x15\x42\x2b\x2d\x71\x02\xd7\x76\x76\xc0\xb2\xb2\x09\x46\xf9\x3e\x48\xc2\x3c\xa4\xf4\x90\x11\x15\x11\xa2\x11\x44\x92\x74\x3d\xe4\x28\x8d\x35\x0d\x1c\xbf\x34\xb4\x63\xca\xed\x40\x29\x74\x66\xd3\x40\xc4\xb1\x28\x6d\x5a\x4e\x1e\xe5\x46\xcb\x17\xe1\x30\xc3\x30\x03\x44\x9a\x26\x40\x2b\xa4\x7d\x08\x18\x56\xca\x43\x1c\xf7\x7d\x2c\x21\xff\xb1\x42\x2c\xf7\xc1\xef\xe5\xbf\x2a\xc2\xa1\x18\x58\x2a\x46\xb9\xd9\xbc\xa1\x89\x09\xa6\x9c\x48\x08\xad\x2e\x23\x07\x70\x27\x55\x9a\x76\x87\x56\x11\x2c\xcb\x27\x7a\x40\x32\xd2\xa1\xf8\xd7\xc2\xe5\xa1\xd1\xfc\xf6\x96\x2f\x31\x0f\x4b\x20\xcc\x68\x8f\x5b\x54\x93\x58\x59\x01\xe1\x9a\xc8\x0a\x06\x40\x4b\xf5\x7b\x70\x10\x33\xae\x3c\x64\xb2\x64\xcb\x71\x06\x83\x81\x3d\xd8\xb0\x85\xec\x39\xcd\xb5\xb5\x35\x47\xf5\x7b\x08\x72\xc2\x51\x73\x0d\x41\x44\x68\x2f\xd2\x79\xbb\x4b\x19\xf3\x10\x17\x9c\x20\x50\x5a\x8a\x7d\xe2\xa1\x20\x95\x92\x70\x7d\x4d\x30\x21\x51\x65\x2b\x28\x56\x58\x8c\x72\x12\xe0\xc4\x43\x52\xa4\x3c\x44\xd5\xe1\x3b\x82\xf2\xc5\xf1\x72\x6b\x04\x58\x52\x6c\x45\x34\x0c\x09\xf7\x90\x96\x29\x99\x7a\x1e\x4b\xab\x39\xbf\x97\x49\xda\xab\xe2\xc0\x43\x6b\xb0\x06\xcd\x4d\x68\x6e\x22\xe8\x8a\x20\x55\xd8\x67\xc4\x43\x5d\xcc\x14\x99\x8b\x04\x40\x2b\xc1\x3a\x82\xd0\x43\x6f\x37\x37\x60\xfd\xff\x71\x13\x9a\x60\xac\xd7\xad\x26\x34\xdf\xda\xa8\xf6\xad\xe6\xed\x37\x67\x7d\x68\x5a\xcd\x68\x93\x35\xad\x8d\xc8\x65\x4d\xd8\x88\x36\xab\x73\xd0\xfc\x19\x02\x67\x61\xab\x80\xca\x80\x11\x08\x0e\x3c\xb4\xde\x44\x10\x0c\x3d\xb4\xbe\x81\x40\x7a\x68\xd3\x2c\x6e\x99\xa8\xcf\xf3\xa4\xa5\xe0\xbd\xd9\x25\x29\xfa\xb3\x74\x70\xf0\x34\xb1\x9c\x90\xf6\x8b\xe4\x2c\x9b\x79\x12\xe7\x39\x9b\xdd\x4b\x4c\x39\x48\x61\x42\x61\x9a\x26\xd8\xc3\xbc\x23\x7b\x94\x5b\x5a\x24\x5b\xb0\xb1\x96\x1c\xa0\x59\x92\xd7\xe7\x2d\xb3\xe2\xd0\x72\x41\x74\xbb\x8a\x68\xd3\xde\xa8\xe6\x68\x57\xc8\x18\x70\xa0\xa9\xe0\xe6\x76\xe1\x3e\x41\x10\x13\x1d\x89\xd0\x43\x89\x50\xda\x88\x58\xc5\x49\x3f\xd5\x5a\x70\xd0\xc3\x84\x78\x48\xa5\x7e\x4c\xf5\x94\x62\x5f\x73\xf0\x35\xb7\x12\x49\x63\x2c\x87\x8b\x27\xf6\x85\xd6\x22\xde\x82\xe6\x5a\x72\x70\x19\xb5\x77\x70\x9f\xb4\x9c\x1c\x6f\x7e\x8f\xca\xe9\x8d\xa8\x31\xa2\x17\xd3\x20\x7a\x63\x61\xc1\xde\x9e\x09\x1e\x6a\x77\x6e\xee\x3c\x7d\xf0\xd1\xe4\x83\xcf\x27\xbf\xf9\x65\xcb\x89\xde\x58\xb0\x5b\x46\xde\xdb\x33\x92\xb1\x80\x3f\xbf\xd2\x44\xc8\xea\x49\x91\x26\x4b\xcb\x6a\x16\x06\x11\x09\xf6\x61\xd6\xb4\x28\x37\xd7\xa6\xc6\x14\xa0\x45\x79\x92\xea\x65\x73\x2b\x1b\x47\x45\x94\x25\x0e\xa9\x40\x85\x44\x77\xb8\xba\x89\x63\x82\x80\x86\x1e\xc2\x8c\x86\x5c\x21\xe8\x63\x96\x92\x59\x57\xf0\x80\xd1\x60\xbf\x1c\xb8\x66\x20\x49\xf8\x83\x94\xaf\xac\x22\x18\x8d\x68\x17\xc8\x5d\xb8\x64\x77\x6e\xee\xd8\x06\x0b\x4a\xc3\xf1\x38\xc8\x97\x8e\x46\x84\x87\xe3\x71\xed\x91\x19\xf6\x09\xab\x39\x72\x36\x6e\x2e\xb0\x9c\x1e\xa4\xce\x1e\xe0\x4a\x36\xb9\xf2\xe2\xfe\x93\x17\xbf\xfa\xf0\xe8\x1f\x1f\xaf\xd6\xed\xe2\x64\x70\x35\xe1\x76\x2a\x69\xfe\x1d\x61\x21\xe4\x2a\x11\xe1\x94\x85\xb2\x3b\x65\x21\x1f\x38\x9d\x85\xc2\xf0\xc2\x58\x28\xf0\xea\x59\xe8\x64\x93\x2b\xcf\xdf\xff\xd7\xf3\xbf\x3d\xfa\x3e\xb0\x10\x30\x91\x86\x5d\x86\x25\x99\x32\x51\x1d\x9a\xb2\x31\x1b\x3c\x9d\x91\x0a\xc0\x85\xb1\x52\xc1\xac\x67\xe6\xda\x74\xc1\x45\x50\xa2\x62\xcc\x58\x99\xa6\x7b\x11\x61\x09\x9a\x3b\xa0\x26\x07\x1a\xcc\x7f\x56\x9c\x6a\x12\x22\xf3\x5d\x33\x26\x4b\x82\x58\x87\x7f\x76\x95\xcc\x23\x93\x05\xa0\xc3\xd5\x76\x67\x7a\x73\xb6\xc3\x1b\x59\x70\xda\xdb\x9d\xe3\x7d\xab\x49\x0d\xc1\xb5\x14\xac\x92\x09\x25\x66\xd1\x2c\x52\x60\x34\xca\xe8\xdc\xee\x8c\xc7\xcb\x1a\xbf\xe4\xd2\x05\xb9\xb7\x43\x02\x49\x34\x9a\x7a\x58\xf4\xdb\xf9\xef\xf9\xdc\x2c\xb0\x4a\x57\xcb\xee\xbc\xbb\xf9\xe8\xf1\x2e\x37\x4e\x0c\xc3\xf2\xa2\xf3\x7c\x91\xb7\x7f\x7c\x7b\xf3\x3c\x5f\x62\x78\x59\x7e\x2a\x7a\x53\x47\x55\x1e\xe4\x5c\x4d\xb2\x45\xbe\x38\x40\xc7\xcb\x8e\x89\x33\x4d\xfa\x9b\x7b\x84\xe3\xec\x2d\x96\x33\xb1\x9d\xf4\x37\xaf\x17\x23\x33\xe5\x30\xa3\x76\x3e\x0c\xe6\x8f\xdf\xd3\x45\xe3\x4c\x92\x51\xdd\xbf\xfd\xf4\xfe\xa3\xc9\xbd\x2f\x26\xf7\x1e\x3d\xfb\xe4\xcf\xc7\xa4\x52\x1d\xcb\xdf\x30\x9b\xb3\xad\x53\xc9\x50\xfb\xe9\xc3\x3f\x3d\x7f\xf4\x87\xdd\x5b\x37\x4e\xcb\xdf\x3c\xb4\xc6\xe6\xa4\x54\x36\xa1\xda\x35\x6b\xa6\x01\xce\x2c\xb2\xf7\x43\x48\x54\x20\xa9\x4f\x42\x7f\x38\x9b\x2b\xb4\x6b\x96\xe8\x59\xb0\x77\x6f\xdd\xa8\x49\xf3\x39\xdd\x5b\x00\x38\x51\xfc\x9e\x7e\xf9\xc7\xdd\x5b\x37\x5e\x3c\x38\x7c\xfe\xe4\x93\xc9\x6f\x7f\x77\xf4\xd5\xef\x9f\x7d\xfa\x8b\xc9\xfb\x7f\x7d\xf6\xd5\xc7\x26\x8f\x27\x0f\xfe\x3e\xf9\xec\xf0\xeb\xc3\xf7\x26\x5f\xbc\xf7\xef\x7f\x7e\x6a\x5e\x63\x6a\xcb\x71\x70\x42\x2d\xb3\x8b\x4d\x13\x5b\xf9\x0e\x4d\xbe\x3e\xfc\x79\x39\x17\x0f\x69\x62\xd3\x84\x26\x36\x37\x57\xff\x24\x6d\xbd\x30\xbe\x42\x61\x1e\x0c\x0a\xb5\x3b\x79\xe3\x78\xc6\x8c\xef\x58\x12\x5c\xcf\xd4\x34\x7a\x25\x60\x85\xbb\x4e\x39\x24\xc5\x40\x79\x68\xe3\x58\xea\x0a\xdb\x3c\xfa\xed\xc6\x68\x64\x81\xc4\xbc\x47\xe0\x12\x7d\x1d\x2e\xf5\x61\xcb\x83\x9c\xca\x02\x71\x3c\x6e\x8c\x46\x97\xfa\xd9\x8f\x05\x84\x87\x60\x8d\xc7\x35\x9f\xbb\xf2\xe8\xa7\x73\x3f\x77\x82\x93\xf9\x3f\x7a\x7c\xf8\xfc\xf3\x0f\x8f\x1e\x1f\x1e\x3d\xfe\xcb\xe4\xe1\xc3\xc9\xbd\x8f\x5e\x82\xb2\x57\x2f\xa4\xee\xff\xbe\x90\xba\xcb\x42\xea\xd6\x0a\xa9\xfb\x6a\x84\xd4\xfd\xf6\x84\xd4\xbd\x60\x21\x9d\x46\x34\x5b\x32\x0b\xe7\xee\xb1\x42\xea\x2e\x0b\xe9\x82\x86\xb9\x53\x0d\x3b\x55\x54\xdd\x8b\x11\x55\xf7\x64\x51\x75\xeb\x44\xb5\xef\xda\x99\xae\x32\xec\xdc\x51\x82\xff\x57\x64\xd5\xbd\x68\x59\x75\x6b\x64\xd5\x3d\xa3\xac\xba\x67\x95\x55\xf7\x55\xc9\xaa\xfb\x9d\x92\xd5\x97\xa9\x4a\x1d\x57\x78\x6a\x39\xe6\xf0\x27\x57\xe8\x8c\xc7\x99\x49\xcb\xf0\x91\xe8\x7c\x41\x37\xe5\x59\xfd\x0c\x7a\x44\xff\x24\x25\x72\x78\xdb\x90\xe6\x33\xb2\xb2\x3a\x2a\xe0\xfa\x58\xc2\x5d\x33\x05\x1e\x0c\x28\x0f\xc5\xc0\x66\x22\xc0\xc6\xca\x56\x04\xcb\x20\xb2\x55\xea\x2b\x2d\x29\xef\xad\xac\xaf\x5e\x2e\xcc\x68\x17\x56\x0a\x33\xcf\x03\xa4\x70\x9f\xbc\xb3\xef\x65\x75\xd5\x55\x18\x4d\x8f\xbe\x88\x18\x49\xd2\x05\x0f\x90\x33\x2b\xb8\x62\x46\xa4\x5e\x41\x47\x4f\x3e\x9b\x7c\x79\xff\xe9\xaf\xef\x4d\x3e\x78\x88\xca\x37\x7f\x9e\x07\xf9\xff\xcb\x2e\x34\xca\xf3\xfb\xa4\x2b\x24\xc9\xde\x57\x1e\x42\x8d\x79\xd7\x97\x4b\x4d\xd3\xf3\x85\x22\x48\x63\xc2\xb5\xdd\x23\xfa\x3a\x23\xa6\x79\x75\xb8\x1d\xae\x14\x4f\xb5\x55\x3b\xa4\x59\xc9\x37\xf4\x20\xab\xf9\x56\x7c\xaf\x6c\x39\x2b\x50\x9c\x8a\x97\x4b\x5b\xf5\xbc\x8d\xd3\x8e\x52\x79\x89\xae\xda\x94\x73\x22\xdf\x7a\xf7\xed\x1b\x26\x86\x57\x82\x80\x28\xf5\x23\x32\x84\xed\x0e\x3a\x13\x4c\xf1\x26\x3b\x16\xa7\x98\x3f\x03\x56\x7e\xb7\x16\x80\x4a\x11\x94\x38\xb6\x03\xc1\x95\x60\xc4\xc6\x8c\x0e\x53\x6e\x07\x22\x76\x62\xcc\x71\x8f\x38\x78\x1f\x15\xac\xce\xf3\xb4\x5c\x8c\xfa\x9e\xf0\x74\x3e\x76\xde\x15\xfb\x84\x9f\x9b\x92\x92\x8e\x3c\xc8\x76\xc0\x1d\x1c\x04\x22\xe5\xda\xd1\x33\xfc\x45\x46\xea\x0b\x52\xa7\xb3\x72\x7c\x28\xac\xd2\x91\x4a\x64\xc1\x3b\x23\x1f\x2f\x9f\x0c\x46\x8e\xce\x6a\x55\x50\x8e\xbe\x7d\xae\x42\xac\x22\x7b\x16\xfb\xec\xee\x24\x52\x74\x29\x23\xd9\xdf\x1e\x19\x63\x6a\x8e\x32\x23\x82\x61\x5e\x73\x04\x0f\x5e\x1b\x8d\x66\xb5\xc2\xf1\xf8\xb5\x7c\x8d\x1a\x50\x1d\x44\x2b\xc5\xb2\xa9\xfe\x07\x58\xcd\xea\xec\x5b\x15\xe9\x5e\xd6\xcd\xe9\x94\x2f\x09\xde\xbf\x3c\xa7\xd1\x25\x50\x51\xda\xad\x02\x2d\x5f\xec\xb3\x00\x55\xaa\x91\x55\xb0\xfa\x9c\x3c\x19\x30\x24\x5d\x9c\x32\xfd\xcd\x9d\x1b\x67\x9f\xd8\xf2\xc3\xda\x68\x39\x91\x8e\x59\xfb\x3f\x01\x00\x00\xff\xff\x5f\x1d\x0c\x9b\x53\x1f\x00\x00")


func staticPagesWritingHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticPagesWritingHtml,
		"static/pages/writing.html",
	)
}

func staticPagesWritingHtml() (*asset, error) {
	bytes, err := staticPagesWritingHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/pages/writing.html", size: 8019, mode: os.FileMode(420), modTime: time.Unix(1599030966, 0)}
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
	"static/pages/writing.html": staticPagesWritingHtml,
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
		"pages": &bintree{nil, map[string]*bintree{
			"writing.html": &bintree{staticPagesWritingHtml, map[string]*bintree{}},
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
