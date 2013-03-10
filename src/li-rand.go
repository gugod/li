package main

import (
	"os"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed( int64( os.Getuid() + os.Getpid() + os.Getppid() ) )
	for {
		number := rand.Int()
		fmt.Println(number)
	}
}
