package omni

import (
	"fmt"
	"os"
	"os/exec"
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

func RunRuby(basePath string, command []string) error {
	oldPath := os.Getenv("PATH")
	fullPath := fmt.Sprintf("%s/bin:%s", basePath, oldPath)
	os.Setenv("PATH", fullPath)

	run := exec.Command(command[0], command[1:]...)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	run.Env = os.Environ()

	return run.Run()
}
