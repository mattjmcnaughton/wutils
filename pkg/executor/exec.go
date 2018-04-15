package executor

import (
	"fmt"
	"os"
	"os/exec"
)

// Executor provides an interface for running shell commands. We do this instead
// of directly using `os/exec` so it is easier to unit test.
type Executor interface {
	Run(string, ...string) error
	RunWithBoundOutput(string, ...string) error
	RunWithBoundInputOutput(string, ...string) error
	RunInDir(string, string, ...string) error
	RunInDirWithBoundOutput(string, string, ...string) error
	RunInDirWithBoundInputOutput(string, string, ...string) error
}

// ShellExecutor is a "production" executor which executes commands in the shell.
type ShellExecutor struct{}

// Run executes the given command with no input/output binding.
func (s *ShellExecutor) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	return cmd.Run()
}

// RunInDir executes the given command with the given dir as the working directory.
func (s *ShellExecutor) RunInDir(dir string, name string, arg ...string) error {
	cmd := func() error {
		return s.Run(name, arg...)
	}

	return executeInDir(dir, cmd)
}

// RunWithBoundOutput executes the given command with just the output bound to
// the current shell.
func (s *ShellExecutor) RunWithBoundOutput(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// RunInDirWithBoundOutput executes, in the given directory,
// the given command with just the output bound to the current shell.
func (s *ShellExecutor) RunInDirWithBoundOutput(dir string, name string, arg ...string) error {
	cmd := func() error {
		return s.RunWithBoundOutput(name, arg...)
	}

	return executeInDir(dir, cmd)
}

// RunWithBoundInputOutput executes the given command with the input and output
// bound to the current shell.
func (s *ShellExecutor) RunWithBoundInputOutput(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// RunInDirWithBoundInputOutput executes, in the given directory,
// the given command with the input and output bound to the current shell.
func (s *ShellExecutor) RunInDirWithBoundInputOutput(dir string, name string, arg ...string) error {
	cmd := func() error {
		return s.RunWithBoundInputOutput(name, arg...)
	}

	return executeInDir(dir, cmd)
}

func executeInDir(dir string, cmd func() error) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Unable to get current directory: %v", err)
	}

	err = os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("Unable to chdir to %s: %v", dir, err)
	}

	defer os.Chdir(currentDir)

	return cmd()
}
