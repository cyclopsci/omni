package omni

import (
	"encoding/json"
	"fmt"
	"path"
)

type ExecResult struct {
	Platform string
	Version  string
	Log      string
	ExitCode int
}

func (e ExecResult) String() string {
	return fmt.Sprint(e.Log)
}

type ExecOptions struct {
	Output string
	Format string
}

type ExecTask struct {
	Platform string
	Version  string
	Dir      string
	Command  string
	Args     []string
}

func Exec(basePath string, platform string, version string, command []string, opts ExecOptions) error {
	result, err := doExec(basePath, platform, version, command, opts)
	if err != nil {
		return err
	}
	return writeOutput([]ExecResult{result}, opts)
}

func ExecMultiple(basePath string, platform string, versions []string, command []string, opts ExecOptions) error {
	results := []ExecResult{}
	for _, version := range versions {
		result, err := doExec(basePath, platform, version, command, opts)
		if err != nil {
			return err
		}
		results = append(results, result)
	}
	return writeOutput(results, opts)
}

func doExec(basePath string, platform string, version string, command []string, opts ExecOptions) (ExecResult, error) {
	task := ExecTask{
		Platform: platform,
		Version:  version,
		Dir:      path.Join(basePath, platform, version),
		Command:  command[0],
	}
	if len(command) > 1 {
		task.Args = command[1:]
	}

	switch platform {
	case "puppet":
		return ExecRuby(&task), nil
	default:
		return ExecResult{}, ErrInvalidPlatform
	}
}

func writeOutput(results []ExecResult, opts ExecOptions) error {
	switch opts.Format {
	case "json":
		out, err := json.Marshal(results)
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	case "text":
		for _, result := range results {
			fmt.Print(result)
		}
		return nil
	default:
		return ErrInvalidFormat
	}
	return nil
}
