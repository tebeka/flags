package flags

import (
	"flag"
	"net/url"
	"os"
	"time"
)

var (
	Add = NewFlags(flag.CommandLine)
)

type Flags struct {
	fs *flag.FlagSet
}

func NewFlags(fs *flag.FlagSet) *Flags {
	return &Flags{fs}
}

func (f *Flags) File(ptr *os.File, mode byte, name, usage string) {
	f.fs.Var(File(ptr, mode), name, usage)
}

func (f *Flags) Float(ptr *float64, check func(float64) error, name, usage string) {
	f.fs.Var(Float(ptr, check), name, usage)
}

func (f *Flags) Int(ptr *int, check func(int) error, name, usage string) {
	f.fs.Var(Int(ptr, check), name, usage)
}

func (f *Flags) Port(ptr *int, name, usage string) {
	f.fs.Var(Port(ptr), name, usage)
}

func (f *Flags) String(ptr *string, check func(string) error, name, usage string) {
	f.fs.Var(String(ptr, check), name, usage)
}

func (f *Flags) Time(ptr *time.Time, layout string, name, usage string) {
	f.fs.Var(Time(ptr, layout), name, usage)
}

func (f *Flags) URL(url *url.URL, name, usage string) {
	f.fs.Var(URL(url), name, usage)
}
