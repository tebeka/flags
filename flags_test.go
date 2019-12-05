package flags

import (
	"bytes"
	"flag"
	"fmt"
	"net/url"
	"testing"
	"time"
)

func newFlagSet() *flag.FlagSet {
	var buf bytes.Buffer
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(&buf)
	return fs
}

func TestInt(t *testing.T) {
	var i int
	fs := newFlagSet()
	isPositive := func(i int) error {
		if i < 0 {
			return fmt.Errorf("%d < 0", i)
		}
		return nil
	}
	fs.Var(Int(&i, isPositive), "port", "port to listen on")
	err := fs.Parse([]string{"-port", "8080"})
	if err != nil {
		t.Fatal(err)
	}
	if i != 8080 {
		t.Fatal(i)
	}

	err = fs.Parse([]string{"-port", "-8"})
	if err == nil {
		t.Fatal("no error on negative number")
	}
}

func TestFloat(t *testing.T) {
	var f float64
	fs := newFlagSet()
	isPositive := func(f float64) error {
		if f < 0 {
			return fmt.Errorf("%f < 0", f)
		}
		return nil
	}
	fs.Var(Float(&f, isPositive), "ratio", "compression ratio")
	err := fs.Parse([]string{"-ratio", "3.2"})
	if err != nil {
		t.Fatal(err)
	}
	if f != 3.2 {
		t.Fatal(f)
	}

	err = fs.Parse([]string{"-ratio", "-8.3"})
	if err == nil {
		t.Fatal("no error on negative number")
	}
}

func TestURL(t *testing.T) {
	var u url.URL
	fs := newFlagSet()
	fs.Var(URL(&u), "url", "server url")
	url := "http://example.com"
	err := fs.Parse([]string{"-url", url})
	if err != nil {
		t.Fatal(err)
	}

	if u.String() != url {
		t.Fatal(u.String())
	}
}

func TestString(t *testing.T) {
	var s string
	host := "example.com"
	fs := newFlagSet()
	check := func(s string) error {
		if len(s) == 0 {
			return fmt.Errorf("empty string")
		}
		return nil
	}
	fs.Var(String(&s, check), "host", "host name")
	err := fs.Parse([]string{"-host", host})
	if err != nil {
		t.Fatal(err)
	}

	if s != host {
		t.Fatal(s)
	}

	err = fs.Parse([]string{"-host", ""})
	if err == nil {
		t.Fatal("empty string")
	}
}

func TestTime(t *testing.T) {
	var tm time.Time
	fs := newFlagSet()
	fs.Var(Time(&tm, time.RFC3339), "start", "start time")
	ts := "2019-11-26T19:23:42Z"
	err := fs.Parse([]string{"-start", ts})
	if err != nil {
		t.Fatal(err)
	}

	if tm.Format(time.RFC3339) != ts {
		t.Fatal(tm)
	}

	err = fs.Parse([]string{"-start", "2017-11-26"})
	if err == nil {
		t.Fatal("parse bad time")
	}
}

func TestFile(t *testing.T) {
	t.Skip("TODO")
}

func TestPort(t *testing.T) {
	fs := newFlagSet()
	val := 9876
	var port int
	fs.Var(Port(&port), "port", "port to listen on")
	err := fs.Parse([]string{"-port", fmt.Sprintf("%d", val)})
	if err != nil {
		t.Fatal(err)
	}
	if val != port {
		t.Fatalf("%d != %d", val, port)
	}
}
