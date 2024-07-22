// pkg/utils/errors.go
package utils

import "errors"

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrNotFound         = errors.New("not found")
)
