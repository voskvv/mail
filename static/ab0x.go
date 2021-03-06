// Code generated by fileb0x at "2018-04-27 13:25:42.422197217 +0300 MSK m=+0.007172437" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(432a5369214f19739586a840cb1259d1.3daf2b4d0cc8675651cd8e10adf6cbc7)

package static

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path"

	"context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileSwaggerJSON is "swagger.json"
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5c\xcd\x6e\xdc\x38\x12\xbe\xe7\x29\x08\xcd\x02\x3e\xac\xbb\x95\x9d\x5d\xec\xc1\xb7\xc1\xb6\x27\x31\x66\x66\xe1\xb5\x9d\x5c\xc6\x86\x41\x8b\xd5\x6a\x26\x12\x29\x93\x94\x63\x8f\xe1\x77\x5f\x90\xd4\xbf\xa8\x96\xfa\x4f\xed\x4e\x9c\x4b\xd4\x52\xb1\x48\x56\x7d\x5f\x55\x89\xa4\xfc\xfc\x0e\x21\x2f\xe0\x4c\xa6\x31\x48\xef\x04\xfd\xf9\x0e\x21\x84\x3c\x9c\x24\x11\x0d\xb0\xa2\x9c\xf9\x5f\x24\x67\xde\x3b\x84\x6e\x8e\xb5\x6c\x22\x38\x49\x83\x61\xb2\xf2\x1b\x0e\x43\x10\xde\x09\xf2\xfe\x39\x7d\xef\x99\x7b\x94\xcd\xb9\x77\x82\x9e\x6d\x5b\x02\x32\x10\x34\xd1\x6d\xb5\xd4\x1f\x98\x46\x93\x2b\x88\x93\x08\x2b\x10\x48\x82\x78\xa0\x01\x20\x2a\x8b\xcb\x39\xd7\xb7\x19\xa1\x2c\x44\x10\x63\x1a\xc9\xa9\xd1\x8b\x90\xa7\xa8\x8a\xa0\xad\x45\x0f\xe7\xc5\x0e\x1d\xab\x85\x2c\xfb\xf6\x63\x90\x12\x87\x50\xde\x42\xc8\x0b\x41\x55\x7e\xb6\x47\xb8\x50\x2a\x91\x27\xbe\x1f\x2c\xa6\x89\x6e\x3b\x0d\x38\x53\x98\x32\x10\x69\x3c\x65\xa0\x7c\x9c\xd0\x09\xe1\x81\xf4\x63\x4e\xd2\x08\xa4\x1f\x2c\x26\x7a\x9c\x13\x95\x0d\xc8\xa7\x8c\xc0\xe3\x74\xa1\xe2\xe8\xa7\x10\xd4\xa4\x18\xc5\x71\xd9\xa9\xc2\x61\x69\xe0\xec\xde\x1f\xb9\x5c\x71\xf3\xa6\xd2\x42\xa6\x71\x8c\xc5\x93\x1e\xe2\x07\x50\x28\x57\x8a\x22\x2a\xd5\xb4\xaa\x9a\x27\x20\x8c\xaf\xce\x88\xb1\x95\x15\xfc\x9d\x4a\xf5\x01\xd4\x47\xcc\x48\x04\xa2\x2a\x9f\x60\x81\x63\x50\x20\x9a\x03\x7a\xae\x5c\x23\xe4\xfd\x4d\xc0\x5c\x2b\xfc\xc9\x2f\x5b\xf8\x9f\x24\x88\x0b\x1e\xc1\x47\xc0\xc4\x7a\x22\xff\xf7\x72\xdc\xad\x4a\x3d\x25\xc6\x8f\x52\x09\xca\x42\xef\xb8\xfe\x94\xe1\xd8\x3c\xd5\xe6\x6f\x3e\xa3\xc6\x47\xf7\x29\x88\xa7\x2d\x77\x06\xe2\x76\x95\x0e\x9d\x3e\x12\x20\x13\xce\x64\x0d\x70\xe6\xc1\xcf\xef\xdf\x37\x6e\xb5\x81\x97\x79\xd4\x38\x14\x85\xa0\x50\xae\xad\x39\x24\x19\x2c\x20\xc6\x2d\x7d\x35\x17\x11\x98\x53\x46\xb5\x6a\xe9\x57\x10\x70\x91\xab\xac\xb5\x7c\xe9\xb2\xa4\x47\x60\x8e\xd3\x48\xb5\xc7\x5e\xf6\x54\xcc\xd9\x07\x21\xb8\x70\x1b\xa9\xa2\xd5\x7b\x9c\xc4\xa0\x16\x9c\x4c\x1e\xa8\xa4\x77\x34\xa2\xca\x80\x3a\x11\xf4\x01\xab\x62\x64\xb6\x6d\xd6\xae\x64\xb2\xff\x9c\x5d\xdd\x52\xf2\xf2\x2a\x68\x3d\x09\x78\xf2\xb4\x55\x6a\x4b\xca\xc2\x08\x72\x86\x0f\xe0\xf6\x81\xf1\xba\xf4\xa0\x9b\x6c\x3a\x8a\x37\x9f\x08\xb8\x4f\xa9\x00\x3d\x6b\x25\x52\x18\x87\x89\x3b\x20\xe1\x07\x38\x5c\x0e\xea\xac\x5c\xa5\x5c\xc2\xe5\x98\x9c\xd3\xdd\x17\xa4\x9b\x73\x31\x11\x20\x79\x2a\x02\x98\xc4\x98\xe1\xb0\x0e\x7e\x17\x09\x2f\xf5\xf8\xfb\x08\xa8\x85\x72\xea\xa1\xb9\xe0\x31\xa2\x4c\x81\x60\x38\xca\x2b\x94\x65\x84\xbc\xa4\x71\x12\x81\xd6\xb1\x01\x21\x73\x9e\xdc\x71\xf2\xe4\x66\x88\xeb\xc9\xaa\x80\x2c\x87\x7a\x01\xf7\x29\x48\xb5\x04\x8f\xab\x52\xec\xe7\xc1\x14\x93\x66\x14\xa6\xe0\xdb\x1a\xd5\xaa\x33\x7b\xbd\x54\x4b\xef\x22\x1a\x74\x30\x2d\x87\xff\x9e\x0b\x57\x1c\x95\xcf\x7a\xab\xd7\xab\x42\x70\x48\x8e\x2b\xd4\xf6\xd6\xaf\xb9\xde\xf1\x0b\xd8\xed\x27\x96\xfa\xac\xb7\x9a\x5f\x0a\xf3\x1f\x58\x99\x77\xbc\xa7\x6c\x12\x08\xc0\x0a\x8a\x47\xdb\x43\xf7\x7f\x8c\x62\xc4\xe0\x5b\x01\xf2\x21\xf0\xb6\xcd\x5e\x47\x19\x37\x4e\x02\xca\x67\xbe\xc5\xbc\xf3\x8f\x5e\x06\x5a\xb7\x13\xe4\xf0\xfb\xae\x26\xf1\x3a\x8b\xb9\x22\x14\xf9\xcf\xda\xdd\x2f\xfb\xcf\x34\x0f\x20\xa4\xb6\xe9\x84\xcf\x0b\xb9\x6b\xb6\xfd\xbe\xf2\xfb\xba\x1f\x99\x40\x40\xe7\x34\xc8\x3b\xdf\x6e\x92\xcb\x5e\xe4\x56\x09\x03\x07\xf6\x2a\xd7\x36\xdb\x6e\x57\x69\xcc\xff\xaf\xf3\xa5\xf1\x2d\xb7\x2f\xc9\xed\xe9\x98\x01\x25\x4d\x48\x35\xb5\xef\x96\xe7\x9f\x4c\x67\x07\x9b\xea\x47\xa6\xdd\xf7\x58\x74\xf4\xbf\xec\x5a\x40\xfe\x90\x45\x47\x11\x02\x08\x44\xa0\xe0\x7b\x8d\x02\x33\x33\xbb\x95\xa2\x80\x6d\xf2\x16\x05\x46\x28\x3b\xb6\xcf\xe8\xdc\xd1\xc8\xc2\x9a\x78\x07\x55\xf0\xe7\x97\x2f\xaf\x65\x35\x17\x1e\x15\x30\x02\x64\xfb\x2b\xb8\x8a\x23\xcc\x9e\xec\x46\xf2\xd2\x85\xdb\xcd\x96\x6c\x87\xc1\xb6\x2b\xfe\x1f\x66\x2a\xdd\xfb\xd2\xf1\x56\xd7\x8c\x0f\x70\xb5\xb8\x38\xf6\x50\x99\x48\x79\xf8\xe1\x54\x88\x2a\xbd\x1b\x46\x3c\x15\x02\x4d\x26\x48\x2a\xcc\x08\x16\x04\x49\x10\x14\x47\xf4\x2f\x7c\x17\x01\xfa\xe5\xfc\x0c\x99\xf1\x5e\xb3\x6c\xdb\x4a\xcb\x06\x9c\x69\x71\x65\x1f\xe5\x0c\x3b\xb9\x66\x7f\x47\xd7\x1e\x65\x0f\x38\xa2\x04\xa5\x12\x84\x86\xdf\xb5\x67\xef\xdf\xa7\x5c\x61\x04\x8f\x01\x00\x01\x92\xdf\x35\xb2\x86\x7d\x59\x3f\xde\x35\x9b\x4e\xa7\xa0\x82\xe9\x74\x7a\xcd\xce\x66\xba\xbf\x94\xd1\xfb\x14\xb2\xde\x28\x01\xa6\x74\xee\xb6\xad\x02\x4e\xe0\x9a\xcd\x40\x61\x1a\x49\x2d\xcc\xcd\xcc\x70\xa4\x47\xa9\xe0\xb1\x31\x48\x89\xbe\x52\x46\xb0\xed\x7c\x4e\x21\x22\xe8\xe8\x02\xcc\xa1\x16\x79\x84\xe2\x54\x2a\x74\x07\x88\x71\x36\xf9\x0b\x04\x47\x0f\x38\x4a\x8b\x19\x30\xae\x10\x30\x9e\x86\x0b\xa4\x68\xb8\x50\x52\xc7\x94\x39\x00\x41\x21\x4f\x16\x20\x72\xb9\x7c\x37\x0a\x1d\x7d\xe0\xe4\x08\x11\x0e\xf2\x48\x21\x78\xa4\x52\x69\x91\x5f\x75\xaf\xf5\xa1\x4a\x50\x88\xcf\xd1\x57\x78\x9a\x98\x1e\x51\x82\xa9\x28\xd7\xd9\x8b\xa0\xc2\xef\xbe\x40\xa0\xca\xfb\x89\xd0\xc1\x4b\xd1\x06\x7d\x3c\x62\xcd\xd1\xe4\x54\xae\x06\x0b\x81\xeb\xa4\xf7\xa8\x82\x58\xb6\xf1\xdb\x88\x66\x9d\x04\x78\x9c\x84\x7c\x92\x07\x9b\xcc\x19\x9e\x13\xd6\xc6\xe6\xad\x91\xb9\xa9\x68\x2d\xe5\xd6\x43\xc9\x30\x1d\xa7\x42\x9c\xcd\xdc\x2a\x32\x48\x74\x59\xc9\x11\xc1\xeb\xf3\xcc\x18\xe1\x56\x2e\x15\x56\xa9\xbc\xd5\xa9\xb2\xab\x03\xca\x14\xd4\x37\x2b\xb5\x7d\xb8\x88\xb1\xca\x1e\xff\xfb\x5f\x4b\xba\xbf\x34\x3d\x7c\xbc\xba\x3a\xaf\x8c\xa0\x55\x6d\x9b\x26\x09\x0e\xbe\xda\x99\x7a\x21\x55\xad\x74\x1d\x2c\xfc\x5a\x76\x16\xfe\x03\x30\xc2\x85\x1f\x52\xb5\x48\xef\xa6\x01\x8f\xfd\x4a\x1b\x3f\x58\x80\xc8\x0b\xab\xbc\xaa\xb0\x76\x5e\x1a\x66\x2c\x99\x05\x24\x02\x24\x30\x25\xeb\xbc\x3e\x9b\xad\x8f\x77\xcd\xe9\xc1\x70\xf8\x8d\x56\x4b\x87\x9a\xcf\x56\x00\xd5\x65\x0d\x55\xfb\x34\xfb\x6f\xf5\xc9\x3b\x0c\xaf\x25\x1a\xa6\xd7\x16\xd3\x21\xc7\x26\xa3\x96\xe1\x9b\xd0\xac\xc0\x32\xad\xe3\x72\xcc\x99\x5e\xf6\x21\xec\xb2\x05\xb1\xfc\x2c\xe1\xd9\x6c\xc9\x6c\x1b\x4c\x1f\x6d\x4a\xbf\x36\x63\x61\x73\x4a\x65\xaa\xa8\x4c\x09\x13\x42\xb3\xbc\x51\x26\x0c\x1b\x56\x97\xcc\xb1\x49\xa5\x52\xcb\x79\x07\xa9\xdc\xb1\x7f\x7c\x88\x3b\xce\xcb\x74\x5b\xac\x2d\x6c\xad\x97\x5d\x2b\x6e\x96\x63\xf3\xc2\x51\x64\x45\xeb\xda\x81\x27\xdb\x44\xba\xc5\x6a\x95\x2c\x52\x92\xc9\x2e\x4c\xd0\xc6\xdb\x74\x3d\xce\xdb\x45\x42\xf2\x8b\x1a\x9a\x08\x07\x27\xb0\x33\xb2\x87\xc4\x98\x43\xe1\x36\x93\x5e\xaf\x8b\x7c\xed\xe4\xbf\xfa\xb7\xb3\x1f\x5d\x7f\xde\x6e\x60\x9c\x4f\x12\x44\x97\x81\x1e\xb0\xa0\xba\x3e\xee\xac\xb0\x1a\xf8\x19\xc6\x38\x97\x86\x61\x15\xd7\xe7\x62\x3c\x5b\x4f\x49\xc9\xd7\x50\xbf\xc5\x43\x5e\xcf\x35\x58\xf9\x3b\x95\xea\x94\x29\xf3\xbe\xdd\xc3\xc9\x42\x54\x33\xd2\xa8\x34\x27\xbc\x6b\xc7\x97\x7f\x58\x26\xce\xbe\x03\xb2\xcc\xf6\x82\xbe\xff\x99\x45\xbe\x41\xe8\x33\xa2\x1a\x7d\x66\x61\xb0\x8a\xbe\x0d\xc1\x77\xbe\x24\x5a\x0e\xaa\xf2\x9d\x4e\x39\x07\xb1\x45\xc5\xe3\x79\x64\x78\x9a\x1e\x12\x12\xd0\x37\xaa\x16\xd6\x63\xdf\xb7\x83\xda\x09\xf8\xd6\xa0\x72\xe3\x97\xf8\xde\x6f\x01\x6c\x00\x1f\x96\x6a\xda\xe7\xd7\x47\x42\x96\xbc\x54\x5c\xe0\x10\x3e\xeb\x7a\xb7\x1f\x5a\x35\x71\x27\xbc\x10\x65\x48\x5a\xa1\x83\x4e\x3b\x6f\xd5\xda\x5b\xb5\xd6\xc7\xa1\x0b\x08\x68\x42\x81\xa9\x25\xc4\x29\x64\xec\x1b\x53\xfe\x83\xb2\x39\x2f\x3e\xc5\x43\xba\x3f\x64\x17\xa6\xfb\x49\x53\xd9\xb3\xf8\xb3\x56\x26\x55\x9c\xd3\xd8\x4a\xf4\xcc\xfe\x4c\x6e\x98\x9b\x5e\xfe\x59\xf1\x75\x51\x73\x5a\xed\x6c\xe7\xc5\xdc\x46\x14\xea\xa6\xce\x1b\xbc\xbd\xea\x06\x54\x37\xc0\xdb\x5f\x03\x94\x6b\x03\x81\x59\x1b\x28\x21\x9e\x4a\xca\xc2\xec\xf7\x40\xb4\x77\x2f\xc6\x47\xf8\x69\x47\x6b\xc0\x33\xa3\x7b\x9d\xb4\xe0\x02\x85\x93\xaf\xd9\x93\x2c\x1c\xdc\x12\xac\x70\x15\x08\x37\x35\x0d\xc9\x12\x30\x05\x3c\x8e\x39\xbb\xed\x82\x6b\xdf\xe8\x86\xc3\xb6\x67\xe7\xa2\x05\x5f\x47\xc6\x35\x43\x75\x00\xd9\xd1\xb6\x69\x9b\x25\xb3\x6a\x97\x6c\x4b\xca\xb6\x25\xa5\x5b\x19\xcd\x57\x9b\x56\xd1\x4e\x0e\xdd\xd8\xec\x29\x14\x46\xa1\x75\xef\xcb\xc4\xe8\x8c\x36\x49\x67\xa7\xc5\xb9\x9e\x92\xdd\xdf\x19\x18\x74\xad\xf0\xd8\x31\x37\x1b\xe2\x72\xd7\x58\xa1\x6c\x8f\x59\x5f\xf1\xb9\x76\x83\xca\xbe\xe9\x5f\xdf\x0d\x25\xf1\x36\x48\xd6\x05\x27\xba\xb2\xb6\x6c\xce\x71\x35\xfd\x4d\x2f\x8e\x56\x5c\xef\xda\xfd\xad\x8f\x06\x87\x25\x5e\x23\x6b\x59\x6a\x2f\x5d\x24\xad\x7c\x0a\xb8\x59\xad\xe9\xfa\xa8\x27\x7f\xd5\x70\x96\x50\x83\x0b\xcf\x42\xf3\xa6\x6e\x1b\x69\x41\xee\xad\x60\xec\x42\x6f\x7f\x7e\x59\xbd\x6e\x5c\x03\xc0\x5d\x40\x3b\xb4\xa5\xd9\xab\x36\x31\x9a\x06\xcd\x45\xb4\x19\x8b\xf3\x93\x46\xd7\xba\x3c\x6f\xbe\x45\x3a\x8e\xf5\x9a\xe2\xac\x76\x4e\xb0\x86\xd1\x9b\x43\x58\xea\x71\x14\x98\xc3\x81\x30\xab\x55\xee\x23\xbc\x98\xe6\x16\x5e\x3b\x75\xa6\x8d\x28\xf2\xe2\xf0\xf0\xba\xca\x3f\x67\xed\x47\xe7\x85\x1c\xb2\x65\xd6\x16\xae\x2f\x61\x16\xac\xd9\x6c\xdf\x62\x57\x2b\x12\xd9\xf7\x7b\x7b\x3a\x7f\xf6\x39\xef\x7e\x3f\xbe\x1d\x90\x51\x9c\xf2\xd5\x68\x98\x6d\x7d\xb4\xce\xb4\xae\xec\x61\xd7\xdf\x0f\xd8\xea\x9b\x8a\x03\xd6\xc3\xbc\xe4\xf8\xa4\x62\x67\x6e\xaa\xfc\x59\xb0\xca\x29\xee\xec\x78\x6c\xe3\x43\x89\x8a\xd7\xf2\x19\x2f\x3f\xb3\x92\x2a\x1a\x49\x7f\xa1\x54\xa2\xaf\xfc\xec\x0f\xa1\x4d\xbf\x48\xce\x7a\x3f\xc8\x28\x47\xd6\x3e\x08\xed\xd9\x93\x3c\xdd\x28\xb2\x87\x65\x9a\x07\x7e\x1c\xe7\x9d\x07\x4e\xc4\xea\x6b\x4c\xa0\x71\xec\xcd\x71\xf0\xf8\xdd\xcb\xff\x03\x00\x00\xff\xff\xee\xa2\xd9\xbc\x59\x4e\x00\x00")

