package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version     bool
	file_output string
	file_input  []string
)

// ascii img is build from http://www.network-science.de/ascii/
// is there any better way to print logo? please contact us
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

// is for cmd
func dna_flag() {
	flag.StringVar(&file_output, "o", "a.nex", "output file")
	flag.BoolVar(&version, "v", false, "version")
	flag.Usage = usage
	flag.Parse()
	if version {
		fmt.Println("SeqCombGo version: SeqCombGo/0.0.1")
		os.Exit(0)
	}
	file_input = flag.Args() // []string{"foo", "bar"}
	// 这里应该加个判断，如果输入不符合，后面的正则会报错
	if len(file_input) == 0 {
		fmt.Println("please use it in right way! see -h for help")
		os.Exit(0)
	}
	print_logo()
	fmt.Println("\n[import files:]", file_input)
	fmt.Println("[export file :]", file_output)
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
