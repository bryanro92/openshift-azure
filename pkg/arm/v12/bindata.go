// Package arm Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5d\x73\x1b\xb7\xce\xbe\x2e\x7f\x05\xba\x72\x9b\x26\x35\xb5\x76\xda\xbe\x6f\x47\xad\x33\xe3\x38\x4e\xc7\xa7\x6e\xec\x63\x27\xa7\x17\x69\x26\x43\x2d\xb1\x12\x23\x2e\xb9\xe5\x87\x64\x55\xd1\x7f\x3f\x03\xee\x4a\x96\x2c\xd9\x49\x1a\xf7\xe4\xc2\x91\x48\x10\x00\xc1\x87\x78\x00\xaa\xf3\x65\xde\x57\x26\xef\x0b\x3f\x04\x8e\x57\x8c\x75\xe0\xb9\x75\x10\xd0\x07\x65\x06\x3d\xd0\x76\x00\xc2\x48\x90\xce\xd6\x20\xb4\x86\xe0\x44\x59\xaa\x02\xc2\x50\x04\x98\xd8\xa8\x25\x38\x1b\x03\xc2\x58\x09\x08\x43\x84\x4a\xf8\x80\x0e\x8e\x4f\x9f\xb2\x0e\x5c\x1c\x5f\x9e\xbd\xba\x38\x3a\xfe\xe5\xe2\xec\xd5\xf9\x41\x36\xb5\xd1\x71\x87\xde\x46\x57\x20\x1f\x38\x1b\xeb\x8c\x75\xe0\xec\xf2\xed\xf3\x7f\x3f\x7b\x71\x90\xd9\x1a\x8d\x1f\xaa\x32\x74\x77\xd6\x56\x76\xad\x17\x12\xc7\xdd\x42\xdb\x28\x33\xd6\x61\x1d\x50\x75\x10\x7d\x8d\x1e\xf8\x09\x9c\xbc\x38\x7f\xf5\x12\xb8\x87\x9d\x6f\xa4\x1a\xc0\xb7\x7e\x68\x5d\x80\x6c\xa7\xd5\x9b\xc1\x7b\x08\x42\x69\xe0\xfb\x0f\x81\xbf\x83\xd3\xb3\x5f\x80\x73\x6d\x07\xbc\x76\x58\xaa\x2b\xc8\x7e\x7d\xf5\xf4\x18\x48\x14\x9e\x5d\x9c\x9d\xf7\xb2\xcf\xd3\x4f\x3a\x18\x9b\xcd\x40\x95\xd0\x3d\xb2\xa6\x54\x83\xee\x25\x16\xd1\xa9\x30\x3d\x17\xa1\x18\x9e\x8b\x62\x24\x06\xe8\x61\x3e\x67\xda\x0e\x06\xe8\x80\x87\x36\x70\xdc\x07\xe1\x42\xac\xbb\x7e\x08\x99\x32\x3e\x08\xad\x95\x19\x80\x43\x09\x14\xf2\x42\x1a\x28\x92\xce\xe8\x44\x50\xd6\x80\x35\xb0\xf3\xcd\xd0\xfa\x60\x44\x85\x0f\x33\x56\x88\x00\x4f\xf2\xb1\x70\xb9\x56\xfd\x7c\x1a\xab\xbc\xd0\x0a\x4d\xe0\x05\xba\xd0\xad\xb1\x82\x9f\x7f\x7e\x70\x7c\xf6\xfc\x01\xb9\x78\x84\x2e\x1c\xfa\xa7\xd3\x80\x7e\xe9\x2b\x8d\xa9\x52\x15\x22\xa0\xef\xb6\xbe\x5e\x60\x6d\xbd\x0a\xd6\x4d\xd3\x34\xbc\x87\xcb\xe0\xc8\xaf\xf9\x9c\x1d\x9f\x3d\xbf\xdd\xe8\x08\xa7\x37\x6d\x9e\x3b\x35\x16\x01\x7f\xc5\xe9\x27\x5a\xfe\x15\xa7\x1b\x86\x3f\x3a\x80\x87\x17\x67\xe0\xdb\x53\x80\x58\x4b\xb2\x01\xaf\x67\xb3\x56\x9f\xff\x97\x55\xe6\x03\xc7\x95\xed\x42\x06\xf3\xf9\x9b\x8d\x90\x97\xd6\x81\x08\x01\xab\x3a\x80\x32\x30\xdb\xef\x76\x7f\x98\xff\x04\xd2\x32\x80\x69\xac\xa0\x75\x03\xf8\x14\xf8\x9f\xf0\x69\x36\x93\x49\xf8\xfa\x6b\xe8\x3b\x14\x23\x06\x70\xe7\x86\x5f\x2f\xdc\xd8\x99\xb5\x9f\xe6\x6f\xb6\x6f\xbd\xf5\xa9\xc1\x50\x29\x94\x46\x99\x31\x20\xcc\xbe\x7e\xbd\xb2\x1a\xb8\x0e\xf0\x03\xbc\x79\xf3\x13\xdd\x6e\x03\x5e\x23\xd6\xb0\xff\x13\xa0\xf6\x08\x78\xa5\x02\x7d\x29\x15\x93\xd6\xe0\x07\x4e\xc3\x61\x65\xc7\x9f\x06\x66\x8a\x5e\xa1\x51\x18\x4a\x3e\xcc\x55\xc0\x5d\x09\x77\x82\xfb\x0e\x10\xb2\xd9\x0c\x8d\x9c\xcf\x29\xcb\x15\x0e\x45\x40\xb2\x1e\x84\x32\xe8\xa0\x8e\x5a\x53\x94\x1c\x06\x56\x8d\xa4\x72\xc0\xeb\x6b\x65\xd6\xa9\x81\x32\x79\x57\xda\x62\x84\xee\x06\xdc\xd7\x27\xf3\x66\x47\xdd\x77\xde\x9a\x55\xd8\x77\x9f\xa1\x53\x63\x94\xdd\x23\x5b\xf5\x95\x41\x79\x52\x89\x01\x9e\x47\xad\x2f\x93\xd5\x05\x10\x36\x20\xae\x0d\xe5\x9e\x5b\xac\x41\xee\xac\x0d\x39\x6d\xe9\xe5\xd9\xb3\xb3\x1e\x48\xd4\x18\x30\xa5\xe2\xd2\x6a\x6d\x27\xa4\x29\xa5\xda\x66\xcf\x14\x65\x51\x52\x8a\x56\x01\x94\x87\xbe\x18\xa1\x04\x65\x82\x05\x1b\x1d\xfc\xe7\x37\x50\xe4\x97\x67\x69\x8d\x90\x12\x78\x09\xed\xb6\x99\x2a\xe1\x4b\x18\x38\x5c\x89\xcc\xc2\x0d\x0c\x45\x5e\xfa\x20\xfa\x0d\x50\x18\x80\x9f\xfa\x80\x55\x11\x34\xf8\x60\xeb\x56\x07\x4f\xa7\x19\xeb\x6e\x50\x15\xba\x0f\x4a\x79\x74\x63\x55\xe0\x6d\x72\x2b\xf3\xd5\xa8\xf4\xdd\xab\xd2\x93\xbb\xb9\xc4\x71\x2e\x95\x1f\xe5\xe2\xaf\xe8\x30\x5f\x52\x4e\x2d\x5c\xd8\x67\x00\x58\x0c\x2d\x3c\xb8\x5b\x0c\x36\xf6\x08\xa4\x1e\x06\xae\xfe\x33\xda\x20\x00\xf6\x60\xef\x01\x3c\x79\x72\xbd\x75\x72\xc3\x46\x13\x6e\xae\x64\x00\x0e\x7d\xb0\x0e\x0b\x6b\x80\x5f\x6c\x99\x6f\x10\x45\x9a\x5a\x14\x49\x81\x95\x35\x37\x50\xc4\x00\x32\x22\x2e\x49\x48\x72\x59\x0f\xb2\x77\x36\x3a\x23\xb4\xcc\x76\x69\x4e\x2a\x4f\xac\xc5\x35\x0e\x44\x31\xe5\x0e\x07\xca\x07\x37\xcd\x7a\x10\x5c\x44\xd6\xe0\x69\x3d\x96\xc2\x85\xcd\x60\x6e\x17\xb8\x71\x76\xa5\x62\xac\x8d\x4c\xba\x3c\x84\xf1\x36\x97\x25\x68\xfb\xee\x0b\x2b\x31\x65\xaf\x27\x29\xd4\x86\xa4\xbe\xde\x8a\x22\x0c\x85\xdc\x86\xa1\xe5\xa9\xde\x3c\x2b\x5f\x78\xb5\x9f\xeb\x68\xf6\xe0\xfd\xfb\x66\x77\xb7\x1d\xeb\x8a\xe8\x0d\x83\xcd\x81\x4a\x2c\x45\xd4\xc1\x7f\xd4\x81\xd2\xba\xdb\x8f\x33\xcd\x52\x5c\x88\x11\xa4\x4f\x6c\x10\x8a\x7a\xf7\xc7\xef\xbf\xff\x3e\xf1\xc1\x17\xb5\xb3\xc1\x1e\xec\xcc\xa4\x0f\x5f\x7d\xb5\xfb\x68\xce\xbe\xa8\xad\x0b\xcd\x40\xa7\xf3\x68\x77\xce\xbe\xb8\x2e\x3d\x0e\x53\x69\x74\x72\x71\xfc\xfb\xe1\xe9\xe9\xdb\xc3\xd3\xd3\xb3\xdf\x29\x2b\xed\x24\x25\xc0\x2b\x3a\x9d\x80\xc0\x79\xf3\xff\x8b\xe3\xdf\x69\x70\x31\xcd\x25\xa9\x86\x9d\xf4\x97\xbf\x83\xc3\xa3\xa3\xe3\xf3\x97\xc0\x27\x6d\xae\x5e\xd8\xe1\x5e\x8c\xb1\x05\x9f\x9f\xfa\x26\x7d\xe5\x8b\x59\xca\x2c\x93\x94\xf9\x09\x09\x94\x4c\x0c\x9d\xea\x44\x88\x01\x9a\x90\x8a\x43\x83\x61\x62\xdd\x08\x62\x50\x5a\x05\x85\x1e\x06\x36\x31\x4c\xb0\xe0\x44\x91\xb2\xac\x54\x94\x79\xba\x54\x59\x95\xcb\xc5\x2e\x1a\x0f\x7d\x2c\xad\x43\x90\xc6\x53\x3a\x1a\x19\x3b\x31\x10\x6c\x4a\x60\x8d\x25\x04\x34\x12\x62\x0d\x13\x15\x86\x40\xac\x34\x05\x9f\x32\x24\x9b\x0c\x95\xc6\x44\x58\x4b\xd2\x00\x2e\x1f\xc2\xc1\x01\x64\x59\x22\x2d\x69\xaf\x29\xeb\x23\x28\x8a\x80\x4c\x7b\xdc\xc4\xf2\x65\x23\x05\xf3\xf9\xdd\x7c\x7f\xf7\x8d\xb8\xd6\xf2\x79\x94\xfe\xd1\x56\x3e\x95\xd9\xff\x6f\xef\x36\x6a\xa7\xaa\xfb\xc5\xd9\xcb\xe3\x1e\x9c\x18\x28\x63\x88\x0e\x77\xa1\xb2\x63\x6c\x7a\x01\x65\x4a\xeb\xaa\x96\xc5\x63\xf0\x4a\x22\xd8\x12\xd0\x8c\x95\xb3\xa6\xa2\xe3\x1e\x0b\xa7\x1a\x4c\x75\x98\xc7\x00\xdf\x5e\x31\xbc\x4a\xe8\xbc\x3c\xbc\x7c\x75\x71\x72\xf0\x60\x65\x2b\xbf\xa5\x48\xb4\x3b\x69\xe6\x61\x3e\x7f\x90\x16\xf2\xab\x45\xe2\x71\xd1\x00\xe7\xb5\x53\x63\xa5\x71\x80\x12\x38\xa7\x22\x81\x2f\x20\x49\xa8\x00\x3e\x86\xbc\x97\xd3\xc7\xde\x5f\xc0\xb1\xb5\x76\x77\xdc\xda\x13\x60\xd1\x90\xc1\x66\x05\x63\x4d\xf1\xc4\x0b\xc1\x83\x8b\x3e\xd0\xdd\xf0\x18\xd2\xad\x88\x35\x0c\xd0\xe0\x58\xa4\xd3\xa4\x11\x1f\x44\x31\x02\xe1\xc1\x5b\xe2\x5c\x9f\x20\xbd\x5e\xee\x28\x0f\x5a\x28\x49\x01\x83\xfe\x94\x75\x92\x48\x6b\xfa\xba\x36\xd9\x6d\x56\x6a\xeb\xd1\x41\x18\xaa\x74\x51\xda\x2b\x72\x8b\x70\x65\x1d\xb2\x0e\xb9\xe2\xa1\x74\xb6\x5a\x93\xad\x9d\x2d\xd0\x7b\xba\x59\x13\x45\x55\xcf\x50\xd5\xa4\xaf\xf1\x9f\x35\x6e\x78\x04\x3f\x6c\xfa\xbb\x48\x75\x59\x81\x20\x40\x8a\x29\x58\xa3\xa7\xb4\x9b\x3a\x39\x83\x04\x45\xcf\xf2\xe8\x5d\xae\x6d\x21\x74\xea\x27\xc5\x5f\x1e\x0b\xd9\x6e\x96\xaa\x97\xbe\xf0\xa8\x95\xa1\xdb\x09\xe7\xfb\xcf\x3e\x28\xef\x6d\x19\x26\xc2\x7d\xb4\x7c\xa1\x45\x25\xc6\x0b\x69\xd6\x01\x34\x84\xb4\x94\x9e\x1a\x0a\x5b\x3f\x95\x96\xea\x3c\xbb\x66\xba\x68\x2a\xe1\x47\x50\x49\x2f\x17\x4c\x08\x8d\x9d\xf5\xaf\x95\x35\xd7\x23\xa5\x8e\x68\xc2\xf2\xfb\x8a\xba\xd6\x81\xfb\x52\xd7\x6c\xe2\xf3\xb4\xb1\x0e\x9c\x2b\x03\xa3\xd8\xc7\x26\x72\x09\x45\xd1\x23\xa4\xc8\x82\xa8\x15\x27\x59\x74\xcc\xd3\x55\x52\xc0\x1d\x42\xe6\x3b\xdf\xc0\xa3\x66\xbc\x07\x0f\xbb\x8f\x3a\x7f\xec\x0f\x43\xa8\x7d\x2f\xcf\x57\x6a\xf5\x4e\xd6\xf0\x77\x5b\x9e\x36\x89\x2c\x17\xb2\x52\xa6\x7b\x6d\xf1\xde\x14\x2f\x9f\x0a\x78\x33\x70\xaf\x36\x88\x77\xd2\x9f\xfb\xd7\xea\xe5\x3d\x84\x23\x95\xfc\x49\x4d\xdb\x6c\x30\x36\x9b\x71\xca\xf0\x06\x61\xa7\xfb\x54\x14\xa3\x58\x3f\xd5\xb6\xff\x82\x08\x31\xcb\x3e\xf8\xd0\xb0\xe4\x76\x2a\x69\xc6\xe8\xa6\x1b\x8d\x18\x65\xba\x40\x34\x0a\x03\x0c\xe9\xde\xf7\x93\x95\xd4\x93\x5d\x94\xeb\x25\x50\xfe\x88\x11\xc7\x90\x1f\xcf\x94\x3b\x58\x9f\x6b\xd7\x35\x3d\xd6\xce\x8a\xdc\xdf\xa6\xe6\xe3\x50\xc8\x66\xcf\x9f\xc9\xce\x6b\x8a\xfe\x49\x82\x5e\x37\x74\x7f\x1c\x7d\xa7\x9f\xd2\x4e\x8c\xb6\x42\x52\x10\x9b\x43\xc8\xd6\x69\x74\x93\x39\xff\x60\x90\xd8\x73\xe3\xfe\xf5\x36\x87\xb6\x09\xa7\x07\xbb\xda\xd9\xb1\x92\xe8\xf2\x5e\xfe\x56\x8a\x20\xf2\xb7\x44\x77\xad\xf4\x2a\x00\x7a\xb9\x8d\x44\xd1\x34\xf5\xa1\x98\x11\x94\x9a\x4d\x34\x9a\x78\xbf\x85\xfb\x01\xad\xbc\x71\x03\xe6\xf3\x56\x48\xa6\x77\xcd\xc4\xbd\x07\x64\xac\x05\x63\x57\xf6\x5b\x01\x51\xa4\xb9\x45\xa8\xee\x0e\x68\x6b\x7f\x21\x4c\x47\xb8\xb8\x26\x8f\x17\x1d\xc2\xdf\xc5\x74\x53\xfe\xd0\x9e\x3f\x13\xd3\x6b\x8a\xfe\x49\x4c\xaf\x1b\xfa\x1f\x61\xba\x89\x72\xe2\x75\x23\x6a\x3f\xb4\xe1\x93\x30\x4d\x28\xea\x2d\x3f\x2d\xa7\x56\xf3\x55\x6f\xfd\x5b\x83\x4e\x8e\x70\xfc\xf2\xe8\xd9\xd1\xcb\xd3\xb7\x87\xe7\x27\x07\xd9\x77\xd9\x2d\xa0\x5d\x0f\x0a\xc9\x90\x96\x44\xe8\xad\xbf\x0b\xa0\xac\xdd\x84\x0d\x5c\xd2\xbd\xe1\x94\x30\xd7\x73\xa9\xc1\x49\x2b\x90\x5a\x9f\x95\x8c\xdd\x0e\x2b\xa3\x82\x12\x9a\x17\x3a\xa6\x3b\x9a\xb5\x31\xdc\x4b\xff\x0e\x16\xfc\xb2\x36\xda\x7b\xfc\xdd\x8f\x7b\xbb\xab\x43\xfb\x5b\x05\xf7\x37\x05\x1f\x6f\x15\x7c\x9c\x04\xb3\xed\x2e\xf1\x60\x47\x68\x52\x58\x78\x69\x1d\x4f\x3d\xfb\x0d\x51\x21\xc7\xe8\x82\xf2\xc8\x6b\x44\xc7\xa3\xd3\x1e\xb6\x50\x63\x32\xc3\x58\x35\xde\x8c\x52\xfe\xe8\xc6\xd8\xc6\x5b\xe2\x32\x9e\x6b\x94\xb4\xd6\xe7\xdf\xd0\xfb\x31\xc8\xc4\xd4\x74\x66\x89\x9e\xa9\x8b\x9d\xcf\x19\x0b\xd1\xa0\xe4\x42\x56\x54\x88\x97\xd4\xc0\x5e\x17\x33\x54\xc6\x3b\xab\x79\xad\x45\xea\xb9\xa8\x46\x17\xda\x5b\x30\x88\xf2\x5a\xae\x9b\x0a\xb6\xee\xd8\xea\x58\xa1\x07\x02\x46\xf3\xa0\x29\x17\xed\xf4\x55\xe9\xa1\x79\xa6\x2a\xa8\x89\xa6\x4e\x7b\xf1\xac\x59\xc1\xde\xff\xff\xb0\xb7\xed\x79\xf3\x16\xfd\xe4\x47\xf3\xb2\x94\x4a\x04\x3f\xf5\xda\x0e\xc0\x2b\xea\x09\x26\x08\x95\x30\x62\x80\x80\x54\x37\x84\x21\x89\x84\xa1\xb3\x71\x30\x84\xc5\xe3\xd4\x4a\x1d\xdb\xbe\x50\x2d\xb4\x6c\xad\x74\x6d\xbd\x31\xcd\x3a\x60\x6c\xc0\x1e\x88\x60\x2b\x55\xf0\xeb\x88\xa5\x37\x82\xc2\x09\x3f\x04\x6d\x6d\xed\x21\x9a\xa0\xf4\xe2\x67\x28\xe5\x21\xd6\x9b\x55\xf9\x56\x2d\x4b\x63\xf7\xf1\xd3\x8d\x2f\x86\x28\x63\x0a\xd8\xea\xad\x74\xd8\xb7\x36\x50\xd9\x5d\xd8\xaa\x4e\x2f\xb5\xdb\x5e\xe7\x33\xe6\x87\x31\x10\xb1\x50\x0a\x6b\xd6\x7c\xfb\x98\xcd\x66\x94\x22\xe7\xf3\x8d\xbe\xe0\xce\xfd\x2c\x1f\xc8\x16\xcf\xdf\xff\x0d\x00\x00\xff\xff\x97\x41\xe5\x9c\xf5\x1b\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x5d\x6f\xdb\xc8\x15\x7d\x9f\x5f\x71\x56\x5a\xc4\x5d\x6c\x28\x25\x05\xd2\x02\xce\x26\x40\x9a\x64\x81\x74\xd1\xb5\x61\x6f\xdb\x87\x20\x0f\x23\xce\x25\x35\xab\xe1\x5c\x66\x3e\x24\x2b\x8a\xfe\x7b\x71\x49\x2a\x96\x2d\x5b\x4e\x90\xbe\x49\xe4\x9d\xfb\x79\xee\x39\xc3\xf1\x0f\xd3\x99\xf5\xd3\x99\x8e\x73\x14\x74\xa5\xd4\x66\x03\x5b\x61\xf2\x9a\x7d\x65\xeb\xc9\x25\x95\x39\xd8\xb4\x3e\xd7\xa9\x9c\x9f\xeb\x72\xa1\x6b\x8a\xd8\x6e\x95\xe3\xba\xa6\x80\x22\xc1\xb3\xa1\x22\x26\x1d\x52\x6e\x27\x71\x8e\x91\xf5\x31\x69\xe7\xac\xaf\x11\xc8\x60\xae\x13\x4a\xe3\x51\x76\x1e\x73\xd0\xc9\xb2\x07\x7b\xfc\xf8\x97\x39\xc7\xe4\x75\x43\x3f\x8d\x54\xa9\x13\x5e\x4e\x97\x3a\x4c\x9d\x9d\x4d\xd7\xb9\x99\x96\xce\x92\x4f\x45\x49\x21\x4d\x5a\x6a\xf0\xcb\x2f\x27\x6f\xcf\x7e\x3d\x91\x04\x5f\x53\x48\xaf\xe2\x3f\xd6\x89\xe2\x97\x4c\xe5\x99\xad\x6c\xa9\x13\xc5\xc9\x90\xe9\x05\xb5\x1c\x6d\xe2\xb0\xee\x5e\xe3\x33\x2e\x53\x90\xbc\xb6\x5b\xf5\xf6\xec\xd7\xfb\x83\x2e\x68\x7d\x3b\xe6\x79\xb0\x4b\x9d\xe8\x37\x5a\x7f\x63\xe4\xdf\x68\x7d\x10\xf8\x2b\xdb\xf7\xea\xe2\x0c\x71\x98\x00\x72\x6b\x24\x02\xde\x6f\x36\x83\xb7\xf8\x4f\xb6\xfe\x81\x51\x8d\x1e\x63\x84\xed\xf6\xc3\x41\xc3\x2b\x0e\xd0\x29\x51\xd3\x26\x58\x8f\xcd\xd3\xc9\xe4\xd9\xf6\x39\x0c\x2b\x60\x9d\x1b\x0c\x69\xa0\x58\xa3\xf8\x88\x6f\x8b\xd9\x85\xc4\xa3\x47\x98\x05\xd2\x0b\x05\x1c\x29\xf7\xfd\x2e\x89\x1f\x37\xc3\xaf\xed\x87\xbb\x0b\x1f\x32\xea\xf1\x53\x69\xeb\xc8\x8c\x14\x04\xad\xef\xdf\xef\x9d\x46\xe1\x12\x9e\xe1\xc3\x87\xe7\x48\x73\xf2\x88\x8e\xa8\xc5\xd3\xe7\x20\x17\x09\x74\x65\x93\xfc\xa9\xac\x32\xec\xe9\xe8\x24\x02\x35\xbc\xfc\x36\x18\x4b\xe7\x4a\x47\xda\x43\x3b\xa7\x42\x83\x22\x54\x38\x0a\xeb\x23\xf0\x53\x9b\x0d\x79\xb3\xdd\x2a\x35\x46\x19\x48\x27\x92\xe8\x49\x5b\x4f\x01\x6d\x76\x4e\x7a\x14\x28\xa9\x66\x61\x6c\x40\xd1\x5e\x3b\xe3\x60\x6b\xeb\xa7\x13\xc3\xe5\x82\xc2\x2d\xa0\xdf\x7c\x39\xed\x2b\x9a\xfc\x19\xd9\xef\x03\x7e\xf2\x86\x82\x5d\x92\x99\xbc\xe6\x66\x66\x3d\x99\x77\x8d\xae\xe9\x3c\x3b\x77\xd9\x45\xdd\x81\xe0\x00\xdc\xce\xa3\x88\xf7\xa5\x82\x69\x60\x4e\x53\x29\xe9\x8f\xb3\x37\x67\xa7\x30\xe4\x28\x91\x8c\x0a\x15\x3b\xc7\x2b\xf1\x54\x07\xce\x6d\x5f\xb3\x74\x59\x57\x89\x02\x6c\x82\x8d\x98\xe9\x05\x19\x58\x9f\x18\x9c\x03\xfe\xf3\x2f\x58\xc9\x2b\xaa\xee\x8c\x36\x06\x45\x85\xa1\x6c\x65\x2b\xfc\x80\x3a\xd0\x5e\x67\x76\x69\x50\x2a\xa7\x55\x4c\x7a\xd6\xc3\x44\x01\x71\x1d\x13\x35\x65\x72\x88\x89\xdb\xc1\x47\xd1\x4d\x33\xb7\x93\x64\x1b\x0a\x0f\x5a\x45\x0a\x4b\x5b\xd2\x7d\x76\x7b\xef\x9b\x45\x15\x27\x57\x55\x94\x74\xa7\x86\x96\x53\x63\xe3\x62\xaa\x3f\xe5\x40\xd3\x40\x91\x73\x28\xa9\x68\x75\x48\x4f\x15\x40\xe5\x9c\x71\x72\xdc\x0c\x07\x35\x42\xdc\xa3\x0e\xed\xc7\xcc\x49\x03\x4f\xf0\xe4\x04\x2f\x5f\x5e\x97\x2e\x69\x70\xf6\xe9\xf6\x49\x05\x04\x8a\x89\x03\x95\xec\x51\x5c\x1c\xbc\xdf\x6c\x0a\xd9\x3b\xfa\x88\xc9\x05\x3b\x12\xd2\xaa\x82\x96\xad\x57\x40\x0f\x36\x09\x32\x00\xcc\x68\x6a\xd8\xdf\x02\x98\x02\x46\x8e\xeb\xc2\x08\xc8\xc2\xe8\x14\xa3\x3f\x39\x07\xaf\x9d\x19\x3d\x96\x77\xc6\x46\x3d\x73\x54\x38\xaa\x75\xb9\x2e\x02\xd5\x36\xa6\xb0\x1e\x9d\x22\x85\x4c\xaa\x87\x9a\xe4\x41\xde\xf4\x71\xf7\x3b\xae\x43\x3a\x6c\xf9\xdd\x06\xb7\x26\x5c\x59\xa5\x86\xfe\x75\x2b\x26\x9b\x30\xb0\x5d\xb7\x00\x71\xf2\x3b\x1b\xea\xf8\xed\x65\x37\x10\x2f\x56\x8f\x04\xd0\xab\x8e\x6e\xc4\xb5\x60\x58\xe8\x04\x2b\xad\x6b\xf2\x09\xda\x1b\x78\x4a\x2b\x0e\x0b\xe4\x64\x9d\x4d\x96\x22\x6a\xee\x68\x2d\x31\x82\x2e\xbb\xe5\x36\x56\x00\x3f\x51\x63\x69\xef\xee\x70\xc8\x3e\x62\x46\x15\x07\x82\xf1\x51\xb6\x60\xe1\x79\xe5\x91\xb8\xdb\x9b\x3e\x12\x75\x9d\xc8\x2d\x56\x36\xcd\x21\x54\xb8\x46\xec\x16\x53\xad\xe6\xd6\x51\xc7\x92\x5f\xb8\x0a\x85\xf9\x09\x2f\x5e\x60\x34\xea\x98\xd2\xf0\x35\x4f\x3e\xc8\x8b\xd2\x17\xa9\xf0\xb0\x35\x97\xbd\x15\xb6\xdb\xe3\x02\x73\xbc\xc1\xd7\x5e\xbe\x47\x43\xbe\x3a\xc6\xb7\x4a\xc9\xdf\x9e\xdc\xa7\x25\x63\x35\xc6\xef\x67\x7f\xbc\x3d\xc5\x3b\x8f\x2a\xa7\x1c\xe8\x31\x1a\x5e\x0a\xbd\x69\xe9\x42\xc5\xa1\x19\x84\x23\xa7\x68\x0d\x81\x2b\x90\x5f\xda\xc0\xbe\x91\x51\x2f\x75\xb0\x02\xfb\xa8\xc6\x2a\x52\xc2\xcf\x57\x8a\xae\x5a\x0e\x09\x97\xaf\x2e\xff\x7d\xf1\xee\xc5\xc9\x5e\x29\xff\xe5\xb0\xa0\x30\x54\xd2\xbf\xc7\x76\x7b\xd2\x1d\x2c\xae\x76\x28\x0e\xd9\xa3\x28\xda\x60\x97\xd6\x51\x4d\x06\x45\x21\xba\x54\xec\xe0\x28\x88\x40\xb1\xc4\xf4\x74\x2a\x3f\x4f\x3f\xa1\xa0\x21\xda\xf1\xbe\x0d\x13\x50\xd9\x4b\xc0\xfe\x84\x52\xbd\x5a\x17\xa5\x2e\x52\xc8\x31\xc9\x5e\x44\x4a\xdd\x46\xe4\x16\x35\x79\x5a\xea\x6e\x96\xf2\x24\x26\x5d\x2e\xa0\x23\x22\x0b\xcd\xc7\x0e\xce\x37\x15\xd6\x46\x38\x6d\x8d\x34\x0c\xb3\xb5\x1a\x77\x26\x43\xe8\x6b\x39\x7c\xdc\x9f\x74\x1c\x29\x20\xcd\x6d\xb7\x24\xc3\x7a\xdc\x63\xdc\x70\x20\x35\x96\x54\x22\xaa\xc0\xcd\x0d\xdb\x36\x70\x49\x31\xca\x56\xad\xac\x08\xed\xdc\xb6\xe2\xaf\xcf\x5f\xf5\x69\x44\x42\x9c\x73\x76\xa6\xeb\x31\xfb\x92\xa0\x61\xf4\x1a\xec\xdd\x5a\xaa\x69\xbb\x64\x48\xa0\x18\xd5\x34\xc7\x30\x75\x5c\x6a\xd7\xdd\xb8\xf5\xa7\x48\xa5\x19\x8a\x15\xc1\x9c\xe9\x48\xce\x7a\xd9\x4c\x9c\x3f\x7d\xf3\xa0\x7d\xe4\x2a\xad\x74\xf8\x6a\xfb\xd2\xe9\x46\x2f\x77\xd6\x6a\x0c\xf2\x82\xb4\x8e\x9a\x7a\x3e\xbc\x39\x95\x81\x37\xa3\xba\xa6\xcd\xec\x1b\x1d\x17\x68\x4c\x34\x3b\x5a\x45\x1f\xe7\xe6\xdf\x86\xfd\xf5\x93\xca\x65\xf2\xe9\xcb\xff\x3d\x77\x43\x02\xff\x2f\x77\x7d\x11\xdf\xe7\x4d\x1d\x93\xb6\x94\x3d\x99\x42\x9b\x46\xe0\x51\x09\xa5\x72\x4b\x3e\xce\x6d\x95\x0a\x01\x57\x60\x57\xb4\x4e\x7b\xea\x75\x49\x28\xe2\x81\x53\xc2\x64\xfb\x22\x26\x22\x42\xd0\x2e\x32\x3c\x91\xb9\xb6\x9c\x74\x83\x9d\x2c\xd9\xe5\x86\x22\xe4\xba\xd7\xdf\x08\xcd\x4e\x18\x44\xeb\x7b\x9d\x2f\x45\x0e\x44\x33\x76\xf7\xc2\x06\x4f\xfe\xfe\xec\xc9\x5d\xf7\xc3\x7b\xfc\x4b\x1e\xbd\xfe\x76\xb7\xdf\xb8\x8e\x8e\x6b\x44\x2b\x08\x5f\x11\x1a\xed\x75\x4d\xa0\x25\x85\x75\x9a\x8b\x49\x9a\x07\xce\xf5\x1c\x3b\x09\xdf\x9b\xca\xa0\xe3\x3b\x2f\x77\xce\x8d\xdb\x83\xd7\x6a\x0c\xcf\x89\x4e\xa1\x13\x37\xb6\x2c\x6e\xf6\x0c\x65\x90\x0f\x56\xc7\xdc\x46\x64\x9f\xac\x43\xa3\x63\x77\x49\x8c\xc8\xed\x21\xc6\xee\xf4\xf2\x25\xd8\xf7\x7f\xf3\xc6\x72\x4e\x26\x77\xed\xda\xfb\x1e\x40\xa0\x19\x73\x12\xe2\x28\xb9\x69\xbb\x8b\xee\x5d\x9f\x36\x23\x15\xe7\x39\x19\x91\xf4\xa2\x18\xce\xfc\xfc\x57\xb9\xfe\xbb\x48\xdb\xed\x01\xc6\x8f\x56\x83\xcf\x9f\xfb\x2b\xd2\xee\xeb\xe1\x7f\x01\x00\x00\xff\xff\x38\xb6\xac\x3c\xe3\x0f\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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