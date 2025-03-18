package main

import (
	"fmt"

	"github.com/andrewbatallones/csgo/examples"
)

func main() {
	fmt.Printf("%s\n", examples.Runlength("aaaabbccca"))
	fmt.Printf("%s\n", examples.Runlength("foobar"))
	fmt.Printf("%s\n", examples.Runlength("aaaaaa"))
	fmt.Printf("%s\n", examples.Runlength(""))

	output, _ := examples.DecodeRun("1f2o1b1a1r")
	fmt.Printf("%s\n", output)
}
