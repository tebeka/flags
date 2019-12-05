package flags

import (
	"flag"
	"net/url"
	"os"
	"time"
)

var (
	// Add allows a shortcut to write: flags.Add.Port(&port, "port", "listen port")
	Add = NewFlags(flag.CommandLine)
)

// Flags is a struct for adding flags
type Flags struct {
	fs *flag.FlagSet
}

// NewFlags return new Flags
func NewFlags(fs *flag.FlagSet) *Flags {
	return &Flags{fs}
}

// File adds a File flag. Mode can be 'r', 'w' or 'a'
func (f *Flags) File(ptr *os.File, mode byte, name, usage string) {
	f.fs.Var(File(ptr, mode), name, usage)
}

// Float add a Float flag
// If check is not nil, it'll be used to validate the value
func (f *Flags) Float(ptr *float64, check func(float64) error, name, usage string) {
	f.fs.Var(Float(ptr, check), name, usage)
}

// Int adds an Int flag
// If check is not nil, it'll be used to validate the value
func (f *Flags) Int(ptr *int, check func(int) error, name, usage string) {
	f.fs.Var(Int(ptr, check), name, usage)
}

// Port adds a Port flag
func (f *Flags) Port(ptr *int, name, usage string) {
	f.fs.Var(Port(ptr), name, usage)
}

// String adds a String flag
// If check is not nil, it'll be used to validate the value
func (f *Flags) String(ptr *string, check func(string) error, name, usage string) {
	f.fs.Var(String(ptr, check), name, usage)
}

// Time adds a Time flag
// layout should be one that time.Parse understand
func (f *Flags) Time(ptr *time.Time, layout string, name, usage string) {
	f.fs.Var(Time(ptr, layout), name, usage)
}

// URL adds a URL flag
func (f *Flags) URL(url *url.URL, name, usage string) {
	f.fs.Var(URL(url), name, usage)
}
