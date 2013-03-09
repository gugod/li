package main

import (
	"flag"
	"os"
	"io"
	"bufio"
	"fmt"
	"time"
)

func main() {
	var lps = flag.Int("lps", 1, "line-per-second")

	flag.Parse()

	line_reader :=  bufio.NewReader(os.Stdin)

	for {
		line, error := line_reader.ReadString('\n')
		if ( error != nil ) {
			if ( error != io.EOF ) {
				fmt.Println(error)
			}
			break
		} else {
			time.Sleep( time.Second / time.Duration(*lps) )
			fmt.Print(line)
		}
	}
}