// FileVendorGithubComContainerumCherrySwaggerJSON is "vendor/github.com/containerum/cherry/swagger.json"
var FileVendorGithubComContainerumCherrySwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x54\x41\x6f\xdb\x3c\x0c\xbd\xe7\x57\x10\xfe\x3e\x20\x87\x2d\xce\x30\x0c\x3b\xf4\x36\x20\xed\x16\x14\x03\x8a\xa5\x47\x03\x83\x62\xd1\x36\x57\x5b\x72\x29\xba\x4b\x5a\xf4\xbf\x0f\xb2\xdc\xc6\x76\x93\xac\x5b\x76\xf0\xc1\xd4\x23\xf9\x48\x3e\xf2\x61\x02\x10\xb9\x9f\x2a\xcf\x91\xa3\x33\x88\xde\xc7\xef\xa2\xb7\xde\x56\x2b\x29\x5c\x74\x06\x0f\x8f\xed\xaf\xc6\x8c\x0c\x09\x59\xd3\x1a\x27\x00\x00\xd1\x39\xf3\xf3\x4f\x8b\x71\x29\x53\xed\x41\x3e\xd4\x39\x33\xcc\x66\xe0\x44\x19\xad\x58\x83\x43\x26\x55\xd2\xbd\x5a\x97\x08\x9f\xae\x96\x80\xcc\x96\x13\xf3\x15\x9d\x53\x39\x7a\x6c\x6a\x8d\x87\x4b\x78\x82\x2a\xbc\x9c\x25\xe6\x0d\x24\x11\x99\x3b\x55\x92\x86\xc6\x21\x1b\x55\x61\x12\x05\xfb\x6d\x63\x45\x01\x6e\x52\x44\x8d\xfa\xc9\xda\x62\x95\xe7\xd2\xe5\x89\x12\x13\xc7\x31\x4a\x1a\xc7\x71\x62\x96\x0b\x9f\xaf\x31\x74\xdb\x60\x97\x8d\x34\x1a\xa1\x8c\xd2\xe0\x95\x5a\x8d\x89\x59\xa0\x28\x2a\x9d\x07\xdb\xb6\x32\x55\x7a\x96\x82\x9b\x11\x49\x07\x37\x64\xb4\x0a\xc9\x33\xc2\x52\xc3\xf4\x1b\xd6\x25\xa5\xca\x4d\xa1\x6a\x9c\xc0\x1a\xc1\x58\x33\xbb\x47\xb6\x70\xa7\xca\xe6\xb9\x02\x63\x05\xd0\xd8\x26\x2f\x40\x28\x2f\xc4\x81\x58\xc8\x10\x35\xe4\xb6\x2e\x90\x9f\x70\x8c\xce\x36\x9c\x22\x4c\x3f\x5b\x3d\x05\x6d\xd1\x4d\x05\x70\x43\x4e\x3c\xe4\xc2\x67\x1d\x52\x75\x28\x60\x33\xb8\xc1\xed\xac\xcd\x08\xb5\x22\x76\xed\x84\xdb\x91\xc9\xb6\x46\x3f\x2b\xbb\xfe\x81\xa9\xec\xec\x35\xdb\x1a\x59\x08\x5d\x6f\xc0\xed\x88\xdb\x76\x0c\x8c\xbd\x30\x8a\x59\x6d\x9f\xa3\xb4\x4f\x24\x58\x8d\xf1\x3d\x0f\x27\x4c\x26\x8f\x7a\x8f\x8f\x03\xf7\xcd\x2c\xb7\x33\x3f\x6d\x0f\xee\x86\xb1\x43\xf7\xb0\xa1\xe7\x2f\x98\xfd\xcf\x98\x79\xd7\xff\xe6\x3d\x05\xcf\x43\xa7\xf6\xc7\x21\xfd\xba\x18\xe7\xcc\xcb\xc5\xfe\x10\x9d\x24\x0e\x75\xa9\xab\xf9\x70\x9d\xdd\x46\xec\x0f\xee\x44\x49\xe3\xbe\x17\x22\xf5\xa1\x04\x64\x04\xfd\x3e\x0f\x32\x64\x96\x2b\x25\xdd\xf3\xc7\x0f\x47\xd2\xaf\xda\x0c\x5f\xae\xaf\xaf\x7a\x0c\x26\x23\x26\xc1\xa5\x56\xe9\x4d\xa8\x34\xca\x49\x8a\x66\x1d\xa7\xb6\x9a\xfb\x05\x51\x64\x90\x9b\x6a\x9e\x16\xc8\xbc\x0d\x81\x3a\xe7\x28\xb4\xee\xe8\xe5\x08\xfb\xc9\x58\x33\x3a\x34\xe2\x86\xab\xba\x5c\xfc\xbd\x84\xfd\x9a\xbe\x7a\xc2\x97\x1e\xbc\x7f\x0c\x7f\xa0\x93\xd5\x40\x28\xff\xb8\x93\x97\xc3\x7a\xf6\xf4\xd2\x23\x46\xdd\xf4\x4d\xf0\x87\xa1\x6d\xe7\xcb\x5e\x8e\x05\xd4\x13\x4f\x33\x54\xcf\x89\xe4\x57\xbf\xd3\xc1\xea\x85\x10\x1c\xf2\x1d\xa5\x08\xcb\xc5\x91\x02\x46\x2b\x76\x0a\xcb\x8b\xf1\x5d\x19\xb3\xdc\x9d\xdd\x1e\x4b\xa5\x35\x75\x37\x78\x77\x7c\xc3\x89\x3a\x42\x7b\xac\xe1\x5d\x94\xab\x03\x6a\xde\x7f\x47\x4f\xd3\xd6\xc4\x7f\x8f\xbf\x02\x00\x00\xff\xff\x85\xa0\x34\xb4\x1a\x08\x00\x00")

