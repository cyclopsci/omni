package omni

import (
	"fmt"
	"os/exec"
)

func verifyPythonSupport() error {
	if _, err := exec.LookPath("virtualenv"); err != nil {
		return err
	}
	return nil
}

func EnterPython(path string) {
	fmt.Println(path)
}

func ExitPython() error {
	fmt.Println("exiting virtualenv")
	return nil
}
