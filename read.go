package xio

import "io"

var _ io.ReaderFrom = (*ReaderFromFunc)(nil)

type ReaderFromFunc func(r io.Reader) (int64, error)

func (f ReaderFromFunc) ReadFrom(r io.Reader) (int64, error) {
	return f(r)
}

func AppendReadBytes(r io.Reader, b []byte) ([]byte, error) {
	if len(b) == 0 {
		b = make([]byte, 0, 512)
	}
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
	}
}

func ReadBytes(r io.Reader) ([]byte, error) {
	return AppendReadBytes(r, nil)
}

func ReadString(r io.Reader) (string, error) {
	b, err := AppendReadBytes(r, nil)
	if err != nil {
		return "", err
	}
	return BytesToStr(b), nil
}
