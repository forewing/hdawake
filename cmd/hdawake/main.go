package main

import (
	"flag"
	"time"

	"github.com/forewing/hdawake"
)

const (
	defaultPeriod = time.Minute
)

var (
	name   = flag.String("path", "./.awake", "path to target")
	period = flag.Duration("time", defaultPeriod, "period of wake up action")
)

func init() {
	flag.Parse()
}

func main() {
	hdawake.Keep(*name, *period)
}
