package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()
	fmt.Println("Hello, World!")
	fmt.Printf("port: %s\n", *port)
	fmt.Println(flag.Arg(0), "nono")
	fmt.Println(flag.NArg(), "flag.NArg()")
}
