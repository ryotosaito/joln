package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	printHelp = flag.Bool("h", false, "Print help")
	noNewLine = flag.Bool("n", false, "No newline at the end of output if specified")
	separator = flag.String("s", "\n", "Field separator for joining lines")
	columnLen = flag.Uint("w", 0, "Wrap output every specified number of columns (0 for nowrap)")
)

func main() {
	flag.Parse()

	if *printHelp {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	var colCnt uint
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	colCnt = 0

	for {
		buf, isPrefix, err := r.ReadLine()
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		if err == io.EOF {
			if !*noNewLine {
				w.WriteByte('\n')
			}
			w.Flush()
			return
		}
		w.Write(buf)

		// write newline (-w option specified) or separator
		colCnt++
		if *columnLen != 0 && colCnt == *columnLen {
			w.WriteByte('\n')
			colCnt = 0
		} else {
			_, err = r.Peek(1)
			if !isPrefix && err != io.EOF {
				w.WriteString(*separator)
			}
		}
	}
}
