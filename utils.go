package gosrm

import (
	"errors"
	"net/url"
	"strings"
)

// BaseUrl Gets Base url from options
func (o *Options) BaseUrl() (*url.URL, error) {
	u, err := url.Parse(o.Url.Host)
	if err != nil {
		return nil, errors.New("gosrm: bad server Url string")
	}

	u.Path += strings.Join([]string{
		o.Service,
		o.Version,
		o.Profile,
	}, "/")

	return u, nil
}
