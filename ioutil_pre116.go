//go:build !go1.16

package xio

import (
	"io/ioutil"
)

var Discard = ioutil.Discard

var NopCloser = ioutil.NopCloser

var ReadAll = ReadBytes

var ReadFile = ioutil.ReadFile

var WriteFile = ioutil.WriteFile

var ReadDir = ioutil.ReadDir

var CreateTemp = ioutil.TempFile

var TempFile = ioutil.TempFile

var MkdirTemp = ioutil.TempDir

var TempDir = ioutil.TempDir
