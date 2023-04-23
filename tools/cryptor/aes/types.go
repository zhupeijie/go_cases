package aes

import "errors"

var (
	ErrKeyLengthSixteen = errors.New("a sixteen or twenty-four or thirty-two length secret key is required")
)
