package flags_test

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/tebeka/flags"
)

var config struct {
	port  int
	name  string
	url   *url.URL
	start time.Time
	in    *os.File
}

func ExampleUsage() {
	fs := flag.NewFlagSet("example", flag.ContinueOnError)
	fs.Var(flags.Int(&config.port, checkPort), "port", "port to listen on")
	fs.Var(flags.String(&config.name, checkName), "name", "logger name")
	fs.Var(flags.URL(config.url), "url", "url to hit")
	fs.Var(flags.Time(&config.start, time.RFC3339), "start", "start time")
	fs.Var(flags.File(config.in, 'r'), "input", "input file")

	args := []string{
		"-port", "999",
		"-name", "lassie",
		"-url", "http://example.com",
		"-start", "2019-11-26T19:23:42Z",
		"-input", "/dev/null",
	}

	fs.Parse(args)

	fmt.Printf("port: %d\n", config.port)
	fmt.Printf("name: %q\n", config.name)
	fmt.Printf("url: %q\n", config.url.String())
	fmt.Printf("start: %s\n", config.start)
	fmt.Printf("in: %q\n", config.in.Name())

	// Output:
	// port: 999
	// name: "lassie"
	// url: "http://example.com"
	// start: 2019-11-26 19:23:42 +0000 UTC
	// in: "/dev/null"
}

func checkPort(p int) error {
	if p < 0 {
		return fmt.Errorf("negative port - %d", p)
	}
	return nil
}

func checkName(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("empty name")
	}
	return nil
}

func init() {
	// Set defaults
	config.port = 8080
	config.name = "bugs"
	config.start = time.Now()
	config.url, _ = url.Parse("http://localhost:8080")
	config.in = os.Stdin
}
