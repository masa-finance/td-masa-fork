// Code generated for package internal by go-bindata DO NOT EDIT. (@generated)
// sources:
// _template/client.tmpl
// _template/handlers.tmpl
// _template/header.tmpl
// _template/main.tmpl
// _template/registry.tmpl
// _template/string.tmpl
// _template/utils.tmpl
// _template/zero.tmpl
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

var __templateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\xc1\x84\x92\x94\xe0\xdc\x03\x3d\xa5\x3d\x14\xda\x12\x4a\x7e\x40\x95\xd7\x89\x88\xbd\x12\xeb\x75\x92\x22\xf2\xef\x45\xb2\xd2\x93\x96\xa7\xd1\x68\x66\x63\x44\x4b\x9d\x63\x42\x6d\x7b\x47\xac\x35\xee\xf7\x2a\x46\x2c\xc2\xf9\x88\xed\x0b\x16\xcd\xde\xd8\xb3\x39\x52\xe1\x4a\x43\xe8\x8d\x12\xea\x13\x99\x96\xa4\xc6\x22\xdd\x54\x9b\x0d\xde\xf9\xe2\xcf\x24\xb0\x86\xe1\xf2\x0c\x31\x57\x7c\x1e\xf6\xe2\xd5\x43\x82\x85\x35\x7d\x3f\x36\x95\xfe\x06\xfa\x97\x3b\x56\x92\xce\x58\x42\xac\x50\xe8\xb7\xb9\x2e\xad\xde\x60\x3d\x2b\xdd\xb4\xd9\xcd\xe7\x1a\x8e\xc3\xa4\xf8\x71\xdc\xbc\xb1\xf5\x2d\xc9\x1a\x7e\xd2\x07\x7b\xa5\xcc\x56\x20\x11\x2f\xd5\x1c\x6b\x97\x7b\xc1\x0d\xa1\xa7\x21\x4d\x03\xe9\xc9\xb7\x23\x3a\x2f\x39\x90\xe3\x23\xba\x89\xad\x3a\xcf\x23\x3a\xf1\x03\x0e\x1f\x18\xed\x89\x06\x83\x8b\x33\x8f\xa4\x25\x77\xf1\x1b\x55\x26\xab\x39\x73\x6a\x56\x34\xe9\xcf\xe4\x85\x2f\xba\xce\xc2\xa5\x2b\x3d\x8b\x62\x85\xe7\xe2\x90\x9f\x92\x4e\xc2\x78\x9a\x51\x22\xd9\x6e\x5b\x16\x28\xeb\x0a\xb8\x27\xd3\x18\x41\xdc\xa6\x55\xff\x05\x00\x00\xff\xff\x88\xcc\x3e\xfd\xb3\x01\x00\x00")

func _templateClientTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateClientTmpl,
		"_template/client.tmpl",
	)
}

