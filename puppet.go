package omni

import (
	"bytes"
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

func EnterPuppet(basePath string, version string) error {
	return EnterRuby(path.Join(basePath, "puppet", version))
}

func ExitPuppet() error {
	return ExitRuby()
}
