package omni

import (
	"bytes"
	"fmt"
)

type ExecResult struct {
	Log      bytes.Buffer
	ExitCode int
}

func (e ExecResult) String() string {
	s := fmt.Sprintf("Exit: %v\n", e.ExitCode)
	s += fmt.Sprintln("Log:")
	s += fmt.Sprint(string(e.Log.Bytes()))
	return s
}

type ExecOptions struct {
}
