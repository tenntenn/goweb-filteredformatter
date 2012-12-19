package filter

import (
	"io"
	"io/ioutil"
	"bytes"

	"code.google.com/p/goweb/goweb"
)

type FilteredDecoder struct {
	base goweb.RequestDecoder
	filters []*Filter
}

func NewDecoder(base goweb.RequestDecoder, filters... *Filter) *FilteredDecoder {
	return &FilteredDecoder{base, filters}
}

func (self *FilteredDecoder) Unmarshal(cx *goweb.Context, v interface{}) error {

	// read body
	buf, err := ioutil.ReadAll(cx.Request.Body)
	if err != nil {
		return err
	}

	for _, filter := range self.filters {

		filter.Write(buf)
		dst, err := ioutil.ReadAll(filter)
		if err != nil {
			return err
		}

		buf = dst
	}

	cx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bufa))
	return self.base.Unmarshal(cx, v)
}
