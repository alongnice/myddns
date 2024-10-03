// Code generated by go-bindata.
// sources:
// static/pages/writing.html
// DO NOT EDIT!

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

var _staticPagesWritingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xdf\x6f\xdc\xc6\xf1\x7f\xd7\x5f\xb1\xdf\x85\x91\x9c\xbe\x30\x49\xe9\x24\x1f\x52\xe7\xa8\x22\xf1\xa5\x88\x50\xd9\x4d\xad\x28\x48\x9f\x8c\x3d\x72\xee\xb8\x16\xb9\x4b\xef\x2e\xef\x74\x39\x1c\xe0\x04\x75\x9a\x06\x09\xea\xa0\x49\x91\xa6\x49\x9d\x14\x4d\x10\x14\x8d\xdd\xa7\x36\x70\xea\xe4\x8f\x89\xce\x92\x9f\xfa\x2f\x14\xcb\x1f\x47\xf2\x8e\x92\x0e\xb6\x2c\xab\x7e\xb9\xe3\xee\xce\xcc\xce\xce\xe7\xb3\x33\x24\x97\x4d\x4f\x05\x3e\xf2\x09\xeb\xda\xf8\x0d\x0f\xaf\x2d\x2c\x34\x3d\x20\xee\xda\x02\x42\xcd\x00\x14\x41\x8e\x47\x84\x04\x65\xe3\x48\x75\x8c\xe7\x70\x3e\xe0\x29\x15\x1a\x70\x2d\xa2\x3d\x1b\xbf\x6e\x6c\xbd\x60\x5c\xe0\x41\x48\x14\x6d\xfb\x80\x91\xc3\x99\x02\xa6\x6c\xbc\xfe\x92\x0d\x6e\x17\x0a\x7a\x8c\x04\x60\xe3\x1e\x85\x7e\xc8\x85\x2a\x88\xf6\xa9\xab\x3c\xdb\x85\x1e\x75\xc0\x88\x1b\x67\x11\x65\x54\x51\xe2\x1b\xd2\x21\x3e\xd8\xcb\x33\x66\x48\xa4\x3c\x2e\x0a\x46\xae\xd2\x74\x2e\x45\x95\x0f\x6b\x17\x7f\xd5\x6a\x5d\xda\x6c\x5a\x49\x4b\xf7\xff\x9f\x61\xa0\x17\x39\x57\x52\x09\x12\xa2\x0b\x9b\x9b\xc8\x30\xe2\x01\x9f\xb2\x6d\x24\xc0\xb7\xb1\x54\x03\x1f\xa4\x07\xa0\x30\xf2\x04\x74\x6c\x6c\x49\x45\x14\x75\xac\x76\xa6\x68\x06\x94\x99\x8e\x94\x78\x6e\x55\x87\x07\x01\xcf\x74\x9a\x56\x12\xe4\x85\x66\x9b\xbb\x83\xd8\x86\xee\x00\xa1\x2f\x11\x6a\xba\xb4\x87\x1c\x9f\x48\x69\x63\x46\x7a\x6d\x22\x50\xf2\x67\xb8\x44\x6c\xa3\x76\x37\xf9\x97\x1e\x71\x79\xdf\x90\x01\x4e\xd4\xca\x8a\x3a\x24\x84\x32\x10\xc8\x35\x3a\x3e\xec\xa0\xab\x91\x54\xb4\x33\x30\xd2\x58\x19\x6d\x50\x7d\x00\x36\x51\x46\xa8\x49\x32\xa7\x71\x79\x7a\xa3\x2d\x08\x73\x33\x43\xc4\xa7\x5d\x66\x50\x05\x81\x34\x1c\x60\x0a\x44\xc1\x06\x42\x4d\xd9\xeb\xa2\x9d\xc0\x67\xd2\xc6\x9a\x24\xe7\x2d\xab\xdf\xef\x9b\xfd\x15\x93\x8b\xae\x55\x5f\x5a\x5a\xb2\x64\xaf\x8b\x51\x82\x37\xae\x2f\x61\xe4\x01\xed\x7a\x2a\xb9\xee\x50\xdf\xb7\x31\xe3\x0c\x30\x92\x4a\xf0\x6d\xb0\xb1\x13\x09\x01\x4c\x5d\xe0\x3e\x17\xb8\x30\x15\x4a\x25\x0c\x9f\x32\x70\x48\x68\x63\xc1\x23\xe6\xe2\x62\xf7\x55\x4e\xd9\x74\x7f\x36\x35\x46\x44\x50\x62\x78\xd4\x75\x81\xd9\x58\x89\x08\x26\x2b\x0f\x84\x51\x2f\xcf\xa5\x39\xfb\x22\xdf\xb1\xf1\x12\x5a\x42\xf5\x55\x54\x5f\xc5\xa8\xc3\x9d\x48\x92\xb6\x0f\x36\xee\x10\x5f\x42\x29\x12\x08\x35\x43\xa2\x3c\xe4\xda\xf8\x62\x7d\x05\x2d\xff\x84\xd4\x51\x1d\x69\xed\x65\xa3\x8e\xea\x2f\xaf\x14\xdb\x46\xfd\xb5\xe7\xf2\x36\xaa\x1b\x75\x6f\xd5\xaf\x1b\x2b\x5e\xc3\xaf\xa3\x15\x6f\xb5\x38\x86\xea\x6f\x60\x64\x4d\x4d\xe5\x50\xe1\xf8\x80\x9c\x1d\x1b\x2f\xd7\x31\x72\x06\x36\x5e\x5e\xc1\x48\xd8\x78\x75\x4a\xb8\xa9\x01\x28\x43\xa6\x04\x67\xdd\xc9\x76\x49\x9b\x39\x31\x2c\x32\xa1\x98\xe5\xd2\x5e\x4a\xd3\xec\x32\xa1\xb3\x66\x6f\xbc\x3d\x09\x65\x48\x70\x1d\x12\x7d\xa9\x83\x3e\x48\x1a\xa2\x4b\x99\xa1\x78\x78\x1e\x2d\x9f\x0b\x77\xf0\x2c\xd9\x05\xef\x1f\xc0\x65\xdf\x08\x5c\xa3\x81\x78\xa7\x23\x41\xe9\xeb\x95\x22\x6f\x3b\x5c\x04\x88\x38\x8a\x72\xa6\x77\x1c\xe9\x01\x46\x01\x28\x8f\xbb\x36\x0e\xb9\x54\x38\xf6\x6c\x22\xde\x8e\x94\xe2\x0c\xa9\x41\x08\x36\x96\x51\x3b\xa0\x6a\x02\x7b\x5b\x31\xd4\x56\xcc\x08\x05\x0d\x88\x18\x4c\x7b\xdf\xe6\x4a\xf1\x20\x59\xc0\xf3\x78\x6d\x93\xf4\xa0\x69\x25\xf6\xca\x73\x14\xbc\xd7\x79\xce\x07\x35\x4d\x0d\xef\xdc\x94\xc0\x95\x2b\x3a\x8c\x78\xad\x75\x69\xf3\xfe\xa7\xef\x8f\xdf\xfd\x62\xfc\xd1\xdb\x4d\xcb\x3b\x37\xa5\x37\x6b\xf9\xca\x15\x9d\x46\xca\x8b\x9c\x16\xd5\x21\x32\xba\x82\x47\x21\x2a\x46\xb9\x20\xec\x93\x36\xf8\xc5\x88\xcb\xc0\xa8\x23\x7d\x11\xeb\xc6\xc3\x78\xad\x69\xc5\x17\x15\xfa\x53\x78\xc9\xc0\x58\x5e\xaa\x98\xa7\xc2\x2d\xc7\x03\x67\x1b\xe5\x97\x06\x65\x7a\xe7\xce\xcc\x5d\x61\x2b\xcd\xe8\x4d\xca\xc2\x48\xcd\x1a\x35\xe2\x7e\x9c\x42\x2d\x88\x4b\x39\x4e\x4b\x47\x8b\xc9\x4b\x24\x00\x8c\xa8\x6b\x63\xe2\x53\x97\x49\x8c\x7a\xc4\x8f\x20\x6f\x72\xe6\xf8\xd4\xd9\xce\x3a\x2e\x68\x93\xe0\xfe\x2c\x62\xb5\x45\x8c\x86\x43\xda\x41\x70\x0d\x9d\x31\x5b\x97\x36\x4d\x6d\x0b\x4d\x14\x47\x23\x27\x91\x1d\x0e\x81\xb9\xa3\xd1\x5a\x5a\x65\x2a\xdc\x3f\x5d\x9e\xcf\x38\x5e\xed\x74\x89\x2b\x05\xa7\x13\xa0\x34\x94\x13\x57\xaa\x2d\x20\xf4\x42\x3c\x5c\x7b\xf0\xf1\x0f\x0f\x7e\xf3\xde\xee\xdd\x0f\x16\xab\x67\x3a\x88\x6e\xa5\x4c\x74\xea\xf9\xe5\x32\x19\x72\x77\x82\x52\xd6\x9c\xa0\x94\x74\x1c\x8d\x52\xa6\x78\x72\xfc\x3a\x66\xcf\x8f\x91\x5f\xa9\xc5\x83\xf8\xd5\x8a\x87\x6b\xfb\x37\xbe\xdf\xbf\x7d\xe7\x69\xe7\x97\xe3\xf3\xc8\xed\xf8\x44\xc0\x04\xa9\x62\xd7\x04\xad\xbc\xf3\x68\xc4\x8a\x06\x4e\x8e\x6f\x8f\x79\x25\xb3\xfc\x7b\x78\x02\x16\xcc\x1e\x44\xc2\x0b\x13\x91\xe3\x63\x9f\x0c\x88\xef\x67\x7b\xf3\x8a\x07\x7e\x88\x4b\x8e\x2a\xd8\x51\x48\xff\x18\x41\xa4\xc0\xd5\x25\x3b\x56\xa9\x28\xd9\x55\x73\xa4\x9d\xc7\x70\x2b\x11\x07\xa9\xc5\xe4\x7a\x6b\x92\x4a\xd6\xdd\x8d\x24\x80\x47\xdd\x66\xac\xb7\x8e\xe3\x46\xa3\x82\x8b\x9c\x29\xc1\xfd\x02\xf5\x32\xef\xd2\xcb\x94\x73\xc3\x61\xcc\x9f\xf5\xd6\x68\x54\xb5\xc0\x93\x09\xdc\x26\x38\x42\x3f\x49\xa6\xc1\x4b\x9a\x73\x06\x30\x11\x3e\xa9\x20\x16\x3d\x2d\x34\xcb\xc1\x4c\x7a\x1f\x2d\xa0\x33\x82\xb3\x42\x8f\x72\xeb\xbd\xfe\xca\x6b\xab\x27\x7e\xcb\x1d\xc3\x4d\xc3\xde\xea\x15\x60\x24\x79\x7d\x52\x06\x17\xaf\xdd\xff\xf8\xce\xf8\xe6\x57\xe3\x9b\x77\xf6\x3e\xfc\xfa\xf8\x30\x4d\x72\x71\x9c\xd8\xda\x7c\x07\x57\x26\x6d\x5d\xc8\x2a\x1f\xe1\xe2\x07\xa0\x18\xf0\x92\xeb\x09\x25\xd6\xc3\xde\xea\x4b\x69\x4f\x9e\x91\x75\xaf\x99\x74\x23\xfd\xac\x7d\xf4\xcd\xc0\x49\x6c\xb4\xd8\xfd\x48\xcc\xb1\xa7\xee\xdf\xfa\x72\xff\xce\x5f\xb6\x2e\x6f\x1c\x37\x04\xc5\xd9\xab\x76\x98\x0e\xdc\x96\x96\x99\x84\x3b\xd6\x88\x5f\x5e\xb8\x20\x1d\x41\xdb\xe0\xb6\x07\xf9\x58\x5a\x16\xf2\xfd\x17\x87\x7e\xeb\xf2\x46\xe5\xee\x2b\x15\x95\x29\x13\x87\x56\x96\xfb\xdf\xfc\x75\xeb\xf2\xc6\x83\x4f\xaf\xef\xff\xf0\xe1\xf8\x4f\x7f\xde\xbd\xf7\xf9\xde\x27\xbf\x1e\xdf\xf8\xfb\xde\xbd\x0f\xf4\x4e\x1a\x7f\xfa\x8f\xf1\x67\xd7\x7f\xbc\xfe\xd6\xf8\xab\xb7\xfe\xf3\xef\x4f\x3c\xa5\x42\x79\xde\xb2\x48\x48\x0d\x3d\x8b\x49\x43\x53\xb6\x2d\x1a\xfe\x78\xfd\xcd\x6c\x2c\x18\xd0\xd0\xa4\x21\x0d\x4d\x06\xaa\xd0\x1f\xf6\x4c\xc9\xbd\xc8\x74\x78\x60\x39\x54\x0d\xae\x4a\xce\x7e\x4a\xc1\x8e\xdf\x4a\x3e\xd1\x02\x17\xc7\xcb\xe5\x01\xa1\xfa\xb9\xeb\x28\x0a\xb5\x12\xc1\xe3\x20\x90\x06\x82\x08\x20\xd5\xc4\x99\x40\x39\x71\x2d\xa7\x52\x2b\xeb\x12\xbc\x2f\x6d\xbc\x72\x20\x93\x52\xdd\x84\x0a\x6b\x0b\xc3\xa1\x81\x04\x61\x5d\x40\x67\xe8\x59\x74\xa6\x87\xce\xdb\x28\x61\x56\x6a\x71\x34\x5a\x18\x0e\xcf\xf4\xe2\x3f\x03\x01\x73\x91\x31\x1a\x55\xde\xdc\x64\xce\xcf\x43\xc6\x92\x17\x47\x10\xf2\xce\xef\xf6\xbf\x78\x6f\xf7\xdb\xeb\xbb\xdf\xfe\x6d\x7c\xeb\xd6\xf8\xe6\xfb\x8f\xce\x8d\xc7\x5f\x72\x1a\x4f\xac\xe4\x34\xfe\x77\x4b\x4e\x63\xb6\xe4\x34\x2a\x4b\x4e\xe3\x54\x96\x9c\xc6\xa9\x2d\x39\x93\xf8\xc6\x22\x79\x70\xb7\x0e\x2c\x39\x8d\xd9\x92\x33\x95\xeb\x1b\x93\x5c\x3f\x47\xf9\x69\x1c\x4f\xf9\x69\x1c\x5e\x7e\x1a\x55\xe5\xa7\xd7\x30\xe3\x0a\xe4\x13\x4b\x57\x98\xc2\x88\x0c\x01\x5c\x93\x41\xd4\x30\xc1\x8d\x4c\x87\x59\x5d\x50\xeb\xaf\x98\xa1\x17\x3e\xe9\xf2\xd3\x38\xbd\xe5\xa7\x51\x51\x7e\x1a\x73\x96\x9f\xc6\xbc\xe5\xa7\xf1\x38\xcb\x4f\xe3\x29\x2f\x3f\xe3\x1b\xff\xdc\xfd\xee\x0f\x0f\x6e\xbc\xbf\x77\xef\xf6\x93\x29\x43\x5b\x12\x04\x8b\xdf\xff\x1c\xc5\xde\xbd\x3f\x7e\x37\xbe\xf7\xd1\xde\x87\x5f\xdf\x7f\xe7\x5f\x71\x6c\x4f\xe0\xe1\x36\xf7\x4e\x13\x22\x6f\xe5\xb7\xd6\x59\xdf\x68\x54\x45\xe4\x6c\x34\x23\x71\x85\x0f\x39\xdf\xca\xc2\x87\x73\x6d\xf7\xdb\xbb\xbb\x77\xdf\x1e\xdf\xfe\xed\xf8\xc6\xd7\x67\xc7\xdf\xdd\xdd\xbf\x7d\x7b\xff\xfb\xdf\x8f\x6f\x7c\xf9\x44\xd3\xd1\x2b\x44\xca\x3e\x17\xee\xbc\x60\x8e\xef\xbc\xbd\xf7\xf9\x9b\x8f\x19\xc9\xa4\xe2\x85\x13\xd7\x12\x64\x73\x57\x75\xe8\xf3\x56\x8e\x6c\xd6\x57\x8d\x6c\x36\x3a\x17\xb2\x65\xe1\x27\x86\xec\x41\xc9\x24\xef\x78\xd8\x73\xd2\x83\x8e\x42\x9b\x96\x5e\x64\xc5\xe9\x71\xe5\x09\x6f\xe9\x54\x37\xcc\xc3\xc4\x94\xd1\x8f\xbf\x11\x30\x7c\xfd\x0b\x22\x09\x58\x5b\x00\xd9\xae\x3e\x60\x8e\x0f\x68\x63\x45\x49\xdf\x80\xf3\x68\x79\x65\x72\xf3\xe8\xf3\xae\xc4\x6b\x4d\x2b\x3c\xfc\x4c\x5b\x67\xfd\xe4\x44\x5b\x83\x1e\xaa\x44\xa0\x13\xb1\xf8\x9c\x19\x75\x41\xfd\x32\x02\x31\x78\x4d\x33\xa3\xed\x43\x6d\x11\x0d\x53\x7b\x3d\x22\xd0\x35\x3d\x86\x6c\xd4\xa7\xcc\xe5\x7d\xd3\xe7\x0e\xd1\x6a\xa6\x04\x22\x1c\xcf\x94\x51\x5b\x2a\x41\x59\xb7\xb6\xbc\xf8\x7c\xaa\x46\x3b\xa8\x96\xaa\xd9\x36\xc2\x92\xf4\xe0\x17\xdb\x76\xfc\x51\x42\x6e\x1b\xcd\x58\xf4\x04\x74\x90\x8d\xb0\x95\x7f\xad\x40\x7c\x10\xaa\x86\x77\x7f\xf8\x6c\xfc\xcd\xc7\xf7\xdf\xb9\x39\x7e\xf7\x16\xce\xce\x3f\x92\x82\x98\xfc\xce\xae\x61\x21\xf3\xbf\x0d\x1d\x2e\x20\x7e\x29\xaa\x8d\xe3\x85\xf2\xea\x67\x8f\x15\x27\x1e\xba\xdc\x89\x02\x60\xca\xec\x82\x7a\xc9\x07\x7d\xf9\xe2\x60\xdd\xad\xa5\x6f\x58\x17\x4d\x97\xc6\x5f\x4c\xb8\xc8\x46\xf1\x37\x13\x85\xe5\x17\x66\xcd\xcf\x6b\x8e\x34\x18\xef\x58\x64\x17\x7d\x5e\x38\xca\x99\xc2\xcb\xe8\x45\x93\x32\x06\xe2\xe5\x57\x2f\x6e\xe8\xa5\xbe\xe0\x38\x20\xe5\xcf\x61\x80\xd6\x5b\x78\x0e\x33\xc5\xd7\xb2\x07\x9a\x4a\xdf\x88\xce\x61\x2e\xc9\x10\x53\x86\xb2\xbb\x50\x41\x02\xd3\xe1\x4c\x72\x1f\x4c\xe2\xd3\x41\xc4\xe2\x37\x22\x01\x61\xa4\x0b\x16\xd9\xc6\x29\xb8\x65\xb0\x66\xcf\xe8\x9e\x1a\xb0\x1e\x19\xa2\x57\xf9\x36\xb0\x47\xc6\x25\xc3\x24\x89\xb4\x7e\x3a\x20\x8e\xc3\x23\xa6\x2c\x95\xdb\x9f\x86\xa5\xfa\x08\xeb\x68\x68\x0e\x8e\x46\xb6\x8e\xf2\xe6\x9d\x0f\x92\x87\x20\x84\x4e\x4d\xf3\xaa\x65\xb0\xe3\x53\x81\x97\x4b\xa4\x67\xe6\xf1\x8f\x37\x51\x28\x78\x87\xfa\x10\x3f\x1f\xc6\xa8\xc9\x12\x6c\x3a\x29\xba\xc9\x49\x25\xb2\xd1\xb3\xc3\x61\x7e\xc2\x38\x1a\x3d\x9b\xc8\xc8\x3e\x55\x8e\x87\x6a\xa9\x5c\x8e\xa5\x43\x64\xfe\x8d\xc5\xf9\x42\x32\x9f\xcd\xa3\x93\xa1\xb8\xb8\x3d\x5f\xca\xda\x99\xa1\xf4\xe8\xbb\x68\x68\x76\x8f\xcf\x63\xa8\x70\x84\x59\x34\x56\xcd\xcc\xc3\x0d\xba\xd0\x21\x91\xaf\x1e\x7e\x71\xa3\xb8\xea\x66\xb5\xf6\xb0\xba\xbb\xc1\xbb\x72\xaa\xdc\x92\xab\x64\x07\xd9\x88\x41\x1f\xbd\x7e\x71\xe3\x65\xa5\xc2\xcb\x70\x2d\x02\xa9\x6a\x93\xf2\xaa\x45\x4c\x1e\x02\xab\xe1\x2e\x28\x7c\x16\x61\x2b\xbe\x13\x28\x0b\x48\x60\xee\xb4\x0e\x13\x40\xdc\x81\x54\x44\x81\xe3\xc5\x0f\xa0\x76\xee\x4f\xad\x58\x9b\x75\x5a\x8c\x75\x62\x8d\x4d\xad\x81\x6c\x1b\xad\xa2\x67\x9e\x49\xcd\x2b\xa2\x22\xa9\xfb\xea\x4b\x4b\x45\xcd\x43\xc8\x9c\x78\x59\x22\x72\x3a\x87\x0c\x39\x93\xf0\x2a\xec\xa8\x89\x9d\xd1\x54\x4c\xb3\x3a\x9f\xc4\x2c\x61\x29\xa8\x75\xa6\x40\xf4\x88\x5f\x4b\x47\xce\xa2\x73\xe8\xff\xd1\xf2\xd2\xd2\xd2\x62\x19\x85\xa6\xe5\xa9\xc0\x5f\xfb\x6f\x00\x00\x00\xff\xff\x08\xf1\xcd\x87\x27\x2c\x00\x00")

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

	info := bindataFileInfo{name: "static/pages/writing.html", size: 11303, mode: os.FileMode(436), modTime: time.Unix(1727945567, 0)}
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

