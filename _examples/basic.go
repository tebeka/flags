package main

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

func main() {
	flag.Var(flags.Int(&config.port, checkPort), "port", "port to listen on")
	flag.Var(flags.String(&config.name, checkName), "name", "logger name")
	flag.Var(flags.URL(config.url), "url", "url to hit")
	flag.Var(flags.Time(&config.start, time.RFC3339), "start", "start time")
	flag.Var(flags.File(&config.in, 'r'), "input", "input file")
	flag.Parse()

	fmt.Printf("port: %d\n", config.port)
	fmt.Printf("name: %q\n", config.name)
	fmt.Printf("url: %q\n", config.url.String())
	fmt.Printf("start: %s\n", config.start)
	fmt.Printf("in: %q\n", config.in.Name())
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
