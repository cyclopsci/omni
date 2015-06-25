package omni

import (
	"os/exec"
)

func verifyPythonSupport() error {
	if _, err := exec.LookPath("virtualenv"); err != nil {
		return err
	}
	return nil
}

func EnterPython(path string) error {
	return nil
}

func ExitPython() error {
	return nil
}
