package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	run(in, out)
	_ = out.Flush()
}

func run(in io.Reader, out io.Writer) {
	n := 0
	_, _ = fmt.Fscan(in, &n)

	a := 0
	b := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &a)
		_, _ = fmt.Fscan(in, &b)
		_, _ = fmt.Fprintln(out, a+b)
	}
}