// FileVendorGithubComContainerumUtilsHttputilSwaggerJSON is "vendor/github.com/containerum/utils/httputil/swagger.json"
var FileVendorGithubComContainerumUtilsHttputilSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x95\xcf\x6a\xe3\x30\x10\xc6\xef\x79\x8a\x41\xe7\xf5\xb2\xb0\x37\xdf\x96\xf5\x21\xbe\x2c\x21\x9b\x96\x42\xe9\x61\x62\x8f\x6d\x51\x5b\x72\xf4\x87\x10\x4a\xde\xbd\x48\x72\x9d\x86\x94\xc8\x4d\x71\x0f\x21\x68\xe6\xf3\xcc\x4f\x9f\xc6\xf2\xcb\x02\x80\xe9\x3d\xd6\x35\x29\x96\x02\xfb\xfd\xf3\x17\xfb\xe1\x62\x3d\x2a\xec\xc8\x90\xd2\x2c\x05\xa7\x02\x60\xff\xed\x56\x1b\x6e\xac\xa1\xf2\x4e\x93\xca\xb3\x31\x05\xc0\x04\x76\xe4\x2a\x58\x4d\x2a\xe1\xa5\xaf\xe2\x13\x5c\xb8\xf0\xce\x92\x3a\x9c\x82\x8a\x76\x96\x2b\x2a\x59\x0a\x15\xb6\x9a\xc6\x84\x39\xf4\xbe\x8c\x36\x8a\x8b\xfa\xf4\x40\x25\x55\x87\xc6\x37\xb0\xef\xab\x97\xa4\x0b\xc5\x7b\xc3\xa5\x6f\xb3\x69\x08\xf2\x0c\x64\x05\xa6\xa1\xb7\x3f\x87\x04\x52\xc0\x96\x1a\x6c\x2b\x17\xdd\x37\xbc\x68\x7c\x0e\xcb\x8e\x0b\xae\x8d\x42\x23\x15\x14\x28\xa0\x27\xe5\x9a\x85\x6c\xe1\x0b\xfb\x66\xc7\xd0\x93\xad\x69\x67\x49\x9b\x3c\x5b\x12\x96\xde\xb5\x0b\x0f\x1e\x92\x41\x94\xe4\x19\xbb\x79\x6b\xc1\xb8\x26\x74\xf9\xc8\x39\xa3\x2c\x9d\xa1\x85\x53\xb9\xc6\xe5\x14\xdf\x0c\x35\x58\xf1\x0f\x3b\x9a\xe2\x98\xd3\xc5\xf1\x6e\xf3\xe6\x6f\xcb\x49\x98\xa8\x3f\x41\x36\x17\xc4\x9f\x7a\x0a\x83\x57\xcd\x81\x10\x36\x97\xaf\xae\x11\x04\x4d\x92\xaf\x6e\x05\x38\x4d\x0f\xef\xa7\x61\x6d\xe4\x33\x89\xeb\xb3\xeb\x25\x93\x86\x37\x06\x75\x3e\xd2\xb1\x03\x5b\xcb\x96\xa2\xe7\xe5\x44\x71\x30\x12\xb6\x63\x29\x3c\x0e\x6b\x08\xd7\xe5\x98\x07\x60\xfe\x3e\x62\xc3\xfa\xe9\x8b\xa3\xe6\xde\x25\xdd\x63\x41\x3a\xca\x3f\x4a\xe7\x1a\xfb\x7b\xd9\xda\x6e\x02\x48\xd0\xcd\x45\xb1\xe4\x25\x65\x68\x30\x8a\xe1\x84\x89\x53\x5e\x92\x6c\xa5\x6c\x09\xc5\x27\x50\xfc\x27\x2e\xb0\x2c\xdc\xef\xf8\x1a\x00\x00\xff\xff\xce\xd4\x03\x4a\x77\x07\x00\x00")

func init() {
	if CTX.Err() != nil {
		panic(CTX.Err())
	}

	var err error

	err = FS.Mkdir(CTX, "vendor/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "vendor/github.com/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "vendor/github.com/containerum/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "vendor/github.com/containerum/utils/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "vendor/github.com/containerum/utils/httputil/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "vendor/github.com/containerum/cherry/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileVendorGithubComContainerumCherrySwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "vendor/github.com/containerum/cherry/swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileVendorGithubComContainerumUtilsHttputilSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "vendor/github.com/containerum/utils/httputil/swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
