// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (12.226kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x7b\x93\xdb\x36\x92\xff\x7b\xfc\x29\x3a\x5c\x7b\x25\x9d\xc5\x87\x34\x0f\x8f\x35\xa2\x52\x5e\x6f\x36\xe7\xbb\xcb\xa3\x12\xa7\x6e\xb7\xb2\xa9\x2b\x88\x68\x89\xf0\x80\x00\x03\x80\xd2\x28\x53\xfa\xee\x57\x00\x48\x8a\x7a\xcc\xf8\x95\xbd\xcb\xfc\xa1\x21\x5e\xdd\x8d\xee\x5f\xa3\x1b\x4d\x4e\xbf\xf8\xeb\x77\xaf\xdf\xfe\xe3\xfb\xaf\x20\x37\x05\x9f\x3d\x99\xda\x7f\xc0\x89\x58\xa6\x01\x8a\x60\xf6\xe4\x6c\x9a\x23\xa1\xb3\x27\x67\x67\xd3\x02\x0d\x81\x2c\x27\x4a\xa3\x49\x83\xca\x2c\xc2\xeb\x60\x37\x90\x1b\x53\x86\xf8\x6b\xc5\x56\x69\xf0\xf7\xf0\xa7\x57\xe1\x6b\x59\x94\xc4\xb0\x39\xc7\x00\x32\x29\x0c\x0a\x93\x06\x6f\xbe\x4a\x91\x2e\xb1\xb3\x4e\x90\x02\xd3\x60\xc5\x70\x5d\x4a\x65\x3a\x53\xd7\x8c\x9a\x3c\xa5\xb8\x62\x19\x86\xae\x31\x04\x26\x98\x61\x84\x87\x3a\x23\x1c\xd3\x51\x30\x7b\x62\xe9\x18\x66\x38\xce\x6e\x33\xc9\x04\xfc\x8d\x54\x19\x9a\x69\xec\xfb\xdc\x30\x67\xe2\x16\x72\x85\x8b\x34\xb0\x42\xea\x49\x1c\x67\x54\xbc\xd3\x51\xc6\x65\x45\x17\x9c\x28\x8c\x32\x59\xc4\xe4\x1d\xb9\x8b\x39\x9b\xeb\xd8\xac\x99\x31\xa8\xc2\xb9\x94\x46\x1b\x45\xca\xf8\x3c\x3a\x8f\x5e\xc4\x99\xd6\x71\xdb\x17\x15\x4c\x44\x99\xd6\x01\x28\xe4\x69\xa0\xcd\x86\xa3\xce\x11\x4d\x00\xf1\xec\xd3\xf8\x2e\xa4\x30\x21\x59\xa3\x96\x05\xc6\x17\xd1\x8b\x28\x71\x2c\xbb\xdd\x9f\xc2\xd5\xae\xd7\xd1\x52\xca\x25\x47\x52\x32\xed\xb8\x66\x5a\x7f\xb9\x20\x05\xe3\x9b\xf4\x1b\x3b\x8e\x4a\x11\x33\x19\x27\xc9\xf0\x3c\x49\x86\x17\x49\x32\xbc\x4c\x92\xe1\x55\x92\x9c\xe6\x64\x59\xe9\x4c\xb1\xd2\x80\x56\xd9\x07\xef\xf0\xdd\xaf\x15\xaa\x4d\x7c\x1e\x8d\xa2\x51\xdd\x70\x3b\x7a\xa7\x83\xd9\x34\xf6\x04\x67\x9f\x45\x3b\x14\xd2\x6c\xe2\x71\x74\x11\x8d\xe2\x92\x64\xb7\x64\x89\xb4\xe1\x64\x87\xa2\xa6\xf3\x77\xe3\xfb\x10\x5a\xde\x1d\x82\xe5\xf7\x60\x56\xc8\x02\x85\x89\xde\xe9\x78\x1c\x8d\xae\xa3\xa4\xe9\x38\xa6\xef\x18\x58\xa3\x59\x56\x67\xf9\x68\x08\xf9\x78\x08\xf9\x39\xdc\xdb\xf6\x99\x03\x95\x07\xc0\x04\x76\x08\xb8\xd9\x0d\xae\x91\x2d\x73\x33\x81\x71\x92\xb8\xde\xad\xfd\x21\x7e\x75\x26\xb9\x54\x13\xf8\xd3\x68\x91\x9c\x9f\xd3\x13\x8b\x2e\xf6\x16\x4d\x72\xb9\x42\xb5\xbf\xf4\x72\x34\xa2\xd7\xa3\xf7\x2d\x8d\x56\xa8\x0c\xcb\x08\x0f\x33\x14\x06\x55\x2d\x7d\xc1\x44\x98\xd7\xf3\x47\x49\xf2\xec\xe6\x54\xef\x2a\xf7\xdd\x94\xe9\x92\x93\xcd\x04\x16\x1c\xef\x7c\x17\xe1\x6c\x29\x42\x66\xb0\xd0\x13\xf0\x94\xfd\xc0\x9c\x64\xb7\x4b\x25\x2b\x41\xc3\x53\x7b\x6c\xfa\x16\x8b\xc5\xcd\x27\xaa\xb1\xc3\xc2\x6a\xe1\x7c\x94\x5d\x5f\xc2\x17\xac\xb0\xc7\x1f\x11\x7e\x25\x40\x77\x52\x58\xc8\xdf\x42\xce\x04\x12\x15\x2e\x15\xa1\x0c\x85\xe9\x1b\x59\x0e\xdb\xe5\xc9\x33\xfb\xfc\x62\x34\xbf\xbe\x80\xd1\x85\x6d\x5c\x5d\x8c\x5e\x5e\x27\x70\xee\x46\x5e\x90\xd1\xe5\x0b\x02\x17\x57\xb6\xf1\xf2\x7a\x94\xbc\x18\xc1\xd5\xd8\x36\xe6\x98\x90\xab\x4b\x78\xf1\xd2\x36\x90\x26\xe3\xcb\x6b\x78\xe9\xa6\x2d\x5e\x26\xc9\xe5\x85\x53\xee\xe0\x3d\xe2\xad\x71\x7e\xcb\xcc\x1f\x58\xc2\x63\xc9\x60\x2e\x8d\x91\xc5\xff\xa3\x7c\x0b\xc6\x0d\xaa\x09\x94\x4a\x2e\x19\x9d\xfc\xf5\xef\x6f\x0a\xb2\xc4\xb7\x8a\x08\xbd\x90\xaa\x88\xbe\x61\x99\x92\x5a\x2e\x4c\xd4\x8a\x0d\xda\x10\x65\x5e\x5b\x08\x6a\xa3\xd2\x5e\x2d\x7b\x6f\x08\x28\x68\xa7\xdb\x33\xee\x0d\xbf\xae\x17\xbe\xdd\x94\x98\x26\x30\xe8\x78\x95\xe5\xaa\x50\xeb\xda\x9d\x4a\xa9\x99\x61\x52\x4c\xec\x59\x4f\x0c\x5b\xe1\xa9\xb9\xba\x24\xe2\x68\x01\x99\x6b\xc9\x2b\x83\x07\xae\x36\xe7\x32\xbb\xf5\x7d\x2e\x76\x77\xdd\xb4\x76\xa2\x75\xce\xea\x65\xe0\x18\x41\xa9\xb0\x26\x0f\x25\xa1\x94\x89\xe5\x04\xae\xca\xda\x63\xa1\x20\x6a\xc9\xc4\x04\x92\xdd\x92\x69\xdc\x1c\x70\xd3\xd8\xa7\x29\x4f\xce\xa6\x73\x49\x37\xee\x74\xa5\x6c\x05\x19\x27\x5a\xa7\xc1\xc1\x21\xe2\xd2\x8f\xbd\x09\x36\xeb\x20\x4c\x34\x43\x7b\x63\x4a\xae\x03\x70\x8c\xd2\xc0\x0b\x11\x7a\xf4\x4c\x60\x64\xc5\xab\x97\x1c\xd0\xe3\x21\x5f\x86\xa3\x71\x33\x78\x36\xcd\x47\x0d\x11\x83\x77\x26\x74\x27\x50\x7b\xf6\xd8\x83\x9b\x1d\xa4\x30\xf9\xa8\x21\x1c\x53\xb6\xaa\xe5\xea\x3c\x1e\x88\xf8\xb0\x14\xd7\x50\x3f\xc8\xc5\x42\xa3\x09\x3b\x42\x75\x26\x33\x51\x56\x26\xb4\x2e\x53\xb6\xe3\x67\x53\xd7\x0b\x8c\xa6\x41\xa5\x78\x50\x67\x6b\xee\xd1\x58\x54\xb9\xbd\x04\x0d\x09\x0b\xdc\xd0\xaa\x52\x49\x1e\x40\xc9\x49\x86\xb9\xe4\x14\x55\x1a\xfc\x28\x33\x46\x38\x08\x34\x6b\xa9\x6e\xe1\xa7\x1f\xfe\x0b\x6a\x9d\x33\xb1\x84\x8d\xac\x14\xfc\xa7\x5c\x13\x4e\x80\x50\x6a\xd1\x16\x45\x51\x47\x0c\x07\xbd\x63\x41\xc3\xb9\x11\xbb\x59\x67\xd3\x79\x65\x8c\x6c\x27\xce\x8d\x80\xb9\x11\x21\xc5\x05\xa9\xb8\x01\xaa\x64\x49\xe5\x5a\x84\x46\x2e\x97\x36\x2d\xf5\x5b\xf0\x8b\x02\xa0\xc4\x90\x7a\x28\x0d\x9a\xb9\x01\x10\xc5\x48\x98\x13\x5d\xca\xb2\x2a\xd3\xc0\xa8\x0a\xeb\x4e\xbc\x2b\x89\xa0\x48\xd3\x60\x41\xb8\xc6\x60\xf6\x35\x5b\x21\x14\x08\xde\x8e\x67\x53\xd6\x2a\x86\xc0\x82\x84\x19\x51\x68\xc2\x2e\x55\x46\x29\x8a\x9a\xa6\x43\xc0\x34\xf6\xd2\xf8\x3d\x41\xfd\x37\xad\x78\x43\xa9\xdd\x43\x81\xa2\x82\xbd\x56\xa8\x6c\xa0\x09\x66\xf7\xf7\x8a\x88\x25\xc2\x53\x46\xef\x86\xf0\x94\x14\xb2\x12\x06\x26\x29\x44\xaf\xdc\xa3\xde\x6e\xf7\xa8\x03\x4c\x39\x9b\x4d\xc9\x63\xf0\x04\x29\x32\xce\xb2\xdb\x34\x30\x0c\x55\x7a\x7f\x6f\x89\x6f\xb7\x37\x70\x7f\xcf\x16\xf0\x34\xfa\x01\x33\x52\x9a\x2c\x27\xdb\xed\x52\x35\xcf\x11\xde\x61\x56\x19\xec\x0f\xee\xef\x91\x6b\xdc\x6e\x75\x35\x2f\x98\xe9\x37\xcb\x6d\xbf\xa0\xdb\xad\x95\xb9\x96\x73\xbb\x85\xd8\x12\x15\x14\xef\xe0\x69\xf4\x3d\x2a\x26\xa9\x06\x3f\x7f\x1a\x93\xd9\x34\xe6\x6c\x56\xaf\xdb\x57\x52\x5c\xf1\x1d\x60\x62\x8b\x98\x16\xe6\xce\x6b\x9c\xa8\x5d\x49\x4f\x38\xc1\x32\x6c\xa5\xaf\x01\xa1\x99\xc1\x5b\xdc\xa4\xc1\xfd\x7d\x77\x6d\x3d\x9a\x11\xce\x6d\xa4\x49\x03\xbf\xb5\x76\xd1\x6f\x68\x81\xba\x62\xda\x5d\x80\x66\x8d\x04\x3b\xb1\x3f\xd0\xab\x0f\x0e\x1e\x23\xcb\x09\x9c\x8f\x1f\x39\x75\x74\x11\x9e\x07\xb3\x43\xe8\x29\x7b\x63\x78\x00\x73\xe0\xbd\xcb\xfa\x78\x89\xa8\x7c\x12\x69\xb5\x07\xae\xd9\x11\xee\x03\x99\x59\x15\xcc\x89\xc6\x0f\xe1\xe8\xe2\xc4\x8e\xa3\x6b\x7e\x02\xcb\x1c\x89\x32\x73\x24\xe6\x43\x78\x2e\x2a\x41\x3b\xbb\x74\xde\xfa\xf1\x2c\x2b\xc1\x56\xa8\x34\x33\x9b\x0f\xe5\x89\x74\xc7\xd4\xb7\xff\x45\x18\x38\x8c\x3c\xe7\xb3\x7f\x97\x6b\xa0\x12\x35\x98\x9c\x69\xb0\x47\xf0\x97\xd3\x38\x3f\x6f\xa7\x94\xb3\xb7\x76\xc0\x1f\x5c\x0b\x17\x80\x80\x69\x50\x95\x70\xe7\xb3\x14\x60\x72\x84\xfb\xfb\xe8\x5b\x7f\x80\x6f\xb7\xcd\x51\x1e\xc1\x5b\x69\x23\xf7\x0a\x85\x81\x82\x70\x96\x31\x59\x69\x20\x99\x91\x4a\xc3\x42\xc9\x02\xf0\x2e\x27\x95\x36\x96\x10\xe1\x1c\xc8\x8a\x30\x4e\xe6\x1c\x9d\x12\x34\x48\x05\x24\xcb\xaa\xa2\xb2\x99\x87\x58\x02\x0a\x59\x2d\xf3\x5a\x16\x23\xc1\x9f\x5f\x5c\x8a\x65\x2b\x8f\x2e\x49\x01\xc4\x18\x92\xdd\xea\x21\x28\xfc\xb5\x42\x6d\x34\x10\x85\x60\x18\x52\xbb\x2a\x93\x45\x21\x05\x9c\x2b\x0a\x25\x51\x66\x03\x7a\x3f\x02\x91\x2c\x73\x87\x61\x04\xaf\xc4\x46\x0a\x84\x9c\xac\x9c\x84\xf0\xd6\xdf\xe7\x86\xf0\xb5\xbb\x31\x3f\xb7\x02\xfe\x8d\x64\x38\x97\xb2\x5d\x06\x05\xd9\x34\x7c\xeb\x6d\xac\x99\xc9\x99\xd7\x53\x89\xaa\xb0\x34\x28\x70\x56\x30\xa3\xa3\x69\x5c\xee\x42\xee\xee\x2c\xe7\x61\x2e\x15\xfb\xcd\xc6\x41\x7e\xd2\xca\xe3\xa4\x63\x65\xbb\xd6\x34\xb3\xea\xa4\x8a\x54\x46\xde\xd4\xa9\x51\xc8\x71\xe1\xae\x4f\x6e\xcd\x21\x5c\xeb\x5b\xea\x29\xac\x36\x34\xdd\x75\xc5\x1e\x5d\x13\x38\xf7\x89\x8d\x0f\x49\xd4\x74\x24\xa0\x07\x72\x7a\xa6\xd7\xd7\xe5\x5d\x2b\x47\x93\x1d\xd5\xd2\x3b\x4f\x78\x2b\x0f\xd4\xb5\x62\x1d\x4d\x17\xe4\x16\x81\xc0\x94\x1c\x54\x2e\x6a\xa1\xdd\x1d\x98\xb9\xba\x50\x6c\xd6\x88\xe6\x4b\x1b\xa3\xd2\x1f\x3c\x41\x26\x96\xcf\xc6\x89\x07\xad\x7d\xb0\xe4\x9f\x8d\x13\x26\x8c\x7c\x36\x4e\x92\xbb\xe4\x03\xff\x9e\x8d\x13\x29\x9e\x8d\x13\x93\xe3\xb3\x71\xf2\x6c\x7c\xde\x85\xbb\xef\xf1\x29\x8a\x9d\x83\xda\xf2\x6a\x7c\x20\x00\x43\xd4\x12\x4d\x1a\xfc\x0f\x99\xcb\xca\x4c\xe6\x9c\x88\xdb\x60\xe6\x84\xb5\x51\xcb\xa1\xe3\x54\x9a\x03\x25\xd1\x16\x28\x56\x5a\x87\x9d\xba\xfe\xa5\xa1\xaf\x2b\xe5\xee\x2f\x16\x95\x76\xbf\xce\x81\x45\xcf\x62\xcf\x2a\x65\x10\x4d\xe7\x2a\x9e\xb9\x9f\xd7\xb2\xdc\x84\x8e\x92\xa3\x71\xa4\x47\x5d\x95\xf6\xf6\x11\x75\xf5\x49\x6c\x4a\xcc\x51\xc7\xd7\xc9\xe5\xf5\xd5\xa3\x3b\xd0\x36\x5f\x73\xdb\x68\xc5\x24\x73\xb9\x42\xf0\xd9\xe1\x5c\xde\x01\x11\x14\x16\x4c\x21\x90\x35\xd9\x7c\x31\x8d\xa9\x4b\xc6\x3f\x1f\xb6\xbe\x6a\x15\x96\xbc\xd2\x36\x7d\x65\xd6\x87\xff\x50\x18\xf6\x87\x04\x7c\xcf\x2b\x3d\x84\xb2\x9a\x73\xa6\x73\x20\x20\x70\x0d\x53\x6d\x94\x14\xcb\x99\xeb\xcd\xec\x65\xc5\x35\xa1\x94\xda\x3c\x0c\x08\x2c\xe6\x48\xe9\x09\x48\x7c\x0e\x22\x2c\x4b\x67\xc5\xff\x7b\x0b\x2e\xea\xa3\xf3\x0f\x65\xb5\xe6\x3c\xff\x03\x9b\xec\xc8\x89\xd7\xeb\x75\xd4\x28\xd3\x79\x70\x8e\xbc\x8c\x6d\x9c\xab\x04\x33\x9b\xd8\x1f\x86\x52\xc4\x5f\x32\x9a\x8e\xaf\xc7\x57\x57\xe3\x8b\x97\xd7\x97\x97\xe3\xeb\x8b\xcb\x87\xdc\xbb\xc5\xc5\xa7\x7b\xb7\x4f\xab\xbf\x95\xaf\x2a\x93\xb7\x39\xb5\x87\x4c\x8d\x03\x77\xa1\xa0\xf6\x4e\xa2\x82\x4f\x86\x51\x25\x6c\x6a\x18\x12\x7e\x32\xc5\xfb\x08\x20\x39\x24\x3d\x22\xd9\xfb\xd1\x35\x7a\x0c\x5d\x0d\x82\x2c\x58\x64\x65\xec\x0e\x73\x14\x86\x65\xc4\xda\xa6\x45\xd4\x10\x34\x2b\x4a\xbe\x81\x6c\x67\xf5\x53\xd0\x7a\xd0\x24\xef\x45\xd6\xbe\xd1\x3c\xce\x5c\x92\x57\x48\x8a\x36\xb9\xd3\x95\xce\xb0\x74\x6f\x68\x6c\xc2\xf4\x97\xcd\x6f\x44\x18\x26\xb0\x49\xac\x22\xf8\x4e\xf0\x0d\x54\x1a\x61\x21\x15\x50\x9c\x57\xcb\xa5\xcb\x06\x15\x94\x8a\xad\x88\xc1\x26\x9b\xd2\x35\x26\x5a\x48\x74\xee\x39\x36\xb3\xe5\x9d\x44\xf3\x1f\xb2\x82\x8c\x08\x30\x8a\x64\xb7\xde\x59\x2a\xa5\xac\xb3\x94\xe8\x77\xd3\xe6\x73\x73\xe4\x72\xed\xa6\xf8\x7d\x2f\x18\x72\x97\xdc\x69\x44\xc8\xe5\x1a\x8a\x2a\x73\x1e\x69\x93\x37\xb7\x89\x35\x61\x06\x2a\x61\x18\xf7\xda\x34\x95\x12\x36\x15\xc4\xbd\x1c\xec\xe8\x26\x38\xc5\x62\xf6\x36\xc7\x13\x99\x6f\x7b\x87\x03\x85\xaf\xfd\x74\x28\x95\x34\x98\x59\x73\x02\x59\x12\x26\xb4\xb5\x88\xcb\xf2\xb0\xf8\x80\x3b\x5e\xfb\x54\x3f\xec\xde\x04\xb8\xe1\x38\x86\xaf\xb9\x9c\x13\x0e\x2b\x8b\xf3\x39\xb7\x59\xbb\x84\x5c\xda\xad\x77\xb4\xa5\x0d\x31\x95\x06\xb9\x70\xbd\x5e\x72\xbb\x7e\x45\x94\xb5\x20\x16\xa5\x81\xb4\xae\x96\xd9\x3e\x8d\x6a\x55\x57\xb9\x6d\xd3\xde\xe3\xf7\xc6\x5b\xad\xa7\xf0\xf3\x2f\x37\x4f\x6a\x51\xfe\x8a\x0b\x07\x09\x8b\x6e\xbf\x65\x93\x13\x03\x99\x42\x62\x50\x43\xc6\xa5\xae\x94\x97\x90\x2a\x59\x82\x95\xb2\xa1\xd4\x50\xb6\x03\xa5\xe3\xd6\x10\xe9\xe7\x44\xe7\x83\xba\xd8\xa7\xd0\x59\xa9\x1d\x6b\xfa\xcf\x2c\xea\xfa\x96\x00\x4b\x93\x1b\x60\xd3\x86\x6e\xc4\x51\x2c\x4d\x7e\x03\xec\xf9\xf3\x76\xf2\x19\x5b\x40\xbf\x99\xf1\x33\xfb\x25\x32\x77\x91\xe5\x02\x69\x0a\x5d\x6e\x8e\x61\x4d\x47\x97\x9c\x65\xd8\x67\x43\x18\xf9\xe2\xa8\xfd\x9b\x2b\x24\xb7\x4d\xab\xb6\xa3\xff\xe7\x7e\xb7\x37\xfb\x9a\x71\xca\xdf\xd3\x8d\xaf\x04\x68\x20\xb0\x64\xda\x40\xa5\x38\xd4\x3e\xec\x4d\xd0\x1a\xc4\xcd\xeb\x6a\xe5\x08\x97\xf5\x43\x8d\xa9\x66\x0b\x9e\x4c\xa4\x51\xd0\xfe\x7f\xfc\xf8\xdd\xb7\x91\x36\x8a\x89\x25\x5b\x6c\xfa\xf7\x95\xe2\x13\x78\xda\x0f\xfe\x54\x29\x1e\x0c\x7e\x4e\x7e\x89\x56\x84\x57\x38\x74\xf6\x9e\xb8\xdf\x23\x2e\x43\xa8\x1f\x27\xb0\xcf\x70\x3b\x18\xdc\x9c\xae\x9a\x74\x8a\x3c\x0a\x35\x9a\xbe\x9d\xd8\x02\xff\x50\x47\x04\x0a\x34\xb9\x74\xae\xab\x30\x93\x42\x60\x66\xa0\x2a\xa5\xa8\x55\x02\x5c\x6a\xbd\x03\x62\x33\x23\x3d\x06\x45\x3d\x3f\x75\xd1\xfa\xbf\x71\xfe\xa3\xcc\x6e\xd1\xf4\xfb\xfd\x35\x13\x54\xae\x23\x2e\xfd\x41\x1b\x59\x27\x95\x99\xe4\x90\xa6\x29\xd4\x31\x34\x18\xc0\x97\x10\xac\xb5\x8d\xa6\x01\x4c\xec\xa3\x7d\x1a\xc0\x73\x38\x5c\x9e\xdb\x80\xff\x1c\x82\x98\x94\x2c\x18\x78\x77\x68\x14\x2f\x45\x81\x5a\x93\x25\x76\x05\x74\x17\xe0\x16\x64\x76\x1f\x85\x5e\x42\x0a\xce\x40\x25\x51\x1a\xfd\x94\x88\x12\x43\x1a\xb4\x59\xcc\xba\x69\x69\x0a\xa2\xe2\x7c\x07\x52\xef\x14\x37\x0d\xfc\xf6\xa6\x47\x3e\xd2\x7c\x91\xa6\x50\x09\xea\x54\x4c\x77\x2b\xad\xf1\x7d\x85\x63\x10\xd9\xb8\xb0\x5b\x31\xb8\xe9\xa2\x79\x8f\x1a\xd2\xf7\x91\x43\x7a\x48\x0f\xe9\x03\x04\x5d\xd9\xe8\x31\x7a\xbe\xcc\xd4\x21\xe7\x3a\x1e\xa0\x26\xaa\x62\x8e\xea\x31\x72\xbe\x86\x54\x93\x73\xaa\x7e\x23\x4c\x67\xed\x10\x46\x57\x83\x07\xa8\xa3\x52\xf2\x41\xe2\x42\x9a\x4d\xff\x9e\x93\x8d\xcd\x98\xa0\x67\x64\xf9\xda\xd5\x43\x7b\x43\x17\x71\x27\xd0\x52\x18\xba\x5a\xf2\x04\x7a\xae\x65\xc7\x59\x81\x6e\xd5\x65\x92\x24\x43\x68\xde\xa1\xfc\x85\x58\x27\x54\x15\x6e\x1f\x90\x47\x57\x59\x66\xe3\xfe\xe7\x48\x54\xd3\x68\x65\xaa\xdb\x9f\x21\x55\x1b\x1b\xf6\xc4\x82\x3f\xff\x19\x8e\x46\xf7\x61\x1c\xc7\xf0\x0d\x51\xb7\xae\xee\x53\x2a\x5c\xb9\xda\x50\x3b\xbf\x60\x5a\xbb\x9a\x8b\x06\x2a\x05\xd6\x6b\x3e\xee\xd8\x3f\x92\xb1\x9e\x06\x33\x48\x0e\x05\xb4\xc7\x61\x27\x2c\x9c\x88\x16\x1d\xba\xfb\x81\xe0\x6c\xdb\xe5\xb7\xb7\x92\x15\x08\x5f\xa4\x10\x04\xdd\xc5\x47\x33\xec\x84\x96\xd8\x99\x46\xf3\xd6\xdb\xa2\x5f\x47\xc7\x53\xb1\x6b\x30\x84\xf3\x24\x49\x06\x47\x42\x6c\x77\xea\x7d\x55\xda\xb4\x09\x88\xd8\xb8\x23\xb1\xd5\xad\x4b\x1c\x6d\x0a\x64\x8f\x34\x0e\x99\xe4\xdc\xe7\x2c\xf5\x52\xab\xe0\xba\x46\x96\x42\x38\xba\x39\x11\x45\x3b\x9a\xec\x6c\xed\xd0\x3c\x27\x74\x7f\x68\xa2\x7d\x9d\x1d\x4c\x0e\x47\x7b\x46\xd9\xb3\xd7\x69\xc3\x9c\xb5\x72\xb3\x9d\x46\x0f\xcc\xb5\xb3\xd7\xa1\xce\x3a\xf2\x7b\x3a\xcf\x47\x1f\xb8\x8d\x76\xb8\xac\x74\xde\x3f\x10\x74\x70\x73\x6c\x9b\x37\x06\x95\xcd\x92\xa5\x0d\x59\xd6\x16\xf6\x22\xa0\xf0\xc8\x24\x2e\x55\x57\x18\x2a\x14\x14\x55\x93\x52\xf8\xcc\xde\x26\x80\x7b\x26\xf3\x17\xcb\x2e\x9c\x3e\xd2\x61\x5c\x4a\x26\x05\x02\x00\x1c\x38\x81\x03\xea\x1e\x52\xed\x64\xe4\xa4\xd4\x48\x21\x05\xff\xb1\x49\x7f\x10\x55\x82\xdd\xf5\x07\x61\xdd\x3e\xa4\xd1\x8c\xdf\xb4\x97\xc4\x46\xec\xe7\x29\x04\x53\xa3\x80\xd1\xb4\x17\xc0\xf3\x53\x2e\x68\xa3\x6e\x6f\xb6\x93\xa0\xbb\x14\x60\x6a\xe8\xcc\x95\xbb\xfd\x6d\xed\x9f\x41\xf7\x3d\x7f\xa5\x78\xff\x88\x2c\x59\x11\x43\x94\xa3\x3a\xb8\xe9\x7c\x16\x50\x5f\x13\x33\x6b\x9c\x1b\xf0\xf7\x51\x57\x55\x87\xe6\x13\x12\xdf\x9a\x4b\x45\x51\x85\x8a\x50\x56\xe9\x09\x5c\x94\x77\x37\xff\x6c\x5e\xe8\x4c\x63\x43\x1f\x17\xb5\x54\x38\x3b\x92\xa8\xae\x21\x3f\x87\x60\x1a\xdb\x09\xef\x23\xd3\x6e\xb6\xfb\x2a\x1d\x4e\xbc\xa7\x83\xf6\x45\x77\xdd\x5f\x30\x4a\x39\x5a\x81\x77\xe4\xad\x33\x5a\xfb\x77\x5d\x6a\x9f\x25\xd4\xaf\x27\x76\x6b\xb6\x80\x5c\xe3\x23\x0b\xda\x37\x1d\x3d\x0b\x80\xd0\x6e\x99\x39\x9d\xd7\x57\x6d\xd7\xad\x7a\x4e\x17\xf5\x27\x4b\xb4\x52\x2e\xd7\xea\x87\x35\xc0\x86\xd0\xd3\x36\xf7\xa3\xba\x37\x88\xf2\xaa\x20\x82\xfd\x86\x7d\x1b\x97\x06\x5e\x57\xee\xd5\x49\x70\x7c\x24\x1f\x09\xb3\x7b\x1f\xd2\x6b\x62\x5c\xaf\x56\x62\xaf\xb1\xee\xc5\xee\x66\x3f\x81\xe4\xa6\xf7\x91\x1a\x3a\xcd\x25\x9c\x13\x05\xdd\x46\xd8\x04\x5f\x50\xd2\x72\x6f\xc6\xe6\x44\xf5\x7c\x1d\xc3\xe5\xe7\x42\xae\xd3\xde\x79\xd2\x0a\xe9\x0d\xed\xec\xdc\xab\xb1\x76\x64\x0c\x2b\x65\xe3\x9a\x33\x38\x4f\x7e\x0f\x69\x7d\x2d\xe4\x60\x07\x46\xb1\x12\x29\x90\xcc\xb0\x15\xfe\x0b\x36\xf2\x3b\x28\xf9\xa3\x45\xb4\x38\x6c\x94\xe7\x60\xba\x27\xaf\x1d\x6d\x75\xfb\x6f\xd6\xdf\x20\x76\x1a\x7e\x0e\xc1\xc9\x8d\x3c\x88\xc4\x83\x89\x07\xae\xfd\xb0\xdf\x4f\x63\xa3\x76\x43\xdb\x4e\xb6\xdb\x9c\x24\xc1\x20\xca\x4d\xc1\xfb\xc1\xd4\xb8\x4f\x5e\xac\xcc\x2d\x05\x47\xc0\x77\xef\xa7\x74\xdb\xfd\x8b\x8c\xbd\xbf\xe3\xc1\x3d\x0b\x3a\xc9\x49\x7b\x17\x6b\x32\x11\xd8\xee\xbe\x0c\x8a\x63\xf8\xd1\x10\x65\x80\xc0\x4f\x6f\xa0\x2a\x29\x31\x36\x7a\x49\xb0\xf1\xd1\x97\x9d\x9b\x4f\x87\xe6\x44\x69\x58\x48\xb5\x26\x8a\xd6\xf5\x19\x93\xe3\xc6\xbd\xb1\x6b\x52\x3f\x8d\xe6\x8d\x3d\xc5\x56\x84\xf7\x8f\xee\x7d\x4f\xfb\xbd\xa8\x6b\xf2\xde\x20\x42\x92\xe5\xc7\x13\x5d\xc4\x6a\xf9\xa6\xf0\xad\xbb\x02\xf4\x9f\xf6\x4d\xce\xf4\x20\x22\xc6\xa8\x7e\x6f\x0f\x0c\xbd\x81\xb5\xeb\xa8\x73\x25\x6b\x97\x4f\xf7\xdc\xea\x31\x1a\xbb\x64\xba\x4d\x04\x9a\xe9\x99\xd6\x7d\x8f\xab\xde\xb0\x43\x7b\x1f\x56\xbd\x67\xbd\xd6\x50\x3b\xf7\xde\xed\x23\x3d\x29\xc9\x1e\xe9\x9e\xf5\xb2\xde\x11\x7b\x42\xe9\x6b\xeb\x3f\xfd\xe0\x84\xa7\x1f\xa2\x63\xd0\x2a\xdb\x9f\xd7\x8f\x6a\xd9\x7f\xa4\xf1\x80\x8a\x19\xed\x0d\x22\x5d\xcd\x7d\x6d\xa2\x7f\xd9\x5e\xc0\x9a\x69\x0e\xbc\x87\xa1\xe0\x28\xa1\xb0\x2c\xf6\x93\x8a\xf0\x20\x09\x79\x24\x6a\xd4\x2c\xfd\xae\xb6\x43\xab\xf0\x64\xd0\x96\xb6\xbe\xd2\x36\xb9\xf2\xb5\xff\x35\xce\xb5\xab\x24\x40\x8d\x77\x57\xcd\xf1\x55\x9b\x57\xdf\xbf\xe9\x54\x6e\x5a\x8f\xe8\x3b\xea\xed\xf7\xb6\xa7\xea\x24\x27\x3f\xf0\x5d\xaf\xd7\xf5\xb7\xd8\xae\x88\xdf\x16\x52\x62\x52\xb2\xe8\x9d\x0e\x80\xe8\x8d\xc8\x80\xe2\x02\xd5\xac\x43\xbe\xae\xae\x4c\x63\xff\x81\xdb\x34\xf6\x5f\xec\xff\x6f\x00\x00\x00\xff\xff\x5d\x36\x02\x9e\xc2\x2f\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0x43, 0xb9, 0x79, 0x46, 0x7c, 0xd9, 0xcf, 0x28, 0x85, 0xe, 0xd5, 0x96, 0xc4, 0x7a, 0x59, 0x6f, 0xba, 0x83, 0x0, 0x78, 0x4, 0x80, 0xa9, 0xf0, 0x3d, 0xc1, 0x21, 0xa6, 0x6a, 0x27, 0xa2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"faucet.html": faucetHtml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