func _templateClientTmpl() (*asset, error) {
	bytes, err := _templateClientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/client.tmpl", size: 435, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateHandlersTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x4f\x6f\xea\x38\x10\x3f\xc7\x9f\x62\x8a\x38\x24\x15\x1b\xee\x54\x9c\x5a\xad\xca\xa5\x5b\x69\xb7\x27\x54\xad\xbc\x61\x52\x2c\x82\xc3\xda\xce\x96\x6e\xe4\xef\xfe\x34\xb6\x13\xe2\x04\x5e\x0f\xaf\x97\xc6\xf3\x7f\x7e\xfc\x66\xa6\x6d\x61\x87\xa5\x90\x08\xb3\x3d\x97\xbb\x0a\x95\x9e\x81\xb5\xac\x6d\x61\x7e\x3a\x7c\xc0\x6a\x0d\xf3\xfc\x95\x17\x07\xfe\x81\x41\x6e\xf0\x78\xaa\xb8\x21\x0f\xe4\x3b\x54\x33\x98\x93\x86\x99\xaf\x13\x42\x08\x02\x6b\x28\x1b\x59\xa4\x6f\xa7\x1d\x37\xf8\x58\x4b\x83\x67\xb3\x80\xf0\xac\xb8\xd6\x19\xa0\x52\xb5\x0a\x6e\x5e\xf1\x24\xf4\x89\x9b\x62\x8f\x0a\xb4\x51\x4d\x61\xa0\x65\x00\xd0\x05\xd5\x70\xe4\xa7\xad\x90\xe6\x3d\x08\x98\x65\x8c\xd2\xc0\x0b\x7e\x8e\x23\xa4\xd9\x34\xa8\x8f\xa6\xd0\x34\x4a\x4e\xb4\x5e\x39\x4c\xb7\x1a\xe7\x6b\xed\xc2\x19\x59\x66\xa3\xba\x43\x7f\x71\xd1\x85\x17\xe6\x41\xc9\x9c\xf0\x4d\x0f\xdb\xb8\xa7\xa7\x93\x3f\xee\xb9\x19\xc8\xe9\xd9\xc9\xa5\xc4\x2a\x56\x91\xc4\x69\x85\x14\x06\xfe\xa9\xeb\xaa\x07\x22\x6d\xe0\x3e\xaa\x29\x83\x8a\xff\xff\xb5\x91\xc2\xfc\xae\xea\xa3\x57\xe9\xb4\xf1\xff\x3b\x5b\x9d\x85\x9a\x45\x09\x4d\xee\xa2\x5e\xd0\xf0\x70\x85\xbe\xdd\xbf\x60\xb2\x06\xa3\x1a\x0c\x12\xdf\xd8\x1a\x8e\xfc\x80\x69\xd4\xdf\x02\x2a\x94\x5d\x46\x6f\x97\x65\x2c\x29\x6b\x05\x7f\x2f\xa0\x20\x2e\x10\xcb\x14\x97\x1f\x08\x91\x19\xb4\x2c\x49\x1a\x17\xa2\x3e\x90\x8d\x33\xce\x53\x17\x36\x63\x49\x22\x4a\xb8\xab\x0f\xce\x2c\x21\xb8\x85\x6c\x90\x25\x89\x25\x37\x1f\x62\x4b\xee\xf9\xe6\xe9\x1d\xd6\x40\x9f\x2c\xe9\x7b\xf0\x90\x8f\x2b\x26\x69\x5c\xb1\xb3\xcb\xb2\x8b\x97\xff\x41\xae\x38\x92\xe2\xba\xef\x37\xdd\xfa\x52\x2e\x90\xeb\x4f\x61\x8a\x3d\x14\x7b\x6e\x06\x6d\x13\xe1\xb2\x81\x55\xc1\x35\x82\x2b\x78\xd5\xcb\x06\xad\x6d\xc9\x3d\xb4\x5e\x74\x7c\x8a\xfc\xa8\xde\x2b\xae\xae\xbf\x9b\xde\x96\x20\x1c\xd0\x6d\x3c\x46\x19\x3c\xbb\x61\x49\x0b\x73\x1e\xcf\xc0\x02\xa6\xc4\x73\x5b\x20\x74\xd5\x90\xcf\x6a\x1d\x4f\xd5\xa5\xe1\x20\x58\x85\x36\xcc\xb9\x9b\xc6\x1e\x5f\x1f\xfe\x0a\x9d\x42\xda\x09\xc4\xde\xe0\x82\x6d\xdb\xfe\x16\x5c\xe7\xda\xad\xbe\x3f\xdd\x48\x6b\xb0\xb6\x6d\x69\x3c\xf0\x5f\x98\xeb\x7c\x23\x0d\xaa\x92\x17\x08\xb3\xc1\x4a\x73\x8b\x33\x46\x99\xd6\xa8\xce\x5f\xf8\x91\x96\x67\x8c\xb5\x28\xbb\x45\xd3\xd1\xbb\xc9\xbb\xcd\xb3\x8d\xfc\xfe\xfa\x3a\xe1\xe6\xe9\xfd\x01\x1c\xd7\x61\xf4\x47\xa0\xe5\x3f\x19\xf2\x6c\xe2\x41\x6d\x28\x45\x19\x43\xbe\x94\x62\x74\xe8\x65\x0f\x4e\x7b\xb7\x06\x29\xaa\x2b\xf9\x2e\x3b\x81\xec\x26\x6a\xcb\xe2\x17\x21\x8a\x72\xe7\x01\xf4\x1f\x31\x97\x92\x10\x4c\x0a\xb7\xc6\x7e\xfd\x17\x20\xec\xf0\x3f\x94\xc6\xc1\xb7\xa2\x35\x25\x8e\xaf\x0a\x4b\x71\xee\x41\x0d\x4e\x33\x6b\xd9\x72\x09\xb1\x87\xb5\xcf\xe1\x88\x09\x0d\x7c\xa2\x04\xf7\xe8\xa0\xcb\xfd\x1d\xb8\x15\xc1\x1d\x41\x62\xf5\xe8\x10\x06\xa2\xc6\xfc\xe8\x4f\xe2\x72\x09\x7f\xc8\x49\x5e\x8d\xb4\x24\xc6\xd2\xbe\x8e\xdb\x03\x39\x8d\x95\x76\x67\xfa\x46\xdd\xdd\x9a\xf9\x86\x91\xdd\x95\xbf\xdd\xe0\xf4\xe0\x4f\x2e\x4b\x4f\xc2\x01\x07\xf3\x74\x84\x4c\xd6\xdf\xdd\x6b\x7c\x62\x97\xcf\x1f\x01\x00\x00\xff\xff\x5a\x6e\x95\x1a\xd0\x08\x00\x00")

func _templateHandlersTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateHandlersTmpl,
		"_template/handlers.tmpl",
	)
}

func _templateHandlersTmpl() (*asset, error) {
	bytes, err := _templateHandlersTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/handlers.tmpl", size: 2256, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\xc1\x6a\xc3\x30\x0c\x86\xef\x7a\x0a\x61\x76\x48\x61\x73\x9e\x60\x97\xac\x3b\xec\xd2\x0e\xd6\xfb\x70\x62\xd9\x15\x69\xe4\xe0\x28\x63\xc3\xe4\xdd\x47\xda\xd0\x5e\x04\xd2\x0f\x9f\xbe\xbf\x14\xf4\x14\x58\x08\xcd\x99\x9c\xa7\x6c\x70\x59\xa0\xae\xf1\x2d\x79\xc2\x48\x42\xd9\x29\x79\x6c\xff\x30\x26\xf5\x91\xe4\x19\xf7\x47\x3c\x1c\x4f\xf8\xbe\xff\x38\x59\x80\xd1\x75\xbd\x8b\x84\xa5\xe0\x93\xfd\xdc\x96\x65\x01\xe0\x61\x4c\x59\xb1\x02\xd3\x25\x51\xfa\x55\x03\x26\x0c\xeb\x9c\x34\xb3\xc4\xc9\x00\x98\xc8\x7a\x9e\x5b\xdb\xa5\xa1\x5e\xf9\xb5\xfa\xba\x65\x31\xb0\x83\x55\xe2\x90\x5e\xd2\x78\x13\x64\xe5\x24\x18\x52\xc6\x9e\x68\x64\x89\x78\xe3\x4f\x16\x7e\x5c\xc6\x6f\x7c\xc5\x96\xc5\x36\x73\x08\x94\xcb\x72\x3f\x6e\xbf\x6d\xe3\xba\x3e\xe6\x34\x8b\xaf\x76\xf7\x30\x0c\x6a\xbf\xae\x32\x94\x2b\xe1\xcb\x23\xd9\x14\x6d\x33\xf3\xc5\x5f\x81\x50\x0a\x92\xf8\xb5\xda\x7f\x00\x00\x00\xff\xff\xaa\xa0\x2f\x0c\x35\x01\x00\x00")

func _templateHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateHeaderTmpl,
		"_template/header.tmpl",
	)
}

