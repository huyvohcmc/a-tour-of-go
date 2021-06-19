package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(s []byte) (n int, err error) {
	n, err = r.r.Read(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] < 'Z' {
			s[i] = 65 + (((s[i] - 65) + 13) % 26)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			s[i] = 97 + (((s[i] - 97) + 13) % 26)
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
