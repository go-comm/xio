package xio

import (
	"io"
)

func NewDumpReader(r io.Reader) DumpReader {
	return &dumpReader{r: r}
}

type DumpReader interface {
	io.Reader
	Bytes() []byte
}

type dumpReader struct {
	r io.Reader
	b []byte
}

func (r *dumpReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.b = append(r.b, p[:n]...)
	return
}

func (r *dumpReader) Bytes() []byte {
	return r.b
}

func NewDumpWriter(w io.Writer) DumpWriter {
	return &dumpWriter{w: w}
}

type DumpWriter interface {
	io.Writer
	Bytes() []byte
}

type dumpWriter struct {
	w io.Writer
	b []byte
}

func (w *dumpWriter) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	w.b = append(w.b, p[:n]...)
	return
}

func (w *dumpWriter) Bytes() []byte {
	return w.b
}

func NewTotalReader(r io.Reader) TotalReader {
	return &totalReader{r: r}
}

type TotalReader interface {
	io.Reader
	Total() int
}

type totalReader struct {
	r io.Reader
	t int
}

func (r *totalReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.t += n
	return
}

func (r *totalReader) Total() int {
	return r.t
}

func NewTotalWriter(w io.Writer) TotalWriter {
	return &totalWriter{w: w}
}

type TotalWriter interface {
	io.Writer
	Total() int
}

type totalWriter struct {
	w io.Writer
	t int
}

func (w *totalWriter) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	w.t += n
	return
}

func (w *totalWriter) Total() int {
	return w.t
}