func _templateHeaderTmpl() (*asset, error) {
	bytes, err := _templateHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/header.tmpl", size: 309, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x5f\x73\xdb\xb8\x11\x7f\xd7\xa7\xd8\x78\x72\x29\xe9\x53\xa8\x7b\x76\xce\x0f\xb5\x9d\x6b\x3c\x93\x5e\x33\x71\xd2\x87\x76\x3a\x1e\x88\x5c\xca\x18\x51\x80\x0a\x40\x92\x73\x3a\x7d\xf7\x0e\xfe\x91\x20\x09\x4a\xb2\xab\xa4\xb9\xfa\x49\x06\x81\xc5\xee\x6f\xff\x03\xd8\x6e\x5f\xc3\xe4\x7c\xc6\xd5\x97\x25\x5e\xc0\x8c\xaa\x87\xd5\x34\xcb\xf9\x62\x32\xe3\xaa\x98\xa8\x62\x42\x99\x42\xc1\x48\x35\x99\x21\xcb\x72\xce\x4a\x3a\x3b\x9f\xc0\xeb\xdd\x6e\x34\xda\x6e\xa1\xc0\x92\x32\x84\xb3\x05\xa1\xec\x0c\x76\x3b\x3d\xf6\x72\x39\x9f\xc1\xc5\x25\xbc\xcc\x3e\x90\x7c\x4e\x66\xe8\xc6\x15\x2e\x96\x15\x51\x08\x67\x0f\x48\x0a\x14\x67\xf0\x12\x1c\x19\x41\xd8\x0c\xe1\xa5\x34\xab\xee\x94\x58\xe5\x4a\xea\x6f\x93\x09\x68\x82\x32\xbb\xe6\x8b\x05\x32\x65\x29\xbd\x06\x5a\xea\xc1\x1b\x9e\x4b\x3f\xe2\x28\x14\x9a\x42\xf3\xc5\xae\x57\x82\x2e\xe0\x65\xe1\xa7\x22\x8b\xfd\xb4\x24\xdf\x53\x36\x37\x2b\x01\x00\xf4\x70\xc3\xf4\x52\x50\xa6\xee\x2b\x3d\xe1\xac\x35\x35\x4a\xe8\xf3\xc7\xf7\x96\x01\xcd\xc3\x1d\xa2\x93\xc3\x0e\x43\xc9\x05\x08\x2c\x51\x20\xcb\x31\x0b\x29\x68\x3d\xb8\xb9\xbf\x92\x85\x86\x0e\xa4\xc1\x03\xb6\xa1\x9c\x25\x58\x41\x7f\xa1\x58\x15\x21\x54\x65\x0c\xaa\xd2\x32\x7b\x84\x58\xe5\x3e\xb1\x08\x2b\x20\x31\x3b\xb0\x82\x2a\xca\x19\xa9\x52\x48\x18\x57\xd0\x1e\xbc\xe2\xbc\x4a\x1b\xe9\x3f\x4b\x84\x3b\x54\x96\x3b\x2f\x94\xa6\xf5\x97\xee\xe0\x03\x56\x4b\x14\xb2\x05\x88\xe5\x38\x9c\xd5\x32\x25\xcb\xbe\x46\x4d\x73\xdf\xe1\xdb\x18\x97\x61\xdd\x30\x29\xb3\xbf\x63\xae\xb8\x68\x59\x96\xa3\xfa\xe9\xcb\x12\x6f\x6f\x80\x4a\xf8\xf4\x1e\x8c\x12\x68\x01\xbc\x6c\x4f\xca\x46\x39\x67\x52\x45\x57\x5e\xc2\x4f\x8f\x76\xfc\x1d\x3e\xde\xde\xc0\x6e\xf7\x26\x64\x65\xb4\xdd\x36\x4c\xff\x86\x82\xdf\x17\x28\xe8\x5a\x73\xed\xd0\x6e\x3e\x4b\x25\x28\x9b\x75\x26\x68\x8e\xdf\xb2\x9c\x17\x08\x74\xb1\xac\x50\xeb\x58\xc2\x94\xb2\xcc\x8e\x8a\x6c\x54\xae\x58\x0e\x89\x65\xe2\x23\xe6\x48\xd7\xa8\x65\x85\xf3\x16\xbf\xa9\x23\xe3\x26\x5e\xad\xca\x3f\x8b\x99\x99\xa6\x89\x5d\xad\xca\x12\x45\x0a\x28\x04\x17\xb0\x35\xf0\x53\x0f\x43\x48\xf4\xf2\x12\x18\xad\xdc\x0c\xfd\x27\x50\xad\x04\x83\x72\xa1\xb2\xb7\x7a\x71\x99\x9c\xe5\x84\xfd\x49\x01\x5a\xae\x1d\x09\xb2\xd1\x80\x19\x13\x90\x9a\xc2\x59\x6a\x28\xec\x06\x35\xe5\x0d\x20\x60\x35\xfb\xb0\x52\xb7\x37\x49\x44\x0d\x69\xd7\x6c\x07\xfd\xc5\x7b\x82\xf5\x8f\xc0\x7a\xfd\x47\x5a\xc2\x8b\x3e\x98\x59\xa8\xa8\x9c\x2f\x96\x44\xe0\xbd\xd6\xa7\x36\x3f\x0d\x6e\x03\x48\x6c\x6d\x67\x2f\xc3\x8e\xfe\x72\x87\x2a\xe9\x7d\xbd\x65\x05\x3e\x6a\x8d\x39\x84\x3c\xcb\xf1\x00\x76\x8c\xa4\x71\x5f\x3d\x16\x8e\xa7\x08\xf4\x8e\xc8\x3d\x02\x39\x94\x7a\x6e\xee\xf7\x3f\xa8\x7d\x3b\xe1\x9d\xc9\x23\x49\x85\x2c\xaa\xa8\x20\x68\xa4\x16\x43\x1d\x77\xb7\x5b\xb7\x89\xf3\x1b\xd8\xed\x68\xa1\x5d\x17\x2b\xa9\xa7\xde\xeb\x9f\x86\xa9\x71\xc0\xd1\x0d\x5f\x4d\x2b\xac\xf9\x12\x7c\xb3\xdd\xea\x05\xbb\xdd\x7a\xbb\x45\x56\xec\x76\x1a\x77\xab\x83\x03\xbc\x04\xc2\x47\x69\x77\x0c\x68\xbf\xdc\x82\x6f\x9c\x6c\x5e\xbe\xfb\x31\xac\x1b\x5e\x04\xdf\x1c\x02\xbb\xc1\x21\xd8\xd9\x7f\xbc\xd5\xe9\xbf\x24\x39\x86\x9f\x9d\x3d\xac\xfb\x41\x60\x20\x10\xac\x18\x99\x56\x08\x8a\x0f\x05\x83\x0b\x28\x8d\xe5\x58\xa0\x3e\x92\x8d\xc7\x0a\x6d\xa4\x83\x0d\x55\x0f\x40\x8d\xfd\xfc\x50\xe8\x30\xad\x43\xc7\x18\x68\xf1\x98\xb6\x36\x6f\xcb\x60\x14\x33\x0a\x58\x46\x21\x34\x36\xeb\x2c\x1e\x03\xd3\x37\x66\xc6\x8b\xff\x8d\x58\x17\xf0\xc3\xc6\xca\x34\xd6\x6c\x34\x82\x05\xae\x6f\x6d\x74\x9f\x89\xd8\xbd\x7e\xd1\xb9\x60\xb7\x4b\xd6\xe9\x7e\xe5\x0f\xd9\x5e\x2f\xda\x0c\xf3\x31\x68\x49\x7b\xac\x68\x4f\x34\xf1\x10\x1d\x93\x60\x9e\xaf\x00\xda\x4e\x3e\x7d\x83\x69\x8c\xe5\x00\xa3\x4f\x34\xa5\x53\x4a\xe1\x0c\xa6\xb6\x95\x21\xfd\x1c\xb0\x91\x43\xb1\x73\xc8\x82\xa2\x06\x15\x49\x1d\x83\x2b\xfa\x89\x2c\x40\x88\xd1\xca\x17\x70\x83\xa9\xad\x8e\xe5\x9d\x5d\x4d\xcd\xdd\x29\x2e\x25\x2a\x09\x6b\x52\xad\xd0\xd7\x75\xcd\xb7\x3c\x58\x6e\xa0\x3e\xbe\x98\xea\xee\x93\xd8\x2d\xf6\x17\xa9\xa9\xab\xe6\x7b\xcc\x87\xe9\x58\x47\x58\x4b\xeb\x6b\x95\x14\xfb\x48\xd6\x6e\x08\x4a\xac\xd0\xea\xd1\x9a\xd5\xf3\xd8\xf9\xcc\xe4\xa9\x18\x2a\x49\x25\x31\xa8\x1a\xfb\xc6\x7e\x32\x84\x0e\x33\x63\x54\xb4\xa7\xe9\x88\xaa\x77\x32\xe9\x37\x3f\xd6\xec\x8f\x37\x51\xdd\x42\x69\x42\x53\xce\x2b\x24\x0c\x36\x0f\x34\x7f\xd0\x71\x4d\xeb\x4b\x6f\x6f\x67\x6d\x88\xd4\xa6\x7f\xbc\x41\x77\x19\x4b\x52\x38\xca\xa6\xc7\xc0\xe7\x86\x99\xb4\x69\x1b\x5e\x9c\xbc\x64\x0c\x02\x84\x61\x6a\xdc\xb2\x86\xe0\xeb\x01\xcd\x8d\xad\x59\x3b\x65\x35\x4d\x5a\x2c\x28\x69\x94\x6f\x30\xd6\x7a\xd9\xd1\x27\xb4\x5e\x76\xc1\x37\x6b\xbd\x0a\x8c\xb7\x5e\x8a\x1f\xdd\x7a\x75\xf3\x60\x93\x47\xae\x39\x93\xab\x05\x0e\x34\x62\xcf\x49\x7f\x03\xec\xf6\xd3\xdc\x7f\xdd\xe4\x85\x71\xf6\x38\x3f\x3f\x9d\x25\xef\x29\xe6\x22\x9c\xf6\xcb\xee\xaf\xe2\x53\x9e\x9f\xce\x76\xdb\x6e\x39\x1d\xe7\xb7\x5f\x3f\xea\x3f\x7b\xcc\xf7\x1e\xd9\x78\xc8\x84\x5a\xed\x4c\xda\x15\x74\xb0\x14\x7f\xbe\x21\x1d\x57\x47\xf9\xbf\xb6\x38\xba\xb9\xa2\xc5\xa3\x96\xe3\xa7\x37\xe6\xd7\xcf\x8d\x88\x66\xe0\xc7\x1f\x61\x3b\x6a\xe3\x77\xb0\xd0\xae\xe5\x65\xec\xd9\x58\x1d\x83\xd7\xb7\xc2\xac\x8f\x9b\xfe\x5b\x13\x61\x5a\xd1\x7f\xfe\xcb\xd2\x71\xc4\x7b\xf3\x0c\xc6\x1a\x09\x6b\xa3\x0e\xea\x66\xe0\xe7\x1a\xa7\x37\xf5\x2f\x83\x7a\x17\xf4\x8e\xa9\xc2\x71\x4d\xad\x65\xd5\xa4\x15\xa7\x05\x1b\xb0\x2d\xd3\xf5\xb2\x76\xe1\xdc\x94\xfc\x7f\x20\xa5\xd4\xf1\x67\xf0\x0c\x20\x54\x5d\x9d\xff\xf7\xa8\x2e\xe8\xae\xf5\xec\x2c\x9e\xea\xf6\x75\xd8\xdf\x07\x22\x07\x4d\xa2\xeb\x98\xed\x66\xea\x0f\x67\x04\x7d\x47\x81\x27\x05\x2e\xc3\x3e\xdf\xc0\x25\x90\xe5\x12\x59\x91\x08\xbe\x19\x5b\xc8\x8e\x89\x0c\x87\x33\xaf\x23\x7b\xb0\xa0\x13\x7c\xd3\xdf\x30\x96\x67\x4f\xbf\xf5\x80\xb4\x03\xf8\xee\x7a\x59\xf7\x89\x67\x6f\xa7\x0a\x51\xdf\x5f\x8e\x3d\xb6\xdf\xea\x01\xb8\x37\x8c\x1d\x7f\x96\xf3\xbc\xa0\xf5\x6d\x71\x1a\x32\xe9\xe7\x47\xa9\xff\x13\x3b\xe8\x57\xa7\xf1\xb3\xcb\xa3\x4e\x9d\xdc\x3d\x6e\xcb\x07\x27\x13\x30\x17\x80\xe6\x3e\x36\x68\x05\xeb\x41\x2e\x9a\x3b\xc3\x70\xe5\x70\x83\xd8\xe9\x0f\x6b\x4a\x49\x1a\x21\x03\x5b\xcf\xe6\xab\x3e\x25\x73\x1a\x16\x74\xac\x6f\x99\x5c\x09\xca\x66\x40\x3d\x05\x09\x94\x41\xce\x17\x4b\x5a\xe1\x6b\x45\x17\xe8\xef\x43\xc2\x1b\x4e\x9d\xee\x93\xd1\x7d\x78\xb3\x08\x97\x7e\x3f\x37\x6b\xbb\x73\x13\x5c\xff\x1b\x9b\x10\x47\xf0\x3e\x26\x55\x7c\xb5\x13\x25\x1d\x05\xb7\xea\x7f\x45\xf5\xc0\x8b\xd6\xfd\x6d\x3d\x04\x94\xad\xf9\x1c\x25\x2c\xec\x48\xaf\xe9\xb5\xd0\x69\x44\x6c\x87\x6d\x6e\xb4\xbf\x64\x07\x1e\x14\xc0\xf7\xf2\xa2\xc0\x78\xde\x5e\x4a\x46\x2c\x4b\xaa\x99\x7c\xd2\xd7\x09\x0d\x85\x2b\xae\xae\x09\xfb\x2c\xbd\x5f\x5c\x13\x06\x53\x84\x95\xc4\x02\xa6\x5f\x60\xca\x95\x1c\x58\xf9\x11\xe5\xaa\x52\xb1\xb1\x3b\xca\x66\xab\x8a\x98\xf0\x6d\xfd\x25\x87\xf3\xeb\x8a\x22\x53\x69\x57\xdb\x49\xeb\x30\x4a\xe0\xbf\x57\x28\xd5\xfd\x92\x08\xb2\x90\xee\xca\x3c\x35\xfe\xd6\x9c\x6b\xd8\x4d\xea\x42\xe6\xdc\xbb\x90\x63\xa7\xb9\x01\xb4\x4d\x8a\xff\xe4\x42\x66\x70\x2b\x68\x70\xf6\xcd\xb3\xe9\x6c\x2c\x8d\x0e\x41\x7f\x80\x18\xa8\x89\xe4\xf3\x7b\xc7\xab\xbf\xd7\x6f\xe7\xa7\x3c\x13\xcb\x3c\xbb\x35\x96\xfc\x91\x6c\x92\x5c\x3d\x8e\xc1\xad\x18\xc3\x2b\xbb\xd1\x81\xb3\x15\x46\x2b\xc3\x63\x27\xec\x85\x48\xb7\xab\x39\xb7\xce\x12\xcf\xde\x56\xb8\x90\x63\x13\x08\x63\xd9\xc6\x07\x20\x3b\xbb\x33\xcf\x9f\x80\x86\x8b\x4e\xa5\x49\xad\x48\x0c\x00\x3e\xbb\xe2\xbc\xba\xae\x88\x94\x67\x46\x3b\x71\x5d\x4e\x39\xaf\x9e\xa8\x3a\xa7\xf0\x2b\xfe\xf8\xcd\x34\xe8\x0f\xe0\x06\xe4\x83\xdd\xae\xa3\xdb\x00\x60\x9f\x9c\x49\x25\xb1\xf9\x16\x4b\x7c\xfb\x77\x88\x98\x42\x88\xca\x15\x91\x58\xd7\xbb\x43\xc6\x71\x6f\x0e\x81\x2f\x2e\x0f\x10\xc8\x92\x73\xbd\xf5\x27\xe1\x8b\x66\xb7\x2d\x9f\xef\xb3\xa7\x20\xc8\x9e\xd2\xb4\xc2\xe3\x56\x6d\x11\x7c\x0e\x7f\x9b\x7f\x05\xcd\xf3\xf9\x01\xbf\x6d\xbb\x6c\xab\x1e\x19\x4a\x32\xc1\xc1\x75\x7d\x08\x6a\xde\xd0\xdd\x36\x39\x3f\x7c\x1b\xd6\xdc\x35\x2c\x05\x4a\x53\xbb\xd4\xf5\x99\x4f\x93\x33\x64\x28\x68\x6e\x1e\x41\x65\x3a\x45\x34\x0d\x49\x9d\x36\xea\x9c\x51\x0e\xe4\x0c\xb7\xcc\xb1\xa7\xab\x91\x47\xa2\xeb\xa5\x0b\xfd\x1b\x66\xd1\xde\xc5\xd7\xa6\xd3\x55\x99\x9a\x69\xdd\xb2\x54\x8f\xe9\xbf\x25\x61\x34\x4f\x4c\x05\xa9\x87\x0c\x7d\x90\x1b\xaa\xf2\x07\xfb\xf4\x61\x96\x25\x9a\x7b\x7b\xcb\xe6\x81\xc9\x4d\x26\x37\x07\xa1\xbe\x58\x93\xe6\x1d\xa3\x5e\x9d\x13\x89\xf6\xb0\x3e\xcf\xea\x42\xd5\xc1\x96\x07\xe8\x78\xcc\xfd\xba\x02\x4b\xb2\xaa\xd4\x85\xe3\x69\xed\x39\xaa\x9f\xf1\x35\x98\xd7\x65\x98\xd3\x7c\x50\x60\xd5\xff\xbb\x7a\xca\xfc\xdf\x2d\x05\x6b\x4a\xf6\x84\x51\x17\xe3\x77\xe6\xa1\x98\x9b\xff\x0f\x14\x3c\x49\xcd\x0d\xcc\x28\xbc\xb3\x68\x81\xdb\xb9\xc0\x20\xe2\x0b\x14\xf8\x5a\xa2\xa0\xa4\xa2\xbf\x11\x45\x39\xab\x8b\xc2\xa6\x2b\xb2\x8e\x16\x23\xa7\x95\xd5\xbe\xbd\x48\x3a\xad\x71\x2b\xe2\xd2\xa2\xd6\xfc\x74\x55\x66\x1f\x10\xe7\xb7\x37\xae\x19\x19\x6a\x42\x86\xb3\x9a\x53\x39\x2d\x82\x17\x2e\xfb\x94\xed\x3c\xd6\xe8\xba\xa5\x6a\x7b\x65\x71\x51\x6f\xe9\xb1\xd3\xb5\x62\xd7\x02\xb2\x7a\xd6\xda\xb7\x59\x79\x50\xb8\x76\x5a\x2a\xf3\xd6\xc4\x35\x95\xda\xae\x0f\xbf\x2f\x31\x72\x1e\xec\xb4\xca\x6c\x4f\x3b\xb5\xeb\x62\xf7\x6a\x1d\x8d\xab\xfa\x5f\x6f\xc0\x51\xb8\x9f\xcc\x86\x36\x84\x5f\x71\xf3\x99\xe1\xe3\x12\x73\x85\xc5\xed\x4d\x42\x8b\xd4\x5f\xdb\x58\xb3\x6c\x5b\xd0\x94\x3f\xa2\x04\xf5\xd0\x75\x97\xa5\xe0\x6b\x6a\x34\x40\xdc\xbb\xd0\x2c\xf4\xaa\x26\x49\x37\xef\x63\xa1\x7e\x2a\x1a\xe4\x9a\x8e\xef\x1c\x71\x9b\x17\x38\x40\xb3\x8b\xef\xe0\xa6\x36\x4a\x84\x9f\xea\x5b\xbd\x9e\x2f\x74\x6e\xf2\xa6\x4f\x7c\xd3\xd2\x02\x3a\x10\xb8\x73\x7b\x67\x6c\xf1\x70\x44\x3d\xc2\xc7\xf6\xb2\xa1\x15\x55\xd8\xb6\x3b\xfe\xee\x64\x9a\xf5\xc1\xbf\x84\x75\xa4\xbb\x3e\xf0\x94\xf5\xe9\x0a\x70\x4f\x70\x8e\x56\xc0\xef\xbf\xc7\xb9\x7d\xee\x63\x9d\xe0\x5d\x73\xff\x5d\x91\x23\x12\xd9\x30\x6b\xf8\x4e\x47\xdd\x7c\xee\x7e\xfe\x27\x00\x00\xff\xff\xfa\x7e\x90\x24\xa1\x2f\x00\x00")

