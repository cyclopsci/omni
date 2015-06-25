package omni

import (
	"fmt"
	"os"
	"os/exec"
)

func verifyRubySupport() (string, error) {
	return exec.LookPath("bundle")
}

func EnterRuby(path string) {
	orig := os.Getenv("PATH")
	os.Setenv("OMNI_ORIG_PATH", orig)
	fmt.Printf("export BUNDLE_GEMFILE=%s/Gemfile\n", path)
	fmt.Printf("export PATH=%s/bin:$OMNI_ORIG_PATH\n", path)
}

func ExitRuby() error {
	if present := os.Getenv("BUNDLE_GEMFILE"); present != "" {
		return os.Unsetenv("BUNDLE_GEMFILE")
	}
	return nil
}
