//go:build !windows
// +build !windows

package testscript

import "os"

func envvarname(k string) string {
	return k
}

func stopProcess(proc *os.Process) error {
	return proc.Signal(os.Interrupt)
}
