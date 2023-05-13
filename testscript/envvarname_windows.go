package testscript

import (
	"os"
	"strings"
	"syscall"
)

func envvarname(k string) string {
	return strings.ToLower(k)
}

func stopProcess(proc *os.Process) error {
	d, e := syscall.LoadDLL("kernel32.dll")
	if e != nil {
		return e
	}
	p, e := d.FindProc("GenerateConsoleCtrlEvent")
	if e != nil {
		return e
	}
	r, _, e := p.Call(syscall.CTRL_BREAK_EVENT, uintptr(proc.Pid))
	if r == 0 {
		return e
	}
	return nil
}
