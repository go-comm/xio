package xio

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateFile(filename string, perm os.FileMode) (*os.File, error) {
	return OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
}

func ViewFile(filename string) (*os.File, error) {
	return OpenFile(filename, os.O_RDONLY, 0)
}

func OpenFile(filename string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(filename, flag, perm)
	if err == nil {
		return f, nil
	}
	if flag&os.O_CREATE == os.O_CREATE {
		if !os.IsExist(err) {
			if err1 := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err1 != nil {
				return nil, err1
			}
		}
		f, err = os.OpenFile(filename, flag, perm)
	}
	return f, err
}

func CloseFile(f *os.File, err error) error {
	if f == nil {
		return err
	}
	if err2 := f.Close(); err2 != nil && err == nil {
		err = err2
	}
	return err
}

func FileSize(filename string) (int64, error) {
	fi, err := os.Stat(filename)
	if err == nil {
		return fi.Size(), nil
	}
	if !os.IsExist(err) {
		return 0, err
	}
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer func() {
		f.Close()
	}()
	return io.Copy(Discard, f)
}

func WriteToFile(name string, data interface{}, perm os.FileMode) (n int64, err error) {
	var f *os.File
	f, err = OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer func() { err = CloseFile(f, err) }()
	n, err = WriteData(f, data)
	return
}

func ReadFromFile(name string, w io.Writer) (n int64, err error) {
	var f *os.File
	f, err = OpenFile(name, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}
	defer func() { err = CloseFile(f, err) }()
	return io.Copy(w, f)
}

func ReadBytesFromFile(name string) ([]byte, error) {
	return ReadFile(name)
}

func ReadStringFromFile(name string) (string, error) {
	b, err := ReadBytesFromFile(name)
	if err != nil {
		return "", err
	}
	return BytesToStr(b), nil
}

func CopyFile(dst string, src string, dperm ...os.FileMode) (int64, error) {
	fs, err := OpenFile(src, os.O_RDONLY, 0)
	if err != nil {
		return 0, fmt.Errorf("%s %w", src, err)
	}
	defer fs.Close()
	fsstat, err := fs.Stat()
	if err != nil {
		return 0, fmt.Errorf("%s %w", src, err)
	}
	dp := fsstat.Mode()
	if len(dperm) > 0 {
		dp = dperm[0]
	}
	os.Remove(dst)
	fd, err := OpenFile(dst, os.O_CREATE|os.O_RDWR, dp)
	if err != nil {
		return 0, fmt.Errorf("%s %w", dst, err)
	}
	defer fd.Close()

	return io.Copy(fd, fs)
}
