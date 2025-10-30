package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case ('A' <= b && b <= 'Z'):
		return (b - 'A' + 13) % 26 + 'A'
	case ('a' <= b && b <= 'z'):
		return (b - 'a' + 13) % 26 + 'a'
	default:
		return b
	}
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := range b {
		b[i] = rot13(b[i])
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}