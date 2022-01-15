package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	file_output string
	file_input  []string
)

// http://www.network-science.de/ascii/
// is there any better way to print logo? plz contact us
func print_logo() {
	fmt.Print("__              ___                _       ___      \n" +
		"/ _\\ ___  __ _  / __\\___  _ __ ___ | |__   / _ \\___  \n" +
		"\\ \\ / _ \\/ _` |/ /  / _ \\| '_ ` _ \\| '_ \\ / /_\\/ _ \\ \n" +
		"_\\ \\  __/ (_| / /__| (_) | | | | | | |_) / /_\\ (_) |\n" +
		"\\__/\\___|\\__, \\____/\\___/|_| |_| |_|_.__/\\____/\\___/ \n" +
		"            |_|                                      \n")
	fmt.Println("  Sequence Combination tool written in Golang")
	fmt.Println("Version 0.0.1  Authors:An,G;Zhang,G  License:GPL-3.0")
}

// __              ___                _       ___
// / _\ ___  __ _  / __\___  _ __ ___ | |__   / _ \___
// \ \ / _ \/ _` |/ /  / _ \| '_ ` _ \| '_ \ / /_\/ _ \
// _\ \  __/ (_| / /__| (_) | | | | | | |_) / /_\\ (_) |
// \__/\___|\__, \____/\___/|_| |_| |_|_.__/\____/\___/
//             |_|

func dna_flag() {
	flag.StringVar(&file_output, "o", "a.nex", "output file")
	flag.Usage = usage
	flag.Parse()
	file_input = flag.Args() // []string{"foo", "bar"}
	// 这里应该加个判断，如果输入不符合，后面的正则会报错
	print_logo()
	fmt.Println()
	fmt.Println("[input  file:]", file_input)
	fmt.Println("[output file:]", file_output)
}

func usage() {
	print_logo()
	fmt.Fprintf(os.Stderr, `
Example:
	SeqCombGo -o export.nex import1.fas import2.fas ...
Options:
`)
	flag.PrintDefaults()
}
