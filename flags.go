package flags

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"
)

type flagVal[T any] struct {
	ptr   *T
	conv  func(string) (T, error)
	check func(T) error
	str   func() string
}

// Flag is a generic flag
func Flag[T any](ptr *T, conv func(string) (T, error), check func(T) error) flag.Value {
	f := flagVal[T]{
		ptr:   ptr,
		conv:  conv,
		check: check,
	}
	return &f
}

func (f *flagVal[T]) Set(s string) error {
	if f.ptr == nil {
		return fmt.Errorf("empty pointer")
	}

	val, err := f.conv(s)
	if err != nil {
		return err
	}

	if f.check != nil {
		if err := f.check(val); err != nil {
			return err
		}
	}

	*f.ptr = val
	return nil
}

func (f *flagVal[T]) String() string {
	if f.ptr == nil {
		return ""
	}

	if f.str != nil {
		return f.str()
	}

	return fmt.Sprintf("%v", *f.ptr)
}

// Int is an integer flag
func Int(ptr *int, check func(int) error) flag.Value {
	v := flagVal[int]{
		ptr:   ptr,
		conv:  strconv.Atoi,
		check: check,
	}
	return &v
}

// Float is a float64 flag
func Float(ptr *float64, check func(float64) error) flag.Value {
	v := flagVal[float64]{
		ptr:   ptr,
		conv:  func(s string) (float64, error) { return strconv.ParseFloat(s, 64) },
		check: check,
	}
	return &v
}

// String is a string flag
func String(ptr *string, check func(string) error) flag.Value {
	v := flagVal[string]{
		ptr:   ptr,
		conv:  func(s string) (string, error) { return s, nil },
		check: check,
	}
	return &v
}

// URL is a URL flag
func URL(ptr *url.URL) flag.Value {
	v := flagVal[url.URL]{
		ptr: ptr,
		conv: func(s string) (url.URL, error) {
			u, err := url.Parse(s)
			if err != nil {
				return url.URL{}, err
			}
			return *u, err
		},
	}
	return &v
}

// File is a file flag
func File(ptr **os.File, mode byte) flag.Value {
	v := flagVal[*os.File]{
		ptr:  ptr,
		conv: func(s string) (*os.File, error) { return openFile(s, mode) },
		str:  func() string { return (*ptr).Name() },
	}
	return &v
}

func openFile(s string, mode byte) (*os.File, error) {
	switch mode {
	case 'r', 'w', 'a':
		// OK
	default:
		return nil, fmt.Errorf("unknown mode: %c", mode)
	}

	if s == "-" {
		switch mode {
		case 'r':
			return os.Stdin, nil
		case 'w', 'a':
			return os.Stdout, nil
		}
	}

	switch mode {
	case 'r':
		return os.Open(s) // #nosec
	case 'w':
		return os.Create(s) // #nosec
	case 'a':
		return os.OpenFile(s, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600) // #nosec
	}

	panic("should not get here")
}

// Time is a time flag
func Time(ptr *time.Time, layout string) flag.Value {
	v := flagVal[time.Time]{
		ptr:  ptr,
		conv: func(s string) (time.Time, error) { return time.Parse(layout, s) },
		str:  func() string { return ptr.Format(layout) },
	}
	return &v
}

// Port is a port flag
func Port(ptr *int) flag.Value {
	check := func(i int) error {
		// port 0 will get random free port
		const minPort, maxPort = 0, 65535
		if i < minPort || i > maxPort {
			return fmt.Errorf("port %d out of range [%d:%d]", i, minPort, maxPort)
		}
		return nil
	}
	return Int(ptr, check)
}
