# flags - More flag Types

[![GoDoc](https://godoc.org/github.com/tebeka/flags?status.svg)](https://godoc.org/github.com/tebeka/flags)
[![Actions Status](https://github.com/tebeka/flags/workflows/Test/badge.svg)](https://github.com/tebeka/flags/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


```go
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
	in      *os.File
	name    string
	port    int
	start   time.Time
	retries int
	url     *url.URL
}

func main() {
	flag.Var(flags.File(config.in, 'r'), "input", "input file")
	flag.Var(flags.Int(&config.retries, checkRetries), "retries", "number of retries")
	flag.Var(flags.Port(&config.port), "port", "port to listen on")
	flag.Var(flags.String(&config.name, checkName), "name", "logger name")
	flag.Var(flags.Time(&config.start, time.RFC3339), "start", "start time")
	flag.Var(flags.URL(config.url), "url", "url to hit")
	flag.Parse()

	fmt.Printf("in: %q\n", config.in.Name())
	fmt.Printf("name: %q\n", config.name)
	fmt.Printf("port: %d\n", config.port)
	fmt.Printf("retries: %d\n", config.retries)
	fmt.Printf("start: %s\n", config.start)
	fmt.Printf("url: %q\n", config.url.String())

	// Output:
	// in: "/dev/null"
	// name: "lassie"
	// port: 999
	// retries: 3
	// start: 2019-11-26 19:23:42 +0000 UTC
	// url: "http://example.com"
}

func checkRetries(n int) error {
	if n < 0 || n > 10 {
		return fmt.Errorf("retries = %d not in range [0:10]", n)
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
	config.in = os.Stdin
	config.name = "bugs"
	config.port = 8080
	config.retries = 1
	config.start = time.Now()
	config.url, _ = url.Parse("http://localhost:8080")
}
```

Now try:

```
$ ./demo --help
Usage of ./demo:
sage of /tmp/go-build652444922/b001/exe/demo:
  -input value
    	input file (default /dev/stdin)
  -name value
    	logger name (default bugs)
  -port value
    	port to listen on (default 8080)
  -retries value
    	number of retries (default 1)
  -start value
    	start time (default 2019-12-03T10:35:48+02:00)
  -url value
    	url to hit (default http://localhost:8080)
```

and

```
$ ./demo \
    -input /dev/null \
    -name lassie \
    -port 999 \
    -retries 3 \
    -start 2019-11-26T19:23:42Z \
    -url http://example.com

in: "/dev/stdin"
name: "bugs"
port: 8080
retries: 1
start: 2019-12-03 10:38:06.553133356 +0200 IST m=+0.000050679
url: "http://localhost:8080"
```


## Types

*This section is still unfinished*

- File
- Float
- Int 
- Port
- String
- Time
- URL