func _templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateMainTmpl,
		"_template/main.tmpl",
	)
}

func _templateMainTmpl() (*asset, error) {
	bytes, err := _templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/main.tmpl", size: 12193, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateRegistryTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6b\xe3\x30\x10\x85\xef\xfe\x15\x0f\x63\x96\x04\x76\xed\x65\xf7\x16\xc8\xa9\x3d\xb4\x90\xb4\x25\xe4\x56\x7a\x50\xec\xb1\xa3\x26\x96\x85\x24\xb7\x31\x42\xff\xbd\x4c\xed\x26\x31\xf5\xa1\x37\x7b\xde\xcc\xd3\x7c\x8f\xf1\x1e\x05\x95\x52\x11\x62\x43\x95\xb4\xce\x74\x31\x42\x88\xbc\x47\xa2\x0f\x15\x16\x4b\x24\xe9\x93\xc8\x0f\xa2\xa2\xa1\xee\xa8\xd6\x47\xe1\x08\xf1\x9e\x44\x41\x26\x46\xc2\x4a\xe4\xfd\x1f\xc8\x12\x49\xba\x12\x1d\x19\x2e\x65\x19\xfa\xef\x37\x32\x56\x36\x0a\x4d\x09\x9b\xef\xa9\x16\x69\x94\x37\xca\xba\x41\x5e\x82\x9f\xbb\xcc\xb1\x13\xa9\xe2\xd3\x35\xcb\xb0\xed\x34\x59\x18\x72\xad\x51\x16\xb5\xd0\x5a\xaa\x0a\xa5\x69\x6a\xb8\x4e\x13\x64\x61\xe1\x1a\x6c\x57\xfd\xaf\x12\x35\xd9\x34\x2a\x5b\x95\xf7\xa3\x6b\xa1\x67\x73\x9e\x7b\x6e\xa5\x72\xff\xff\xbd\x58\x67\xd8\xc1\x47\x18\x5c\x27\x45\xde\xc2\x08\x55\x11\x12\x3a\x52\xdd\x47\xb1\x19\x42\xe2\xdd\x00\xe0\xef\x89\x57\x67\x3d\xbd\xa3\xd3\xfd\x2d\x42\x58\x20\x3e\xd7\x36\xe2\x1d\x21\xc4\xbf\xaf\x91\x80\x10\x5d\x81\xdd\x70\x10\xa6\xcd\x5d\x63\xd6\x42\xf3\x26\x76\x84\x95\x5f\xf4\x11\xd5\x78\x6e\x0c\xc8\x5d\xb3\x39\x76\x52\xa5\x8f\xbb\x57\xca\x1d\x7c\xf4\x9d\x74\xaa\xeb\x07\xd0\x93\xc8\x13\x5e\x5f\xd9\xfe\x3a\xb7\x3f\x88\x9a\x8f\xc8\x07\x84\xa9\x44\xf8\xb8\x86\xca\x47\x00\x00\x00\xff\xff\x82\x8a\xe4\xa1\x97\x02\x00\x00")

