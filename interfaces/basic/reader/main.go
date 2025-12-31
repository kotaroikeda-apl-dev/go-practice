package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	buf := make([]byte, 8)

	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("read %d bytes: %q\n", n, buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error: %v\n", err)
			break
		}
	}
}
