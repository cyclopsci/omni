package omni

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"text/template"
)

func InstallPuppet(basePath string, version string) error {
	bundlePath, err := verifyRubySupport()
	if err != nil {
		return err
	}

	absPath := path.Join(basePath, "puppet", version)
	if err := os.MkdirAll(absPath, 0755); err != nil {
		return err
	}

	const gemfile = `source 'https://rubygems.org'
gem 'puppet', '~> {{.Version}}'
gem 'puppet-lint', '~> 1.1.0'
`

	data := struct {
		Version string
	}{
		version,
	}

	var b bytes.Buffer

	t := template.Must(template.New("gemfile").Parse(gemfile))
	t.Execute(&b, data)

	gemPath := path.Join(absPath, "Gemfile")
	if err := ioutil.WriteFile(gemPath, b.Bytes(), 0755); err != nil {
		return err
	}

	install := exec.Command(bundlePath, "install", "--path", ".gem", "--binstubs")
	install.Dir = absPath
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	if err := install.Run(); err != nil {
		return err
	}

	return nil
}

func InstallAnsible(basePath string, version string) error {
	return nil
}

func verifyRubySupport() (string, error) {
	return exec.LookPath("bundle")
}

func verifyPythonSupport() error {
	if _, err := exec.LookPath("virtualenv"); err != nil {
		return err
	}
	return nil
}

func EnterRuby(path string) {
	orig := os.Getenv("PATH")
	os.Setenv("OMNI_ORIG_PATH", orig)
	fmt.Printf("export BUNDLE_GEMFILE=%s/Gemfile\n", path)
	fmt.Printf("export PATH=%s/bin:$OMNI_ORIG_PATH\n", path)
}

func EnterPython(path string) {
	fmt.Println(path)
}

func Deactivate() error {
	err := ExitRuby()
	err = ExitPython()
	return err
}

func ExitRuby() error {
	if present := os.Getenv("BUNDLE_GEMFILE"); present != "" {
		return os.Unsetenv("BUNDLE_GEMFILE")
	}
	return nil
}

func ExitPython() error {
	fmt.Println("exiting virtualenv")
	return nil
}
