package filter

import (
	"io"
	"bytes"
)

// Filter for goweb.Formatter and goweb.Decoder
type Filter struct {
	io.Reader
	io.Writer
}
