package omni

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func verifyRubySupport() (string, error) {
	return exec.LookPath("bundle")
}

func EnterRuby(path string) error {
	orig := os.Getenv("PATH")
	fmt.Printf("export OMNI_ORIG_PATH=%s\n", orig)
	fmt.Printf("export BUNDLE_GEMFILE=%s/Gemfile\n", path)
	fmt.Printf("export PATH=%s/bin:%s\n", path, orig)
	return nil
}

func ExitRuby() error {
	if present := os.Getenv("BUNDLE_GEMFILE"); present != "" {
		os.Unsetenv("BUNDLE_GEMFILE")
	}
	if origPath := os.Getenv("OMNI_ORIG_PATH"); origPath != "" {
		fmt.Printf("export PATH=%s\n", origPath)
	}
	fmt.Println("export BUNDLE_GEMFILE=")
	fmt.Println("export OMNI_ORIG_PATH=")

	return nil
}

func ExecRuby(task *ExecTask) (ExecResult, error) {
	result := ExecResult{}
	oldPath := os.Getenv("PATH")
	fullPath := fmt.Sprintf("%s/bin:%s", task.Dir, oldPath)
	os.Setenv("PATH", fullPath)

	run := exec.Command(task.Command, task.Args...)
	run.Stdout = &result.Log
	run.Stderr = &result.Log
	run.Env = os.Environ()

	err := run.Run()
	if msg, ok := err.(*exec.ExitError); ok {
		result.ExitCode = msg.Sys().(syscall.WaitStatus).ExitStatus()
		return result, err
	} else {
		//assuming all non-exit errors means everything is ok
		result.ExitCode = 0
		return result, nil
	}
}
