//go:build go1.16

package xio

import (
	"io"
	"os"
)

var Discard = io.Discard

var NopCloser = io.NopCloser

var ReadAll = ReadBytes

var ReadFile = os.ReadFile

var WriteFile = os.WriteFile

var ReadDir = os.ReadDir

var CreateTemp = os.CreateTemp

var TempFile = os.CreateTemp

var MkdirTemp = os.MkdirTemp

var TempDir = os.MkdirTemp
