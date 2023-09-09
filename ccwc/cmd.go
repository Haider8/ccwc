package ccwc

import (
	"flag"
	"fmt"
	"strings"
)

var count_bytes = flag.Bool("c", false, "print the byte counts")

func execute() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), " "))
}
