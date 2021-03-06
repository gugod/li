package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
)

func main() {
	sum := 0.0
	line_reader :=  bufio.NewReader(os.Stdin)

	for {
		line, error := line_reader.ReadString('\n')
		if ( error != nil ) {
			if ( error != io.EOF ) {
				fmt.Println(error)
			}
			break
		} else {
			var n = 0.0
			fmt.Sscanf(line, "%f", &n)
			sum = sum + n
		}
	}
	fmt.Printf("%f\n", sum)	
}
