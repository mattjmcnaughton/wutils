package reporter

import (
	"fmt"

	"os"
)

// Reporter interface is responsible for reporting responses in the event of an
// errors. We make it an interface so its easier to inject a mock for unit
// testing.
type Reporter interface {
	ReportIfError(error, string, ...interface{})
}

// FmtReporter is a production implementation of the Reporter interface
// which writes errors to stdout.
type FmtReporter struct {
}

// ReportIfError outputs the string if the passed error is not nil.
func (f *FmtReporter) ReportIfError(err error, format string, a ...interface{}) {
	if err != nil {
		fmt.Printf(fmt.Sprintf("%s\n", format), a...)
		fmt.Printf("Error: %v\n", err)

		// @TODO(mattjmcnaughton) Make exit code more accurate.
		os.Exit(1)
	}
}
