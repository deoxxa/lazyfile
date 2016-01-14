package lazyfile

import (
	"io"
	"os"
)

type LazyFile struct {
	name string
	err  error
	rd   io.ReadCloser
}

func Open(name string) *LazyFile {
	return &LazyFile{
		name: name,
	}
}

func (l *LazyFile) Read(p []byte) (int, error) {
	if l.err != nil {
		return 0, l.err
	}

	if l.rd == nil {
		rd, err := os.Open(l.name)
		if err != nil {
			l.err = err
			return 0, err
		}
		l.rd = rd
	}

	return l.rd.Read(p)
}

func (l *LazyFile) Close() error {
	if l.err != nil {
		return l.err
	}

	if l.rd == nil {
		return nil
	}

	l.err = l.rd.Close()
	l.rd = nil

	return l.err
}
