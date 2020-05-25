// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../src/contracts/ExampleNFT.cdc (3.833kB)
// ../src/contracts/NonFungibleToken.cdc (4.9kB)

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _examplenftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\x4d\x8f\xdb\x36\x13\x3e\x4b\xbf\x62\xde\x1c\xf2\xca\x68\xd6\x4a\x8b\xa2\x07\x63\xb7\x9b\x62\x37\x06\x7c\xa8\x11\x24\x2e\x7a\x08\x02\x84\x16\xc7\x36\xb1\x12\x69\x90\x23\x3b\xee\xc2\xff\xbd\x18\x52\xdf\x96\xb3\x69\x10\xac\x65\x71\x66\xf8\xcc\x33\x9f\x4e\x53\x58\xed\x94\x03\xe5\x40\x68\xc0\x6f\xa2\xd8\xe7\x08\x8a\xff\x16\xa8\x49\x90\x32\x1a\xcc\x06\x04\xcc\x73\x73\x84\xa5\xd1\x37\xf3\x52\x6f\xd5\x3a\x47\x58\x99\x27\xd4\x71\x9a\xc2\x82\x58\x5f\x1b\x82\xbd\xb0\xc4\xe2\xb4\x43\x30\x9b\x8d\xca\x94\xc8\xc1\x91\xd0\x52\x58\x09\xeb\x92\x40\x11\x08\xe7\xca\x02\x25\x90\x81\x35\xb2\xfe\x01\xed\x09\x9c\x2a\x54\x2e\x2c\xbf\xdd\x99\x23\x14\x42\x9f\x60\x39\x5f\x39\x38\x9a\x32\x97\x2d\x24\x6f\x3b\x33\x16\x61\x53\xea\x8c\xf1\x89\x5c\xd1\x69\x1a\xc7\xaa\xd8\x1b\x4b\x8c\xb1\x86\xe8\x11\xc2\xc6\x9a\x02\xde\x7e\x7b\xfb\x4b\x1c\xef\xcb\x35\x64\x46\x93\x15\x19\xc1\xfb\xe0\xed\x72\xbe\x9a\x5d\x2a\x3d\xc7\x31\x00\x00\x2b\x1c\x3c\x2a\x12\xf9\xa7\x72\xbf\xcf\x4f\x33\xf8\x6b\xa1\xe9\xb7\x5f\x5b\x01\x3c\x30\xae\x87\xca\xee\x42\x2b\x52\x22\x57\xff\xa0\x4c\x26\x03\x99\xbf\x15\xed\xa4\x15\xc7\x44\xc9\xda\xcc\x1b\x8f\x6f\x06\x7f\x48\x69\xd1\xb9\xfb\xa1\xca\x23\xee\x8d\x53\xd4\xd3\x20\xd3\x95\x6f\x14\x2c\x3a\x53\xda\x0c\x61\xd4\xa5\xe9\x62\x39\x5f\xc1\xb3\x97\xae\x35\x72\x24\x68\x0d\xc7\xbd\x33\x76\xbb\x40\x12\x52\x90\x98\xc1\xf3\x27\xb2\x4a\x6f\x67\x10\x3e\xcf\xad\xac\xd2\x8c\x4e\x2b\x5a\x3c\xd6\x86\x26\x9d\x6b\xf8\x9f\xc3\x7c\x33\x55\x12\xee\x20\xc8\x5d\x1e\xd6\x17\xc1\x1d\x3c\x9f\x9b\xe3\xf0\x74\x1e\xf1\xf0\xc1\xe4\x39\xfa\xe8\x8f\x38\xfa\xc1\x9a\x83\x92\x68\xdf\x5c\x1e\x7d\xc4\x0c\xd5\x61\xf4\xa8\x35\xf9\xa1\x5c\xe7\x2a\xeb\xf8\x90\xa6\x20\x55\xc8\x35\x7b\xe2\xfc\x66\x26\x33\xa3\x37\xc6\x16\x4a\x6f\x81\xd8\x80\xeb\x8a\xb3\x00\xd7\x54\x8b\x98\x4e\x7b\x84\xa3\xa2\x1d\x17\xda\xd7\xc0\xd3\x57\x58\x3c\xc2\x46\x61\x2e\x2f\x98\x37\x47\x8d\x92\xf3\x7f\x06\xef\x9e\x83\xf4\x88\xa7\xcb\xf9\x6a\x10\x09\x48\x46\xc9\x6f\xcc\xc1\xed\x4d\x9f\xe1\x2e\xea\x63\x95\x9e\x60\xb1\x30\x07\xf4\x3d\x81\x3d\xf1\x15\x14\xea\xae\xe6\x08\x84\x96\x10\x84\x14\x71\xd1\xfa\x63\x91\xe7\x68\x7b\xbe\x6c\x4a\xdd\x98\x4d\xea\x87\x4e\xa6\xcc\xe0\xdd\x98\x57\x03\x1f\x38\x4f\x3d\xc9\x0c\xbf\xef\xd0\x34\x60\x4d\x9e\xf0\x34\x83\xf6\x82\x09\xdc\xdf\xc3\x5e\x68\x95\x25\xaf\x0a\xe5\x1c\x87\x69\x39\x5f\xbd\x9a\xc4\x3d\xc3\x58\xa8\x41\x55\xfa\x6b\xa6\x4a\xd6\x75\xd9\xdc\x66\xef\xa7\x22\xd4\xdc\xc0\x86\x45\x2a\x2d\x23\xf3\xaa\x57\xa8\x95\xa1\x8c\x81\xc4\x13\xf3\xea\x69\x65\x0a\x85\x94\x3d\x06\x1b\x82\x5d\x27\xe5\xba\x86\x1a\x25\x16\x5f\x3c\xd6\x8a\x4a\x82\xb0\x56\x9c\x2e\xc8\xaf\x2e\x4e\x3c\xb8\x2b\x6c\x0f\x53\xa6\x47\x77\x78\x10\xee\x7f\xf0\xae\xed\x98\xac\x15\x5f\xe8\xb4\xad\x04\xee\x1a\x22\xfb\x62\xec\x81\x94\x1e\xb2\xc6\x63\x65\xbc\xf2\xa1\x53\x63\xc7\x9d\xca\x76\x4d\x1a\xfa\x71\x92\x4b\x30\x1a\x2f\xee\x34\xb9\x5c\x8d\x67\xc6\x67\x25\xbf\x34\x0e\x8c\x84\xbd\xdb\x59\x39\xde\xdc\x55\x5f\x8e\xb6\x44\x47\xd6\x9c\x9a\x7b\xaf\xc4\x7b\x8b\xb4\x78\x74\x55\x6e\xf8\x42\xf2\xe1\xa9\x87\x23\x9f\xd1\x4e\x10\x08\x8b\xa0\xf4\x20\xf6\x17\x41\x0c\xd6\x92\xc9\x0c\x3e\x07\x7e\xbf\x0c\x22\x56\xe5\xe0\xa0\x34\x9e\xf0\xe4\xae\xe0\x5b\x1b\x6b\xcd\x91\xb3\x70\x8b\x14\x1a\xd5\x06\x2d\x6a\xee\x54\xa6\xae\xfb\xeb\xc0\xd2\x14\x9c\x09\x1e\xb4\x85\x0f\x99\xd0\x60\x51\x48\x50\xe4\x9a\xd9\xe1\x33\x96\x05\xea\xb7\x3b\x23\xdd\x85\x87\x0d\x9e\xce\xa0\x9b\xcc\xe0\xf5\x0f\x34\x87\xca\xf7\xd7\x23\xd1\x17\x6e\xdc\xc2\x18\x29\x55\x60\x2f\xfa\x67\x1d\xf0\xbe\xf9\xf1\x19\x95\xa6\xec\x10\x0f\x8f\x7a\x39\xa9\xa2\xac\x4f\x46\xa3\xe7\xc7\x33\x41\x06\x32\x8b\x82\x10\x84\x2f\x03\x2c\xf6\x74\x1a\xf2\x5c\x53\x13\x24\xdf\xb3\x48\x3b\xa3\x92\xd1\xce\xd9\x9e\x77\x9c\x68\xfa\x53\x7d\x67\xd7\xca\x00\xfd\xc7\x66\x5c\x05\xd8\x20\x64\xa1\x34\x18\x0b\xce\x70\xe8\xb8\x8d\xd6\x9b\x5a\x58\xcc\xcc\x51\x57\x9b\x5c\x5d\xde\xbc\x1b\x92\x81\x42\x69\xf2\xce\x35\x74\xa5\x69\x1c\x0d\xf7\x94\x3f\x95\x26\xb4\xbc\x6e\x45\x51\x9a\x7a\x25\x8e\x30\x7f\xba\x8a\x1c\xfe\x1e\x06\xa7\xff\xba\x78\x0c\xa2\x9c\x56\x75\x57\xe5\xff\x21\x57\x2d\x66\x6a\xaf\x90\xb5\x3b\xc3\xaa\xf4\xfd\x9f\x76\xa8\x6c\xf7\x75\x93\xf3\x71\x14\xd5\x6c\x57\x08\x92\xc6\xd0\x0c\x5e\x3f\xbf\xb8\x2c\x9c\x27\xc1\x05\x06\xd6\x8b\x2c\x27\x5b\x14\x45\x3c\xd2\x35\xfa\x7a\x6b\xe3\xe0\xb3\xbd\x5a\x9c\x3a\xad\xb5\xb3\x6b\x4e\x6a\x9b\xdf\x71\xf4\xff\x0e\x44\x96\x99\x52\x53\xcf\xcd\xae\x6f\x51\x23\x3c\x1d\x8c\x83\xdb\x9b\x80\x6a\xd0\xe4\xc6\xd1\xc0\xdd\xb5\x83\x9f\xaa\x8a\x4d\x7e\x9e\xc4\x51\x74\x8e\xa3\x73\x1c\x47\x7e\x31\x9c\xf4\xd7\xa8\x76\x39\xf6\x4e\x78\x1b\xe0\xbc\x91\x46\xcc\x17\x5a\xff\xda\xb7\xbd\xde\xf5\x50\x13\xfc\xd0\x0d\x65\x95\x53\x9c\x16\x4e\x1c\xb0\x1a\xaa\x8e\x8c\x15\xdb\x76\x6c\xf0\xc8\xe8\x64\xc0\x77\x8a\xa2\x81\x52\xb1\x3b\x65\xab\xc9\xed\x4d\xab\x1d\x46\x46\x5a\x5d\x91\x2e\xe7\xab\xd6\xc8\xa4\x87\xb8\x49\x89\xaa\x3b\x64\x62\x2f\xd6\x8a\x7f\xb7\xc0\xc6\xd8\x6b\x4d\xb6\x77\x7b\xae\xf4\xd3\xed\x8f\x64\xe2\xef\x49\x7f\xe0\x86\x2b\xfb\xe8\xde\xf4\x44\x48\xd8\x2d\xd2\x35\x4f\x1a\xd1\xc9\x78\x10\xaa\x0a\xfe\x2f\x01\x28\x82\x4a\xaf\x12\x82\x99\x17\xb8\x0f\x8a\x97\xbc\x07\xe5\x0e\x40\x3f\xe0\xc7\x7f\x91\x45\xe7\xf8\x1c\xc7\xff\x06\x00\x00\xff\xff\x34\xa8\x82\x60\xf9\x0e\x00\x00"

