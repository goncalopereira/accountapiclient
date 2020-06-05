package test

import "fmt"

var (
	ErrBrokenHTTPClient = fmt.Errorf("boom")
	ErrBrokenConfig     = fmt.Errorf("broken config")
)
