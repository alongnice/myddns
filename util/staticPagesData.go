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


var _staticPagesWritingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\xdf\x8f\xdb\xc6\xf1\x7f\xbf\xbf\x62\xbe\x0b\x23\xd1\x7d\x61\x92\x77\xba\xb3\x90\x3a\xe2\x15\x89\x95\x36\x87\x5e\x9c\xd4\x17\x07\x79\x33\x56\xe4\x48\x5c\xdf\x72\x97\xde\x5d\x4a\xa7\x08\x02\xae\x41\xd3\x16\x41\x02\x38\x40\x0a\x04\x69\x52\x37\x45\x1b\x14\x45\x53\xf7\x31\x28\x9a\xe6\x8f\x89\xe5\x38\x4f\xfd\x17\x8a\x25\x29\x89\x94\x78\x3f\x5c\x5f\x1c\x27\x2f\x12\xf7\xc7\xcc\xce\xce\x67\x76\x3e\xe4\x6c\x3b\x32\x31\x07\x4e\x45\xdf\x27\x6f\x44\x64\x67\x6d\xad\x1d\x21\x0d\x77\xd6\x00\xda\x31\x1a\x0a\x41\x44\x95\x46\xe3\x93\xd4\xf4\x9c\x67\xc8\x62\x20\x32\x26\x71\xf0\x56\xca\x06\x3e\x79\xdd\xb9\xfe\x9c\x73\x45\xc6\x09\x35\xac\xcb\x91\x40\x20\x85\x41\x61\x7c\xb2\xfb\x82\x8f\x61\x1f\x4b\x72\x82\xc6\xe8\x93\x01\xc3\x61\x22\x95\x29\x4d\x1d\xb2\xd0\x44\x7e\x88\x03\x16\xa0\x93\x35\x2e\x02\x13\xcc\x30\xca\x1d\x1d\x50\x8e\xfe\xe6\x8a\x1a\x9a\x9a\x48\xaa\x92\x92\x9b\xac\x58\xcb\x30\xc3\x71\xa7\xd3\xb9\xba\xef\xfc\xf4\xe5\xb6\x97\x37\xed\xc0\xff\x39\x0e\x3c\x2f\xa5\xd1\x46\xd1\x04\xae\xec\xef\x83\xe3\x64\x03\x9c\x89\x03\x50\xc8\x7d\xa2\xcd\x88\xa3\x8e\x10\x0d\x81\x48\x61\xcf\x27\x9e\x36\xd4\xb0\xc0\xeb\xce\x04\xdd\x98\x09\x37\xd0\x9a\x9c\x59\x34\x90\x71\x2c\x67\x32\x6d\x2f\xf7\xf2\x5a\xbb\x2b\xc3\x51\xa6\xc3\x76\xa0\xb2\x8f\x00\xed\x90\x0d\x20\xe0\x54\x6b\x9f\x08\x3a\xe8\x52\x05\xf9\x9f\x13\x52\x75\x00\xdd\x7e\xfe\xaf\x23\x1a\xca\xa1\xa3\x63\x92\x8b\x55\x05\xad\x4f\x28\x13\xa8\x20\x74\x7a\x1c\x0f\xe1\x66\xaa\x0d\xeb\x8d\x9c\xc2\x59\x4e\x17\xcd\x10\x51\xcc\x85\x01\xda\x74\x66\x34\xa9\x2e\xef\x74\x15\x15\xe1\x4c\x11\xe5\xac\x2f\x1c\x66\x30\xd6\x4e\x80\xc2\xa0\x2a\xe9\x00\x68\xeb\x41\x1f\x0e\x63\x2e\xb4\x4f\x6c\x94\x5c\xf6\xbc\xe1\x70\xe8\x0e\xb7\x5c\xa9\xfa\x5e\x73\x63\x63\xc3\xd3\x83\x3e\x81\x1c\x70\xd2\xdc\x20\x10\x21\xeb\x47\x26\x7f\xee\x31\xce\x7d\x22\xa4\x40\x02\xda\x28\x79\x80\x3e\x09\x52\xa5\x50\x98\x2b\x92\x4b\x45\x4a\x4b\x41\x31\xc3\xe1\x4c\x60\x40\x13\x9f\x28\x99\x8a\x90\x94\xbb\x6f\x4a\x26\x96\xfb\x67\x4b\x13\xa0\x8a\x51\x27\x62\x61\x88\xc2\x27\x46\xa5\x38\xdf\x79\xac\x9c\x66\x75\x2d\x1b\xb4\xcf\xcb\x43\x9f\x6c\xc0\x06\x34\xb7\xa1\xb9\x4d\xa0\x27\x83\x54\xd3\x2e\x47\x9f\xf4\x28\xd7\x58\xf1\x04\x40\x3b\xa1\x26\x82\xd0\x27\x2f\x35\xb7\x60\xf3\x47\xb4\x09\x4d\xb0\xd2\x9b\x4e\x13\x9a\x2f\x6e\x95\xdb\x4e\xf3\xb5\x67\x16\x6d\x68\x3a\xcd\x68\x9b\x37\x9d\xad\xa8\xc5\x9b\xb0\x15\x6d\x97\xc7\xa0\xf9\x06\x01\x6f\x69\xa9\x80\xa9\x80\x23\x04\x87\x3e\xd9\x6c\x12\x08\x46\x3e\xd9\xdc\x22\xa0\x7c\xb2\x6d\x27\xb7\xad\xd7\xab\x38\x19\x25\x45\x7f\x71\x48\x8a\xf6\x22\x1c\x3c\x3a\x0f\x2c\x2f\x64\x83\x22\x38\x67\x8f\x79\x10\xe7\x31\x9b\x9d\x4b\xca\x04\x28\x69\x5d\x61\x1f\xad\xb3\x47\x79\x43\xf5\x99\x70\x8c\x4c\x2e\xc3\xe6\xa5\xe4\x90\xac\x06\xb9\x92\xc3\x63\x62\x98\x3b\x71\xe8\xb4\x40\xf6\x7a\x1a\x8d\x7d\xde\x2a\xc7\x6b\x4f\xaa\x18\x68\x60\x98\x14\xf6\xa4\xd1\x01\x12\x88\xd1\x44\x32\xf4\x49\x22\xb5\xb1\x09\xad\xb4\xe1\x6e\x6a\x8c\x14\x60\x46\x09\xfa\x44\xa7\xdd\x98\x99\x39\xdc\x5d\x23\xa0\x6b\x84\x93\x28\x16\x53\x35\x5a\xb6\xbe\x2b\x8d\x91\x71\xbe\x81\x67\xc9\xce\x3e\x1d\x60\xdb\xcb\xf5\x55\xd7\x28\x59\x6f\x13\x1c\x47\xb3\x1c\x12\xd1\xa5\xa5\x09\x37\x6e\x58\x47\x92\x9d\xce\xd5\xfd\xfb\x1f\xbd\x3b\x7d\xfb\x93\xe9\x6f\x7f\xd5\xf6\xa2\x4b\x4b\x72\xab\x9a\x6f\xdc\xb0\xe9\xa3\xba\xc9\xe5\xa9\xd6\x45\x4e\x5f\xc9\x34\x81\xb2\x97\x4b\x93\x39\xed\x22\x2f\x7b\x5c\xc7\x4e\x13\xec\x43\x26\x9b\x0d\x93\x9d\xb6\x97\x3d\xd4\xc8\x2f\xe1\xa5\x63\x67\x73\xa3\x66\x9d\x1a\xb3\x82\x08\x83\x03\x58\x3c\x3a\x4c\xd8\x13\xbb\xb2\x76\x8d\x2e\x80\x36\x13\x49\x6a\x56\xf5\x39\x59\x3f\x29\x50\x56\x34\x64\x92\x14\x74\xd1\x11\xfa\x2a\x8d\x91\x00\x0b\x7d\x42\x39\x0b\x85\x26\x30\xa0\x3c\xc5\x45\x53\x8a\x80\xb3\xe0\x60\xd6\x71\xc5\xaa\xc4\xf0\x27\xa9\x68\xac\x13\x18\x8f\x59\x0f\xf0\x16\x5c\x70\x3b\x57\xf7\x5d\xab\x0b\x66\x82\x93\x49\x90\x4f\x1d\x8f\x51\x84\x93\xc9\x31\x46\x57\x9c\x5d\x32\x3a\xdf\xa9\xf5\xc5\xdc\x94\x7a\x0d\x00\xcf\x65\xc3\x8d\x6f\x3e\xf8\xf2\x9b\x5f\xbf\x73\xef\x9f\xef\xad\xd7\xaf\x74\x1c\x5e\x95\xc3\xfc\x24\x03\x14\x0a\x9d\xc8\x70\x0e\xd0\xac\x39\x07\x28\xef\x38\x1d\xa0\x42\xf0\x1c\x01\x2a\x34\x1e\x07\x50\x27\x1b\x6e\x3c\x78\xeb\xdf\x0f\xfe\x7e\xf7\x07\x0c\x50\xc0\x65\x1a\xf6\x38\x55\x38\x07\xa9\xdc\x35\x07\x6a\xd1\x79\x3a\x58\x25\x05\xe7\x08\x58\x49\xeb\x71\xa0\x5d\x99\x4f\x39\x3f\xb4\x74\x4c\x39\x9f\xc5\xf2\x8d\x08\x79\x42\x2a\x86\x1a\x3c\x34\x60\x7f\x9c\x38\x35\x18\xda\x24\x9b\x89\xd4\x24\xd9\xba\x35\x8a\xce\x73\x48\xfe\x99\x93\x3a\x42\xef\x76\xe6\x47\x6f\x37\xdc\xcb\x1d\x78\x1a\x31\xec\x76\xce\x83\x1a\x6a\x42\x51\x0a\xa3\x24\x2f\x45\xde\xcc\xba\xe2\xb1\x08\xb9\xf1\x38\x0b\x9f\xdd\xce\x64\x52\xb7\xc1\xc7\xe3\xb8\x7d\x0c\x94\x7d\xe7\x2f\x9c\x97\x37\xcf\xe8\xc0\x7c\xf2\xe3\x72\x62\xd9\xd2\x52\xb3\xea\xcc\xbc\xf7\xd1\x1c\xba\x32\x71\x75\xd2\xa3\xbc\x2c\xed\xbe\xf2\xda\xf6\x63\x7f\x49\xca\xe0\x66\xc9\x60\xfb\x06\x0a\x9a\x7f\xe9\x56\xc1\x25\x3b\xf7\x3f\xb8\x3b\xbd\xfd\xe9\xf4\xf6\xdd\xaf\xdf\xff\xcb\xf9\x61\x9a\xa7\xe2\x2c\xb1\x75\xe5\x21\xa9\xcd\xd9\x36\xf1\xd7\xbe\x74\x67\xaf\xac\x19\xe0\x15\xd3\xf3\x90\xd8\x4d\x06\xdb\x2f\x14\x3d\x8b\x84\x6c\x7b\xdd\xbc\x1b\xec\x57\xd1\xe9\xb9\xf8\x71\x1c\xb4\xcc\xfc\x54\x9d\xe1\x4c\xdd\xbf\xf3\xe7\x07\x77\xff\x78\xfd\xda\xde\x79\x43\x50\x5e\xbd\xee\x84\x59\xc7\x5d\xb7\x73\xe6\xee\xce\x24\xb2\xcf\xcc\x10\x75\xa0\x58\x17\xc3\xee\x68\x31\x56\xd0\xc2\xe2\xfc\x65\xae\xbf\x7e\x6d\xaf\xf6\xf4\x55\x48\x65\x49\xc5\x89\xcc\x72\xff\xb3\x3f\x5d\xbf\xb6\xf7\xcd\x47\x47\x0f\xbe\x7c\x7f\xfa\xbb\xdf\xdf\xfb\xe2\x0f\x5f\x7f\xf8\xcb\xe9\x5b\x7f\xfb\xfa\x8b\xf7\xec\x49\x9a\x7e\xf4\x8f\xe9\xc7\x47\x5f\x1d\xbd\x39\xfd\xf4\xcd\xff\xfc\xeb\x43\xfb\xd9\xae\x2f\x7b\x1e\x4d\x98\x63\x57\x71\x59\xe2\xea\xae\xc7\x92\xaf\x8e\x7e\x31\x1b\x8b\x47\x2c\x71\x59\xc2\x12\x57\xa0\x29\xf5\x27\x03\x57\xcb\x28\x75\x03\x19\x7b\x01\x33\xa3\x9b\x5a\x8a\x1f\x33\xf4\xb3\x02\xd2\x77\x4a\x70\x99\xbf\x42\x69\x3f\x4e\xf5\xe9\x21\xd4\xc9\x27\x9e\x47\x00\x59\x20\xa8\x42\x5a\x1f\x38\x73\x28\xe7\xa6\x2d\x42\xa9\x33\xeb\x52\x72\xa8\x7d\xb2\x75\x6c\x24\x15\xb2\x79\x28\xec\xac\x8d\xc7\x0e\x28\x2a\xfa\x08\x17\xd8\x45\xb8\x30\x80\xcb\x3e\xe4\x91\x55\x68\x9c\x4c\xd6\xc6\xe3\x0b\x83\xec\xcf\x01\x14\x21\x38\x93\x49\xed\xcb\xcd\xcc\xf8\xb3\x04\x63\xc5\x8a\x93\x03\xf2\xde\xe7\x47\x0f\x3e\x79\xe7\xde\xe7\x47\xf7\x3e\xff\xeb\xf4\xce\x9d\xe9\xed\x77\x1f\x3d\x36\xbe\x7d\xca\x69\x3d\x1c\xe5\x54\x4d\x3e\xc7\x38\x6e\x7d\x7f\x09\xa8\xb5\x4a\x40\xad\x5a\x02\x6a\x3d\x91\x04\xd4\x7a\x62\x09\x68\xee\xdf\x6c\xca\xc2\xb9\xd7\x8f\x25\xa0\xd6\x2a\x01\x2d\x65\xfe\xd6\x3c\xf3\x9f\x81\x8c\x5a\xe7\x43\x46\xad\x93\xc9\xa8\x55\x47\x46\x83\x96\x9b\xf1\x11\xa7\x9e\xe5\x9b\xd2\x88\x4e\x10\x43\x57\x60\xda\x72\x31\x4c\xdd\x40\x78\x7d\x34\xbb\xaf\xb8\x49\x94\x7c\xd7\x64\xd4\x7a\x72\xc9\xa8\x55\x43\x46\xad\x33\x92\x51\xeb\xac\x64\xd4\xfa\x36\xc9\xa8\xf5\x7d\x20\xa3\x87\x29\x48\x1f\x57\x73\x6e\x7b\x76\x4f\x35\x85\xfa\xda\x52\x7a\xa5\x7c\x9e\x2c\xbc\x22\x8c\x33\xcc\x2e\x61\x1c\x6e\x7f\x51\xe5\xfe\xe9\x2a\xa4\x07\xf5\x95\xfc\xac\x12\x9e\x09\x6a\xf6\x06\x5e\x86\xcd\xad\x79\x96\xe7\xb2\xaf\xc9\x4e\xdb\x4b\x4e\xbe\x3e\xb0\xf0\x64\x1b\x69\xdb\x00\x4a\x4c\x3e\xa1\x97\x8a\xac\xa0\x0f\x7d\x34\x3f\x4f\x51\x8d\x5e\xb3\x51\xd6\xe5\xd8\x58\x1f\x17\xea\x06\x54\xc1\x2d\x3b\x04\x3e\x0c\x99\x08\xe5\xd0\xe5\x32\xa0\x56\xca\xd5\x48\x55\x10\xb9\x3a\xed\x6a\xa3\x98\xe8\x37\x36\xd7\x9f\x2d\xc4\x58\x0f\x1a\x85\x98\xef\x03\xd1\x74\x80\x2f\x1f\xf8\xd9\xa5\xcf\x3a\x8c\xe7\x6e\x59\xd6\x18\x29\xec\x81\x0f\xc4\x5b\xdc\x06\x51\x8e\xca\x34\xc8\xbd\x2f\x3f\x9e\x7e\xf6\xc1\xfd\xdf\xdc\x9e\xbe\x7d\x87\xcc\xaa\x7c\x79\xe0\xe6\xbf\xab\x5b\x58\x9b\xd9\xdf\xc5\x9e\x54\x98\x95\x32\x7c\x42\xd6\xaa\x5b\x5f\xad\x3d\xcf\xed\x0b\x65\x90\xc6\x28\x8c\xdb\x47\xf3\x02\x47\xfb\xf8\xfc\x68\x37\x6c\x14\x55\x91\x75\x37\x64\xd9\x7d\x54\xe8\x43\x76\x21\x55\xda\x7b\x69\xc9\x45\x49\xf2\x54\x7d\x39\x31\x94\xed\x5d\x3b\xcd\x94\x52\xf9\x68\xdd\x65\x42\xa0\x7a\xf1\xd5\x97\xf6\xac\x0f\x9f\x0b\x02\xd4\xfa\x67\x38\x82\xdd\x0e\x39\x83\x9a\x72\x21\xe5\x58\x55\x45\x0d\xe3\x0c\xea\xf2\x5c\xb0\xa4\x68\xc6\x14\x8a\xc6\x6e\x20\x85\x96\x1c\x5d\xca\xd9\x28\x15\xd9\x37\x4c\x4c\x05\xed\xa3\x47\x0f\x48\x01\x6c\x15\xaa\xd5\x2a\xf4\x0f\x04\xaa\x47\x06\xe8\x55\x79\x80\xe2\x91\x51\x99\x21\x92\xfb\xd9\xf2\x37\x0d\x02\x99\x0a\xe3\x99\x85\xfe\x65\x50\xea\x2b\xce\xa7\x03\x73\xbc\x37\x66\xfb\x28\xf9\x16\xfc\x33\x22\xf2\xf0\xe1\x60\x73\xd2\x59\xa5\x0a\xd0\xc9\x13\x81\x56\x48\x75\xe4\x2e\xbc\x9f\x1d\xa0\x44\xc9\x1e\xe3\x98\xbd\xbf\x65\x98\xe9\x0a\x68\x36\x19\x86\xf9\xb5\x02\xf8\xf0\xf4\x78\xbc\xb8\x0e\x98\x4c\x9e\xce\xe7\xe8\x21\x33\x41\xd4\x28\xa6\xcd\x79\x20\xa0\x7a\x71\x01\x77\xb9\x94\xc2\x57\xf3\xe7\x7c\x28\x63\xb4\x67\x2b\xb9\x7a\xa6\xa8\xb8\xd6\x29\x2b\x5a\x3d\xdd\x67\x51\x54\xba\x6e\x28\x2b\xab\x8f\xca\x93\x15\x86\xd8\xa3\x29\x37\xff\xfb\xe6\x26\x19\xd5\xce\x08\xf6\x24\xb2\xdd\x93\x7d\x5d\x3a\x24\x16\x17\x7a\x93\x1e\x82\x0f\x02\x87\xf0\xfa\x4b\x7b\x2f\x1a\x93\x5c\xc3\x5b\x29\x6a\xd3\x98\x93\xaa\x9d\xe2\xca\x04\x45\x83\xf4\xd1\x90\x8b\xc4\xcb\xd8\xbf\x3a\xae\x51\x84\xcb\x22\x42\x21\x0d\x47\xda\x50\x83\x41\x94\xbd\x1d\xfa\x0b\x73\x1a\x65\x42\xb6\x09\x31\x93\xc9\x24\xf6\xad\x84\xef\x6f\xc3\x53\x4f\x15\xca\x0d\x35\xa9\xf6\xfd\xe6\xc6\x46\x59\xea\x84\x28\xce\x2d\xac\x44\x70\xa1\x5f\x27\x52\x68\x7c\x15\x0f\x4d\xa1\xe7\xab\xa3\xa3\xc9\x92\x3f\x67\xcc\x9e\xfb\x2b\x8f\x4f\x34\xbb\xc2\xa0\x1a\x50\xde\x28\x46\x2e\xc2\x25\xf8\x7f\xd8\xdc\xd8\xd8\x58\xaf\x22\xd0\xf6\x22\x13\xf3\x9d\xff\x06\x00\x00\xff\xff\x2d\xd3\xe4\xbc\x7a\x25\x00\x00")


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

	info := bindataFileInfo{name: "static/pages/writing.html", size: 9594, mode: os.FileMode(420), modTime: time.Unix(1599713878, 0)}
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