func _templateRegistryTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateRegistryTmpl,
		"_template/registry.tmpl",
	)
}

func _templateRegistryTmpl() (*asset, error) {
	bytes, err := _templateRegistryTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/registry.tmpl", size: 663, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateStringTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6f\xdb\x30\x0c\xbd\xe7\x57\x10\x86\x81\x58\x5b\xa7\xde\x0b\xe4\x92\xad\xc3\x76\xd9\x61\x03\xb6\xc3\x5a\x14\x4e\x44\x17\x02\x64\xaa\x90\x14\x77\x83\xa0\xff\x3e\x48\xb6\x13\x7f\x25\xed\x9a\x43\x00\x8b\x8f\x7c\xef\x91\x94\xbc\x17\x58\x49\x42\xc8\xac\x33\x92\x1e\x1f\x04\x1a\xd9\x60\x06\x21\x78\x0f\xb9\x85\x9b\x0d\xe4\x10\xc2\xea\xfa\x1a\x7e\x24\x04\xc8\xfa\x49\x61\x8d\xe4\x2c\x54\xb5\xe3\xed\x29\x1a\xbe\xaa\x0e\xb4\x87\x22\xa5\xf1\xef\xb8\x47\xd9\xa0\x81\x10\xe0\x5d\x7b\xf4\xad\xac\x11\x42\x60\x5d\x9d\x82\x41\x4b\x09\x7e\x05\x00\x20\x2b\x98\xa7\x6e\x36\x40\x52\x75\x88\xf8\x33\xe8\x0e\x86\x20\x8b\xd0\xbe\x62\x41\x52\xb1\x2c\x41\x42\xfa\x6f\x4a\x03\x76\xd7\x95\xb7\x7c\x7b\x90\x4a\xa0\x49\x21\xbb\xe3\xbf\x8c\x74\xd8\x69\x18\xd5\xc9\xd8\xca\xfb\x0f\x51\x48\x6e\xf9\x4f\xdc\x3b\x1d\x35\x8c\xd2\xb6\x7f\x1d\x16\xeb\xdf\x6b\x96\x4e\x2b\x6d\xe0\xe1\x0a\x30\x36\xc9\x94\xf4\x88\x73\x07\xfc\x56\x61\x6d\x07\x06\x26\x02\x52\x07\x9f\x8c\x24\x57\x20\x83\xf7\x90\x5d\xdd\x51\xc6\x06\x5e\xc6\xcc\xf7\xeb\x56\x23\x2a\x8b\x53\x6d\x47\x4b\xa9\x42\x44\xb5\x9a\xf2\x2a\x0d\xd1\xf2\xcf\x12\x95\xb0\x7d\x5a\x67\x95\xb4\x83\xbc\xe2\x1f\x35\x09\xe9\xa4\xa6\x52\x6d\xb5\x56\x13\xd0\x18\xd0\x07\x97\x26\xc6\xe3\xc9\x08\x9d\x58\x63\xe4\x4b\x69\x8b\x59\xf4\x2b\x09\xfc\x93\xb6\xc2\x1f\x09\x91\xc4\x9c\xff\x3f\xe6\x71\x4a\xfa\xa4\x0f\x3b\x85\xc7\x54\xa3\x9f\xbd\x8f\xad\x0b\xa1\xf1\x1e\x49\x84\x70\x69\x74\xad\xd8\x6e\x39\x06\xfa\x16\x6b\xcf\x06\x3c\x96\x36\x90\xd7\x9c\x38\x8d\x7e\x5e\xf6\x7d\x79\x53\x1a\xc6\x5e\x29\xe6\x9c\xac\xfb\x35\x5b\xa2\x3d\xb7\x72\x47\xe8\x85\xb5\xbb\x73\xa3\x7e\xdd\x40\x36\x15\x79\x4b\x7b\x2d\x70\x36\xc3\xb9\xbf\x17\x46\xc1\x5e\x25\xe8\x6d\xf5\x4e\xbd\x98\xfa\x3b\x5d\xcc\xb3\x3b\xba\x70\x47\xc2\xb9\x8c\xee\xf3\x05\xd2\xfe\x4d\x1a\x40\xba\x17\xd0\xee\x78\xff\x8e\xae\x46\x65\xfe\x05\x00\x00\xff\xff\x66\x7b\xb0\x31\xd2\x05\x00\x00")

