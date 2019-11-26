package flags

import (
	"flag"
	"fmt"
	"testing"
)

func TestValidatedInt(t *testing.T) {
	var i int
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	isPositive := func(i int) error {
		if i < 0 {
			return fmt.Errorf("%d < 0", i)
		}
		return nil
	}
	fs.Var(NewValidatedInt(&i, isPositive), "port", "port to listen on")
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
