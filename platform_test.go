package omni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoverPlatformVersionsBadPath(t *testing.T) {
	assert := assert.New(t)
	_, err := discoverPlatformVersions("fixtures/all/puppet/doesnotexist")
	assert.Error(err)
}

func TestDiscoverPlatformVersions(t *testing.T) {
	assert := assert.New(t)
	versions, err := discoverPlatformVersions("fixtures/all/puppet/")
	assert.NoError(err)
	assert.Len(versions, 2)
}

func TestDiscoverPlatformsBadPath(t *testing.T) {
	assert := assert.New(t)
	_, err := discoverPlatforms("fixtures/doesnotexist/")
	assert.Error(err)
}

func TestDiscoverPlatforms(t *testing.T) {
	assert := assert.New(t)
	platforms, err := discoverPlatforms("fixtures/all/")
	assert.NoError(err)
	assert.Len(platforms, 2)
}
