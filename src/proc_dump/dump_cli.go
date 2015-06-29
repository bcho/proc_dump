// Dump process stats to command line.
package main

import (
	"os"
)

func dumpCommandLine(b []byte, err error) {
	if err != nil {
		b = mustMarshalErrToJSON(err)
	}
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}
