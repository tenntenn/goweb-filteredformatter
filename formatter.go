package filter

import (
	"io/ioutil"

	"code.google.com/p/goweb/goweb"
)

type FilteredFormatter struct {
	base goweb.Formatter
	filters []*Filter
}

func NewFormatter(base goweb.Formatter, filters... *Filter) *FilteredFormatter {
	return &FilteredFormatter{base, filters}
}

func (self *FilteredFormatter) Format(cx *goweb.Context, input interface{}) ([]uint8, error) {

	buf, err := self.base.Format(cx, input)
	if err != nil {
		return nil, err
	}

	for _, filter := range self.filters {
		filter.Write(buf)
		dst, err := ioutil.ReadAll(filter)
		if err != nil {
			return nil, err
		}

		buf = dst
	}

	return buf, nil
}

func (self *FilteredFormatter) Match(cx *goweb.Context) bool {
	return self.base.Match(cx)
}
