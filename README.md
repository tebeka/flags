# flags - More flag Types


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
```

Now try:

```
$ go run basic.go --help
Usage of /tmp/go-build457730440/b001/exe/basic:
  -input value
    	input file (default /dev/stdin)
  -name value
    	logger name (default bugs)
  -port value
    	port to listen on (default 8080)
  -start value
    	start time (default 2019-11-26T20:26:10+02:00)
  -url value
    	url to hit (default http://localhost:8080)
```

and

```
$ go run basic.go \
    -port 999 \
    -name lassie \
    -url http://example.com \
    -start 2019-11-26T19:23:42Z \
    -input /dev/null
port: 999
name: "lassie"
url: "http://example.com"
start: 2019-11-26 19:23:42 +0000 UTC
in: "/dev/null
```


## Types

*This section is still unfinished*

- File
- Int 
- String
- Time
- URL