func _templateStringTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateStringTmpl,
		"_template/string.tmpl",
	)
}

func _templateStringTmpl() (*asset, error) {
	bytes, err := _templateStringTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/string.tmpl", size: 1490, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateUtilsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\xcf\x6a\xe3\x30\x10\xc6\xef\x7e\x8a\x0f\x63\x16\x76\xf1\xda\xf4\x6a\xc8\x29\xa1\xa7\x52\x02\x6d\x4f\x21\x04\xc7\x1e\x17\x11\x45\x76\x25\x85\x26\x08\xbd\x7b\x91\x6c\x27\x8a\x71\x68\x2e\xd1\xcc\x7c\xf3\xe7\x37\x1e\x63\x50\x53\xc3\x04\x21\x96\xf4\x75\x22\xa5\x77\x5d\x29\xcb\xa3\x8a\x61\x6d\xa5\xcf\xa8\x5a\xa1\xe9\xac\xb3\x65\xff\x6f\xcc\x7f\xb0\x06\xd9\x87\xe8\xca\xea\xb0\x76\x52\xd2\x24\x15\xac\x1d\x43\xcf\x8c\x78\xed\x1c\x69\xe4\x5c\xb2\x14\x9f\x84\xa4\x41\xb1\x08\x62\xc6\x80\xb7\xdf\x24\x91\x34\xd9\x6b\x79\x24\x58\x0b\x63\xa0\xe9\xd8\xf1\x52\x13\xe2\x4e\x32\xa1\x77\xfa\xd2\x51\xec\x92\xad\x4d\x5d\x31\x12\xb5\xb5\xd1\xf0\xc2\xf8\xe4\xca\xe5\xa7\x18\x08\xf0\xcf\x18\x8c\x55\xa7\xda\xfe\x19\x05\xdc\x0e\x64\x37\xa4\xc6\x01\xc7\x0c\x62\x04\xe0\xda\xa5\x58\xe0\x8f\x31\x48\xae\x9d\x7c\xd4\xfd\xa6\xd8\xc9\x8d\x3b\xd0\x04\xe8\x05\x66\xf6\x91\xde\xd5\x1b\x06\x77\x66\x48\xf2\x88\x29\xd8\xde\x40\xce\x1a\x24\xd9\xaa\x3d\xed\x39\xbd\x71\x56\xb9\x0e\x9b\xed\x66\xeb\x09\xde\x2f\x1d\xdd\x6d\xd3\x8b\x6f\xb2\x79\x91\xff\x8a\x13\xff\xef\x43\x71\x26\x0e\xfe\xba\xa2\x3c\x8f\xf2\x1c\x2f\xce\x2e\xc2\x53\x61\x29\x12\xa7\xf2\x17\xd3\x0b\xfd\xbe\xca\xba\x46\xc2\xf0\x04\x6b\xff\xfa\xfd\x79\xd1\xcc\x3d\x3c\x6a\x4d\x52\xb6\x32\xec\xbd\x6e\x95\x62\x7b\x4e\xe8\x23\x77\x53\x90\x94\xd3\x01\x9c\x2f\x5b\xb6\xf5\x78\xad\xde\x1e\xf0\x8b\xab\x63\x45\xaa\x92\xac\xd3\xac\x15\x8f\x87\xfb\x09\x00\x00\xff\xff\x20\xf7\xdd\x03\x78\x03\x00\x00")

