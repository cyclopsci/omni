package main

import (
	"errors"
)

var (
	ErrMissingRequiredArgs = errors.New("Missing required arguments")
	ErrInvalidPlatform     = errors.New("Invalid Platform")
	ErrInvalidVersion      = errors.New("Invalid Version")
)
