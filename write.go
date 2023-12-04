package xio

import (
	"errors"
	"io"
)

var ErrUnexpectedData = errors.New("unexpected data")

func WriteData(w io.Writer, data interface{}) (n int64, err error) {
	switch d := data.(type) {
	case []byte:
		return WriteBytes(w, d)
	case string:
		return WriteString(w, d)
	case io.WriterTo:
		n, err = d.WriteTo(w)
	case io.Reader:
		n, err = io.Copy(w, d)
	default:
		n, err = 0, ErrUnexpectedData
	}
	return
}

func WriteBytes(w io.Writer, data []byte) (n int64, err error) {
	var m int
	m, err = w.Write(data)
	return int64(m), err
}

func WriteString(w io.Writer, data string) (n int64, err error) {
	var m int
	if sw, ok := w.(io.StringWriter); ok {
		m, err = sw.WriteString(data)
	} else {
		m, err = w.Write(StrToBytes(data))
	}
	return int64(m), err
}
