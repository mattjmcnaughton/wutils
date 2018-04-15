package executor

import (
	"fmt"
)

// MockShellExecutor is a mock shell executor which records commands for
// inspection later.
type MockShellExecutor struct {
	success            bool
	executedCommands   []string
	executedCommandDir string
}

// CreateSuccessMockExecutor returns a `MockShellExecutor` where all the
// commands simulate success.
func CreateSuccessMockExecutor() MockShellExecutor {
	return MockShellExecutor{
		success: true,
	}
}

// CreateFailureMockExecutor returns a `MockShellExecutor` where all the
// commands simulate failure.
func CreateFailureMockExecutor() MockShellExecutor {
	return MockShellExecutor{
		success: false,
	}
}

// Run mocks executing the given command with no input/output binding.
func (m *MockShellExecutor) Run(name string, arg ...string) error {
	return m.mockCommand(name, arg...)
}

// RunInDir mocks executing the given command in the given dir.
func (m *MockShellExecutor) RunInDir(dir string, name string, arg ...string) error {
	cmd := func() error {
		return m.Run(name, arg...)
	}

	return m.executeInDir(dir, cmd)
}

// RunWithBoundOutput mocks executing the given command with just the output bound to
// the current shell.
func (m *MockShellExecutor) RunWithBoundOutput(name string, arg ...string) error {
	return m.mockCommand(name, arg...)
}

// RunInDirWithBoundOutput mocks executing the given command in the given dir with bound output.
func (m *MockShellExecutor) RunInDirWithBoundOutput(dir string, name string, arg ...string) error {
	cmd := func() error {
		return m.RunWithBoundOutput(name, arg...)
	}

	return m.executeInDir(dir, cmd)
}

// RunWithBoundInputOutput mocks executing the given command with the input and output
// bound to the current shell.
func (m *MockShellExecutor) RunWithBoundInputOutput(name string, arg ...string) error {
	return m.mockCommand(name, arg...)
}

// RunInDirWithBoundInputOutput mocks executing the given command in the given dir with bound
// input/output.
func (m *MockShellExecutor) RunInDirWithBoundInputOutput(dir string, name string, arg ...string) error {
	cmd := func() error {
		return m.RunWithBoundInputOutput(name, arg...)
	}

	return m.executeInDir(dir, cmd)
}

// GetExecutedCommands is a public access on the `executedCommand`.
func (m *MockShellExecutor) GetExecutedCommands() []string {
	return m.executedCommands
}

// GetExecutedCommandDir returns the dir in which the command was executed.
func (m *MockShellExecutor) GetExecutedCommandDir() string {
	return m.executedCommandDir
}

func (m *MockShellExecutor) mockCommand(name string, arg ...string) error {
	executedCommand := fmt.Sprintf("%s %v", name, arg)
	m.executedCommands = append(m.executedCommands, executedCommand)

	if m.success {
		return nil
	}

	return fmt.Errorf("Mock error")
}

func (m *MockShellExecutor) executeInDir(dir string, cmd func() error) error {
	m.executedCommandDir = dir
	return cmd()
}
