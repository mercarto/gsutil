package gsutils

import (
	"fmt"
	"net/url"
)

// ParseGSURL takes a Google Storage URL and returns the associated bucket and path
func ParseGSURL(value string) (bucket string, path string, err error) {
	// NOTE: using named return parameters makes this logic a lot simpler
	var u *url.URL
	if u, err = url.Parse(value); err != nil {
		return
	}

	if u.Scheme != "gs" {
		err = fmt.Errorf("only 'gs' scheme is supported, received: %s", u.Scheme)
		return
	}

	bucket = u.Host
	if u.Path == "" {
		path = ""
	} else {
		// strip the preceding '/' rune from the path
		path = u.Path[1:]
	}

	return
}
