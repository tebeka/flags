package flags

import (
	"net/url"
	"strconv"
)

// ValidatedInt is an integer flag with validation function
type ValidatedInt struct {
	check func(int) error
	ptr   *int
}

func (v *ValidatedInt) String() string {
	if v.ptr == nil {
		return ""
	}
	return strconv.FormatInt(int64(*v.ptr), 10)
}

func (v *ValidatedInt) Set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if err := v.check(i); err != nil {
		return err
	}

	*v.ptr = i
	return nil
}

// NewValidatedInt return a new ValidatedInt
func NewValidatedInt(ptr *int, check func(int) error) *ValidatedInt {
	return &ValidatedInt{ptr: ptr, check: check}
}

// URL is a URL flag
// based on https://golang.org/pkg/flag/#Value
type URL struct {
	url *url.URL
}

func (u *URL) String() string {
	if u.url == nil {
		return ""
	}
	return u.url.String()
}

func (u *URL) Set(s string) error {
	url, err := url.Parse(s)
	if err != nil {
		return err
	}
	*u.url = *url
	return nil
}
