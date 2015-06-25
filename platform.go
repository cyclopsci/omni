package omni

import (
	"fmt"
	"os"
	"path"
)

type Version struct {
	Label     string
	Installed bool
}

type Platform struct {
	Label    string
	Versions []Version
}

func GetPlatforms(basePath string) ([]Platform, error) {
	return discoverPlatforms(basePath)
}

func discoverPlatforms(basePath string) ([]Platform, error) {
	platforms := []Platform{}

	base, err := os.Open(basePath)
	defer base.Close()
	if os.IsNotExist(err) {
		return platforms, err
	}
	dirs, err := base.Readdirnames(0)
	if err != nil {
		return platforms, err
	}

	for _, dir := range dirs {
		versions, err := discoverPlatformVersions(path.Join(basePath, dir))
		if err != nil {
			continue
		}

		platforms = append(platforms, Platform{
			Label:    dir,
			Versions: versions,
		})
	}

	return platforms, nil
}

func discoverPlatformVersions(basePath string) ([]Version, error) {
	versions := []Version{}

	base, err := os.Open(basePath)
	defer base.Close()
	if os.IsNotExist(err) {
		return versions, err
	}
	dirs, err := base.Readdirnames(0)
	if err != nil {
		return versions, err
	}

	for _, dir := range dirs {
		versions = append(versions, Version{
			Label:     dir,
			Installed: true,
		})
	}

	return versions, nil
}

func InstallPlatform(basePath string, platform string, version string) {
	var err error
	switch platform {
	case "puppet":
		err = InstallPuppet(basePath, version)
	case "ansible":
		err = InstallAnsible(basePath, version)
	}
	if err != nil {
		fmt.Println(err)
	}
}
