package reporter

import (
	"testing"
)

// MockReporter mocks the Reporter interface
// which writes errors to stdout.
type MockReporter struct {
	reported bool
}

// ReportIfError mocks outputting the string if the passed error is not nil.
// Really it checks if it printer to screen.
func (m *MockReporter) ReportIfError(err error, format string, a ...interface{}) {
	m.reported = false

	if err != nil {
		m.reported = true
	}
}

// AssertCalled is a helper method for checking if the helper was called.
func (m *MockReporter) AssertCalled(t *testing.T) {
	if !m.reported {
		t.Fatalf("Reporter should have been called")
	}
}

// AssertNotCalled is a helper method for checking if the helper was not called.
func (m *MockReporter) AssertNotCalled(t *testing.T) {
	if m.reported {
		t.Fatalf("Reporter should not have been called")
	}
}
