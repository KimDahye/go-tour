package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	if err != nil {
		return n, err
	}

	for i, v := range b {
		b[i] = rot13Substitute(v)
	}
	return n, err
}

func rot13Substitute(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a = 'a'
		z = 'z'
	case 'A' <= b && b <= 'Z':
		a = 'A'
		z = 'Z'
	default:
		return b
	}

	return (b-a+13)%(z-a+1) + a
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