func _templateUtilsTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateUtilsTmpl,
		"_template/utils.tmpl",
	)
}

func _templateUtilsTmpl() (*asset, error) {
	bytes, err := _templateUtilsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/utils.tmpl", size: 888, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateZeroTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xcd\x6a\xdc\x30\x10\xbe\xfb\x29\xa6\xc2\x07\x7b\xc1\x72\xbb\x90\x50\x02\x7b\x4a\x28\xf4\xd2\x96\xa6\xa7\x5e\x82\x6c\x8f\xbc\x2a\x5a\x69\x2b\x8d\x43\x53\xa1\x77\x2f\xb2\x37\xb1\x93\xdd\x26\xac\x4f\xd2\x68\xbe\xf1\xf7\x23\x85\x00\x1d\x4a\x65\x10\xd8\x5f\x74\xf6\xae\x43\xa7\xee\x91\x41\x8c\x21\x40\xee\xe1\x6a\x03\x1c\x62\xcc\x42\xa8\xa0\x5e\xf5\x96\x1e\xf6\x78\x05\xbd\xa2\xed\xd0\xf0\xd6\xee\xea\xde\x52\x57\x53\x57\x2b\x43\xe8\x8c\xd0\x75\x8f\x86\x7b\x72\x43\x4b\x37\x28\x57\x35\x54\x31\x66\x72\x30\x2d\x14\xe3\x44\xfe\x1d\x5b\x54\xf7\xe8\x20\x46\x58\x4d\xa5\x2f\x62\x87\x10\x63\x09\x3f\xd1\xd9\xa2\x84\xc6\x5a\x0d\x21\x03\x00\x50\x12\x8e\x61\x9b\x0d\x18\xf5\xd8\x91\x3e\x87\x34\x38\x03\xe4\x06\x1c\x6b\xe7\xf2\x95\x0a\x75\x37\xd3\x4d\x60\x27\x4c\x8f\x90\xcb\xe4\x40\xee\xf9\xa7\xd4\xe1\x47\x5b\xaa\x44\xca\x20\x14\xb9\x1c\x99\x97\xe3\xea\xda\x9a\x4e\x91\xb2\x46\xe8\xb1\xb7\x4c\xae\x1d\x14\xbc\x3b\x96\xce\x43\x20\xdc\xed\xb5\x20\x04\xd6\xda\xdd\x5e\x38\xbc\x4b\x09\x30\xc8\x65\xb2\xe2\x48\x9c\x14\xda\x2f\xd5\xa1\xe9\x0e\x74\xa6\x55\x96\xbd\x74\x62\xee\x4b\xa2\xb2\x45\xd4\xcf\x7f\x78\x76\xbc\x27\xec\x9a\x32\x7c\xdc\x29\x09\xd6\x41\xc1\x6f\xb5\x6a\x93\x3f\xfc\xc6\x0e\x8d\xc6\xc3\xb6\x7a\x8a\x70\x22\xa8\x3d\x26\x04\xfe\x86\x82\xff\x78\xd8\x27\x00\x6b\x94\xe1\x9f\x0d\x7d\x58\x7f\x64\x4f\x80\xb9\x16\xe2\x9b\xc8\xf5\xc5\xe5\x11\x72\x7d\x71\xf9\x02\x99\x58\x6e\x85\xff\xe6\x50\xaa\x3f\x8b\x21\xca\x10\x2b\xcb\xd3\x67\x52\x5b\x31\x9e\x4e\xb3\xdf\xff\x9f\x8a\x27\xa7\x4c\x3f\xd3\x60\xec\x15\xda\xd6\xea\xb9\x73\x0a\xfb\x55\x8d\x5f\x9b\x5f\xd8\x12\x2b\xe1\xb4\x9b\x5b\xe1\x6f\x07\xf9\x9c\xf9\xb5\x16\xde\x9f\x46\x54\x31\xf2\xe9\xf5\x2d\x2f\xcd\xe2\x76\xfd\x0b\x00\x00\xff\xff\x84\x08\x29\xb3\x29\x04\x00\x00")

