package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
)

func main() {
	line_reader :=  bufio.NewReader(os.Stdin)
	line_number := 1

	for {
		line, error := line_reader.ReadString('\n')
		if ( error != nil ) {
			if ( error != io.EOF ) {
				fmt.Println(error)
			}
			break
		} else {
			fmt.Printf("%d %s", line_number, line)
		}
		line_number++
	}
}
