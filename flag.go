package main

import (
	"flag"
	"fmt"
)

var (
	file_output string
	file_input  []string
)

func dna_flag() {
	flag.StringVar(&file_output, "o", "a.nex", "files name wait to out")
	flag.Parse()
	file_input = flag.Args() // []string{"foo", "bar"}
	fmt.Println("==============")
	fmt.Println("[input  file:]", file_input)
	fmt.Println("[output file:]", file_output)
}