func examplenftCdcBytes() ([]byte, error) {
	return bindataRead(
		_examplenftCdc,
		"ExampleNFT.cdc",
	)
}

func examplenftCdc() (*asset, error) {
	bytes, err := examplenftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleNFT.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x58, 0xfc, 0x1c, 0xd3, 0x92, 0x3, 0x53, 0xaa, 0xfd, 0x92, 0x1c, 0xb5, 0xba, 0x6b, 0xd4, 0x72, 0x60, 0x4f, 0x19, 0xdd, 0xc, 0x6d, 0xd2, 0x9, 0x18, 0x17, 0xbc, 0x3c, 0xc2, 0xc9, 0xd4, 0x12}}
	return a, nil
}

var _nonfungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4d\x8f\x23\xb7\x11\x3d\xbb\x7f\x45\x79\x0d\x64\x67\x16\x1a\x29\x87\x20\x07\x01\x8b\x6c\x62\x79\x00\x5d\x26\xc6\x44\x41\x0e\x86\x01\x51\xcd\x6a\x89\x58\x36\xd9\x26\xd9\x92\x95\xc1\xfc\xf7\xa0\x8a\x64\x37\x5b\xd2\xd8\xe3\x5b\xf6\xb2\x92\x9a\xac\x8f\x57\xaf\x5e\x55\xcf\xe2\xd3\xa7\xaa\xfa\xee\x3b\xd8\x1c\x10\x1e\xb5\x3d\xc1\x93\x35\x0f\x8f\xbd\xd9\xab\x9d\x46\xd8\xd8\xaf\x68\xc0\x07\x61\xa4\x70\x92\x0f\x6e\x9f\xac\xc9\xcf\xf9\xf1\x16\x6a\x6b\x82\x13\x75\x00\x65\x02\xba\x46\xd4\x58\x55\x64\x6f\xf8\x0a\xe1\x20\x02\x08\xad\xc1\x58\xf3\xd0\x64\xeb\x81\xad\xe7\xdb\x1e\x6a\xdb\x6b\x49\xdf\x1b\xeb\x5a\x08\x76\x5e\xad\x1b\x10\xd0\x7b\x74\x70\x12\x26\x78\x08\x16\x24\x76\xda\x9e\x41\x80\xc1\x13\x98\x26\x0c\xf7\x67\x10\x0e\xa8\xdc\x18\xcd\x89\xcd\x19\x44\x59\x05\x0b\xaa\xed\x34\xb6\x68\x02\x1d\x83\xcb\x24\xc6\x58\xe7\x1c\xfb\xb5\x9d\x83\x38\x52\xc4\xd0\x58\x4d\x30\x51\x32\x64\xc8\xf5\x1a\x3d\x08\x23\xc1\x88\x56\x99\x7d\xc5\xa9\x86\x49\xf6\xbe\xc3\x5a\x35\x0a\xfd\x3c\x21\xf8\xb8\xd9\x82\x43\x6f\x7b\x97\xa1\xaa\xad\xc3\xe1\x27\x08\xe7\x2e\x61\xe6\xb0\x73\xe8\x91\x72\x17\x06\x9e\x1e\x37\xa0\x0c\x5b\xf7\xad\x70\x63\xee\xc9\xf0\xf7\x56\x6b\xac\x83\xb2\x66\x0b\xcf\x13\xfb\xa3\x69\xb2\xea\x83\x75\x14\x35\x43\xfb\xd1\xb3\xdd\x7a\xb8\x3b\xaf\xd6\x54\xca\x5a\xf7\x92\x0f\x35\x78\x82\xa6\x37\xfc\x8c\x4b\x20\x18\x01\x8a\xc2\x9e\x0c\x3a\xfa\x09\x85\x57\xfa\x5c\xb5\xf6\x98\xca\xea\x29\x50\x82\xc5\xf6\x01\x6c\xc3\xa7\x4b\x17\x1c\xef\x8f\xce\x1e\x95\x44\xb7\xe5\x93\xdb\x67\xac\x51\x1d\xe9\xeb\x10\xee\x00\xa2\xe7\x3c\x7c\xf9\x0b\x48\xac\xb5\x70\x58\x04\x77\x52\xe1\x00\xde\xb6\x08\x9d\x43\x36\xda\x59\xcf\x30\x49\xc5\x27\xaa\x84\xea\x2f\xbd\x72\xc8\x41\x8d\x98\x15\xd5\xad\xd1\x05\xa1\x4c\xaa\x29\x1b\xda\xe1\x41\x1c\x95\x75\x43\x37\xf8\xc8\x94\x33\x50\x08\x1e\x3b\xe1\x44\x40\xd8\x61\x2d\x7a\x0a\x33\xc0\x5e\x1d\xd1\xb3\x0f\x66\x30\x7d\x10\x3b\xa5\x55\x38\x93\x27\x7f\xa0\x7b\x02\x1c\x36\xe8\xd0\xd4\x48\x24\x8d\x0c\x2e\x43\xa2\x70\xad\xd1\x67\xc0\x5f\x3b\xeb\x93\xbd\x46\xa1\x96\x91\x75\x63\xee\xca\x80\x35\x08\xd6\x41\x6b\x1d\x56\x09\xf3\x11\xae\x39\xac\xa9\x07\xbd\x4d\x81\x51\x50\xfe\x32\xaa\x56\x7c\x45\xa8\x7b\x1f\x6c\x3b\x14\x21\x81\x36\x69\xa0\x69\x21\xa8\x2d\x2d\x1c\x85\x53\xb6\x27\x93\xca\xec\x53\x2d\xc8\x7c\xe4\xc3\xbc\xaa\xfe\x71\x86\xde\x13\x9e\x83\x65\x4e\x61\x34\x34\x4b\x41\xd9\x86\x29\x39\xe5\xb8\x87\x5a\x18\xf0\x68\x64\x45\xb7\x5c\x24\x4b\x66\x5b\x87\xe8\x1e\x82\x7d\xa0\xff\x67\xec\x9b\x88\x47\x25\x33\x7b\x8a\x8f\x9d\x70\x37\x53\x58\x02\x6a\x24\xab\x1a\x34\xca\x3d\xba\xea\xaa\x9d\x36\x96\x5d\xe5\xae\x23\xd6\x1b\x1b\x0e\xe8\x38\xc4\xd9\x20\x4b\xac\x0d\x9e\xb0\x39\xb3\x69\xe9\x44\x6c\x8d\xa7\xc7\x4d\xd5\x38\xdb\x5e\xd5\x94\x75\xca\x40\x9d\x15\x44\x62\x67\xbd\x0a\x43\x25\xc1\x9a\x89\xaf\x8f\xbe\x9a\x72\xb4\xb6\x54\x89\x10\xe9\x1b\x9c\x30\xbe\x41\x37\xaf\xaa\x4f\x8b\xaa\x5a\x2c\x58\xc9\x5b\x22\x6f\xec\xea\x4b\x69\x9e\xc3\x3f\xd9\x74\xf9\x94\x8a\xa5\x35\x5d\x56\x6d\x67\x5d\x88\x65\x29\xea\xad\x7c\xa1\xed\x8b\x45\xd5\xf5\xbb\x1b\xa6\xaf\x55\xf5\xa5\xaa\x00\x00\x52\x54\xc1\x06\xa1\xc1\xf4\xed\x0e\x1d\x6b\x42\x2c\x1d\x33\x55\xf9\xa8\x7a\xca\x00\xfe\xaa\x7c\xe0\x8e\xa0\xbb\xe4\xea\x28\x5c\xbc\xfc\xaf\xbe\xeb\xf4\x79\x09\xff\x5e\x9b\xf0\xd7\xbf\x0c\xc6\x7f\x38\xc6\x30\x45\x00\x6c\x55\x08\x28\xe1\x44\x18\xa7\x3a\x14\xa1\x52\x1e\x2a\x28\xa1\xd5\x7f\x51\xa6\xeb\x83\x1b\x64\x33\xdf\xa7\xc3\xeb\xf1\xe0\xdd\xfd\x2d\x57\xca\x4f\xbd\x89\x34\xd0\x94\x1f\x98\x60\x66\xf9\x9e\x32\x52\xd5\x22\x30\x1b\x07\xe1\xbc\xd2\xc5\x64\x38\xc0\x49\x14\x46\x80\x78\x34\x2f\xa3\x5d\x2c\x60\x7d\x75\x57\x79\x30\x36\x44\xdd\x05\x51\xd7\xb6\x37\xe1\xa3\x67\xb1\x17\x7b\x9c\xc1\x96\xcc\x6c\xb9\xd4\xb0\x43\xd8\x1a\xa5\xb7\xf3\xdb\x18\xfc\x27\xb9\xbe\x53\x32\x83\x3d\xe3\x28\x96\xf0\x77\x29\x1d\x7a\xff\xb7\x9b\x90\xbc\x85\x47\xe2\x38\x4a\x6e\xa4\xc9\x20\xb8\xc8\x2a\x64\xa4\x92\xd4\xbd\x07\xa8\xd2\xfa\x1b\x09\xad\xe2\x91\x49\x3e\xc1\xde\xca\x66\x3d\x5d\x5a\x12\x85\xfc\x30\xff\xc7\xf5\xe4\xd2\xd3\xf5\xd0\x82\x35\xb1\xef\x85\x4f\x14\x7d\xd0\x1b\xf5\x4b\x8f\xb0\x5e\x25\xd0\x44\x7d\x60\x9a\x1e\x84\x1f\x8e\x92\x41\x8d\x01\xc6\x80\xf9\xd1\xeb\x10\xe7\x73\x9c\x61\xed\x80\x3d\xe9\x49\x0a\x8e\x58\x76\x4b\x40\x29\x87\x7c\x9f\x57\xa9\x46\x99\x38\x83\x52\xe4\x24\x4a\x28\xa3\xe2\x91\xcd\x64\x8f\x15\x9e\x72\xb9\xce\xf5\xe9\x71\xb3\xbc\x4c\xf3\x77\x63\x2f\x30\xb6\xd0\xa2\x54\x34\x39\x33\xdd\x3d\x64\xd9\x2c\x44\xf3\x1d\x58\xe7\x65\x62\x8a\xf7\xa0\xc9\x0e\x69\x39\x19\xd6\xa8\xc1\x47\xc1\x29\x52\xbd\x78\x48\x05\x88\xd3\x38\x22\xe2\x26\xa9\x35\xbd\x19\xcc\xde\xe5\x0f\xeb\x55\xce\xf5\x7e\x09\x5f\xa6\x78\xf0\x45\xda\x43\xa6\x3f\xd1\x3f\x87\xbe\xd7\x61\xae\x24\x7c\xfe\x0c\xa5\xad\x0f\x44\x94\xf5\x2a\x33\x7f\xd4\x82\xd8\x53\x6d\xef\x03\x35\x31\xaf\x82\xa2\x45\x10\xb1\x5d\x68\xb3\x41\x4f\xad\xb0\x5e\x7d\x98\x78\x7b\xad\xa6\x9f\x7e\xa7\x1a\xa9\xa7\x7c\xc6\xe1\x0f\x95\x22\x2f\x72\x59\xff\x93\xa3\x3c\xe9\x82\xf8\x3a\x16\x42\xf0\x27\xe1\xf6\x3d\x53\x99\x6a\x20\xa4\x2c\x4b\x70\xe1\x3a\xb9\xff\xe6\x9b\x5c\x8b\x64\xf6\x8e\x91\x89\xe0\xdf\xbf\x9d\x22\xb7\xca\xa0\x8f\x69\x80\xd7\xb6\x6d\x79\xcb\xca\x17\xba\x7e\xa7\x95\x3f\x40\x63\xdd\xf0\x5a\x31\x89\xe2\x8d\xcc\xc7\x58\x7f\x24\x0b\xf5\x45\x57\xfc\x66\xb8\xe5\xa1\x3d\x86\xf5\xca\xdf\xdd\x2f\xe1\xa7\xc8\xaa\x9f\xaf\x8e\xec\xac\x73\xf6\xf4\xf4\xb8\x29\x44\xed\x7e\x09\x7f\xca\x6d\x7a\x5b\x2a\x52\x42\x89\xfa\xa6\x76\xb4\x48\x4c\x5e\x3c\x0a\x81\xd8\x61\xde\xb1\x65\x7e\xef\x18\xb6\x02\xd2\x98\xac\x2c\x6f\x52\x62\x84\x63\x39\xf4\xe7\x6c\xa0\xc7\xec\x16\x5c\x25\x61\x56\x8a\x9f\x09\xc7\xbb\xe9\xc1\x6a\x39\xea\x71\x8a\xe7\x06\x39\xf2\xc6\x40\xa3\x43\xd2\xd9\x25\x7c\x79\x89\xf8\x2c\xe9\xee\x6b\xf5\x7f\x21\x10\xbf\xd5\x1a\xb1\x33\xae\x5b\x61\x8c\xc5\x83\x1c\xc0\x29\x0d\x0d\x97\x42\xd4\x8f\x74\x51\x49\x10\xce\x89\xf3\xfb\xd8\x58\x1a\x8c\x4c\x04\x87\xa1\x77\x26\xf5\xaa\x13\xe7\x2c\x4c\xf4\x2c\xf6\x94\xc3\x5c\x93\xfa\x76\x4d\xde\xe0\x75\xe9\xec\x39\x7b\x49\xec\x46\x39\xbe\x1f\xc5\x1d\xbc\x7c\x07\xbe\xe1\x67\xb1\x00\x6f\xc7\xc9\x1d\x8b\xc3\x2f\x0e\x0e\x85\x04\x29\x82\x60\x88\x78\xfb\x6e\x31\x1c\xac\x4c\xf3\x46\x85\x3f\xd2\x61\x97\xea\xee\xf0\x86\xb8\x7b\xd4\xcd\x7c\x60\xe1\x4f\x4a\xfe\x0c\xdf\x7e\x06\xa3\xf4\x12\x3e\x90\x0d\x69\x31\xae\x6c\xbc\xf1\x5e\x67\xf5\xed\x7b\x15\xbc\x76\x28\x02\xfe\xd0\x76\xe1\x5c\xbc\x2a\xc4\x5f\xb9\x64\x48\x8f\xae\x35\x1c\xe2\x8b\x54\xc4\xfc\x92\xd2\x25\x90\x67\x86\xd0\x9e\x18\x7e\x5f\x95\x20\xdd\xf4\x4d\x05\xfe\x52\x84\x52\xa8\xe0\xf5\x1c\x4c\x33\x30\x53\x63\xae\xd1\xec\xc3\x81\x06\xe2\x9f\xd3\x1c\x8c\x3e\x64\xd9\x8a\x79\x00\x72\x66\x05\x50\x19\x9a\x57\x7e\x17\xca\xbb\xfc\xf4\xaf\x16\xaa\x1c\x78\xac\x71\x9d\xb6\x67\x94\xd3\xd7\x9a\x55\xdf\xb6\x67\x78\x79\xad\xe0\x7f\x01\x00\x00\xff\xff\x48\xa4\xe2\x85\x24\x13\x00\x00"

func nonfungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_nonfungibletokenCdc,
		"NonFungibleToken.cdc",
	)
}

func nonfungibletokenCdc() (*asset, error) {
	bytes, err := nonfungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NonFungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x8b, 0x74, 0xc1, 0xe1, 0xf0, 0x96, 0xd9, 0x24, 0xfc, 0x1, 0x2, 0x10, 0x8, 0xde, 0x90, 0x9e, 0x67, 0x30, 0x74, 0x11, 0x57, 0x50, 0x84, 0x51, 0x3e, 0x49, 0xe5, 0xad, 0xa0, 0x6f, 0x81}}
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
	"ExampleNFT.cdc":       examplenftCdc,
	"NonFungibleToken.cdc": nonfungibletokenCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"ExampleNFT.cdc":       &bintree{examplenftCdc, map[string]*bintree{}},
	"NonFungibleToken.cdc": &bintree{nonfungibletokenCdc, map[string]*bintree{}},
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
