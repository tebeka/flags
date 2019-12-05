package flags

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"
)

// IntFlag is an integer flag with validation function
type IntFlag struct {
	check func(int) error
	ptr   *int
}

// Int return a new IntFlag
func Int(ptr *int, check func(int) error) *IntFlag {
	return &IntFlag{ptr: ptr, check: check}
}

func (v *IntFlag) String() string {
	if v.ptr == nil {
		return ""
	}
	return strconv.FormatInt(int64(*v.ptr), 10)
}

// Set value from a string
func (v *IntFlag) Set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if v.check != nil {
		if err := v.check(i); err != nil {
			return err
		}
	}

	*v.ptr = i
	return nil
}

// FloatFlag is an integer flag with validation function
type FloatFlag struct {
	check func(float64) error
	ptr   *float64
}

// Float return a new FloatFlag
func Float(ptr *float64, check func(float64) error) *FloatFlag {
	return &FloatFlag{ptr: ptr, check: check}
}

func (v *FloatFlag) String() string {
	if v.ptr == nil {
		return ""
	}
	return strconv.FormatFloat(float64(*v.ptr), 'f', -1, 64)
}

// Set value from a string
func (v *FloatFlag) Set(s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	if v.check != nil {
		if err := v.check(f); err != nil {
			return err
		}
	}

	*v.ptr = f
	return nil
}

// StringFlag is an string flag with validation function
type StringFlag struct {
	check func(string) error
	ptr   *string
}

// String return a new StringFlag
func String(ptr *string, check func(string) error) *StringFlag {
	return &StringFlag{ptr: ptr, check: check}
}

func (v *StringFlag) String() string {
	if v.ptr == nil {
		return ""
	}
	return *v.ptr
}

// Set value from a string
func (v *StringFlag) Set(s string) error {
	if v.check != nil {
		if err := v.check(s); err != nil {
			return err
		}
	}

	*v.ptr = s
	return nil
}

// URLFlag is a URL flag
// based on https://golang.org/pkg/flag/#Value
type URLFlag struct {
	url *url.URL
}

// URL create a new URLFlag
func URL(url *url.URL) *URLFlag {
	return &URLFlag{url}
}

func (u *URLFlag) String() string {
	if u.url == nil {
		return ""
	}
	return u.url.String()
}

// Set value from a string
func (u *URLFlag) Set(s string) error {
	url, err := url.Parse(s)
	if err != nil {
		return err
	}
	*u.url = *url
	return nil
}

// FileFlag object
type FileFlag struct {
	ptr  *os.File
	mode byte // rwa
}

// File returns a new *File
func File(ptr *os.File, mode byte) *FileFlag {
	return &FileFlag{ptr, mode}
}

func (f *FileFlag) String() string {
	if f.ptr == nil {
		return ""
	}
	return (*f.ptr).Name()
}

// Set value from a string
func (f *FileFlag) Set(s string) error {
	if s == "-" {
		switch f.mode {
		case 'r':
			*f.ptr = *os.Stdin
		case 'w', 'a':
			*f.ptr = *os.Stdout
		}
		return nil
	}

	var file *os.File
	var err error
	switch f.mode {
	case 'r':
		file, err = os.Open(s)
	case 'w':
		file, err = os.Create(s)
	case 'a':
		file, err = os.OpenFile(s, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		return err
	}

	*f.ptr = *file
	return nil
}

// TimeFlag is a time.Time flag with specific format
type TimeFlag struct {
	ptr    *time.Time
	layout string
}

// Time return new TimeFlag
func Time(ptr *time.Time, layout string) *TimeFlag {
	return &TimeFlag{ptr, layout}
}

func (t *TimeFlag) String() string {
	if t.ptr == nil {
		return ""
	}

	return t.ptr.Format(t.layout)
}

// Set value from a string
func (t *TimeFlag) Set(s string) error {
	tm, err := time.Parse(t.layout, s)
	if err != nil {
		return err
	}

	*t.ptr = tm
	return nil
}

func checkPort(val int) error {
	// port 0 will get random free port
	const minPort, maxPort = 0, 65535
	if val < minPort || val > maxPort {
		return fmt.Errorf("port %d out of range [%d:%d]", val, minPort, maxPort)
	}
	return nil
}

// Port return an IntFlag that validates a porth
func Port(ptr *int) *IntFlag {
	return Int(ptr, checkPort)
}
