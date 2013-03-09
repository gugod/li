package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
)

func isBlankLine(line string) bool {
	if (line == "\n") {
		return true
	}
	return false
}

func main() {
	line_reader :=  bufio.NewReader(os.Stdin)
	for {
		line, error := line_reader.ReadString('\n')
		if ( error != nil ) {
			if ( error != io.EOF ) {
				fmt.Println(error)
			}
			break
		} else {
			// More strict check needed.
			if ( ! isBlankLine(line) ) {
				fmt.Printf("%s", line)
			}
		}
	}
}
