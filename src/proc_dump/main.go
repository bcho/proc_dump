package main

import (
	"flag"
	"os"
	"time"
)

const (
	ProcPath             = "/proc"
	QueryIntervalSeconds = 5
)

// Commandline options
var (
	queryIntervalSeconds int

	enableDumpCommandLine bool
)

func main() {
	flag.Parse()

	var dumpers []dumperFunc
	if enableDumpCommandLine {
		dumpers = append(dumpers, dumpCommandLine)
	}

	// TODO support missing proc.
	procs, err := parseProcs(flag.Args())
	if err != nil {
		for _, dumper := range dumpers {
			dumper(nil, err)
		}
		os.Exit(1)
	}

	var b []byte
	for range time.Tick(time.Duration(queryIntervalSeconds) * time.Second) {
		b = nil
		if err = procs.Refresh(); err == nil {
			b, err = procs.Dump()
		}

		for _, dumper := range dumpers {
			dumper(b, err)
		}
	}
}

func init() {
	flag.IntVar(
		&queryIntervalSeconds,
		"query-interval",
		QueryIntervalSeconds,
		"query interval (in seconds)",
	)

	flag.BoolVar(
		&enableDumpCommandLine,
		"enable-dump-command-line",
		false,
		"dump to command line?",
	)
}
