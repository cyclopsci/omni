package omni

import (
	"errors"
)

var (
	ErrInvalidPlatform = errors.New("Invalid Platform")
	ErrInvalidVersion  = errors.New("Invalid Version")
	ErrInvalidFormat   = errors.New("Invalid output format.  Valid values: text, json")
)
