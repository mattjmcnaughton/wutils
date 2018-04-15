package reporter

// MockReporter mocks the Reporter interface
// which writes errors to stdout.
type MockReporter struct {
	reported bool
}

// ReportIfError mocks outputting the string if the passed error is not nil.
// Really it checks if it printer to screen.
func (f *MockReporter) ReportIfError(err error, format string, a ...interface{}) {
	f.reported = false

	if err != nil {
		f.reported = true
	}
}

// Reported is a public accessor on the reported field.
func (f *MockReporter) Reported() bool {
	return f.reported
}
