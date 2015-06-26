package omni

import (
	"bytes"
	"fmt"
	"path"
)

type ExecResult struct {
	Log      bytes.Buffer
	ExitCode int
}

func (e ExecResult) String() string {
	s := fmt.Sprintln(e.ExitCode)
	s += fmt.Sprint(string(e.Log.Bytes()))
	return s
}

type ExecFormat string

const (
	FormatText ExecFormat = "text"
	FormatJSON ExecFormat = "json"
)

type ExecOptions struct {
	BasePath string
	Platform string
	Version  string
	Command  []string
	Output   string
	Format   ExecFormat
}

type ExecTask struct {
	Dir     string
	Command string
	Args    []string
}

func Exec(basePath string, platform string, version string, command []string, opts ExecOptions) error {
	task := ExecTask{
		Dir:     path.Join(basePath, platform, version),
		Command: command[0],
	}
	if len(command) > 1 {
		task.Args = command[1:]
	}

	switch platform {
	case "puppet":
		result, err := ExecRuby(&task)
		fmt.Print(result)
		return err
	default:
		return ErrInvalidPlatform
	}
}

func ExecMultiple(basePath string, platform string, versions []string, command []string, opts ExecOptions) error {
	for _, version := range versions {
		if err := Exec(basePath, platform, version, command, opts); err != nil {
			return err
		}
	}
	return nil
}