func _templateZeroTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateZeroTmpl,
		"_template/zero.tmpl",
	)
}

func _templateZeroTmpl() (*asset, error) {
	bytes, err := _templateZeroTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/zero.tmpl", size: 1065, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"_template/client.tmpl":   _templateClientTmpl,
	"_template/handlers.tmpl": _templateHandlersTmpl,
	"_template/header.tmpl":   _templateHeaderTmpl,
	"_template/main.tmpl":     _templateMainTmpl,
	"_template/registry.tmpl": _templateRegistryTmpl,
	"_template/string.tmpl":   _templateStringTmpl,
	"_template/utils.tmpl":    _templateUtilsTmpl,
	"_template/zero.tmpl":     _templateZeroTmpl,
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
	"_template": &bintree{nil, map[string]*bintree{
		"client.tmpl":   &bintree{_templateClientTmpl, map[string]*bintree{}},
		"handlers.tmpl": &bintree{_templateHandlersTmpl, map[string]*bintree{}},
		"header.tmpl":   &bintree{_templateHeaderTmpl, map[string]*bintree{}},
		"main.tmpl":     &bintree{_templateMainTmpl, map[string]*bintree{}},
		"registry.tmpl": &bintree{_templateRegistryTmpl, map[string]*bintree{}},
		"string.tmpl":   &bintree{_templateStringTmpl, map[string]*bintree{}},
		"utils.tmpl":    &bintree{_templateUtilsTmpl, map[string]*bintree{}},
		"zero.tmpl":     &bintree{_templateZeroTmpl, map[string]*bintree{}},
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
