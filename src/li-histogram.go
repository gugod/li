package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"flag"
)

func main() {
	interval := flag.Int("interval", 10, "The interval of histogram")
	flag.Parse()

	histogram := make(map[int]int)

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
			histogram[ *interval * int( n / float64(*interval) ) ] += 1
		}
	}
	// XXX
	fmt.Print(histogram)
}
