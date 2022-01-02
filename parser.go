package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	file_flag := flag.String("o", "a.ast", "files name wait to out")
	// 这里一定要是指针样子
	flag.Parse()
	file_names := flag.Args() // []string{"foo", "bar"}
	// file_out := ""
	fmt.Println("输出在这里", *file_flag)
	for _, v := range file_names {
		// fmt.Println(k, v)
		read(v)
	}
}

func read(file_name string) {
	f, err := ioutil.ReadFile("./" + file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(f)

	i := 0 // DNA行计数
	j := 0 // 非序列行计数
	seq := make(map[string]string)
	section := ""
	// fmt.Println('a', 'c', 'g', 't', '-', '\n', '\r')
	for k, v := range f {
		switch v {
		case 'a', 'c', 'g', 't', '-':
			if j != 0 {
				continue
			}
			if i == 0 {
				i = k
			}
		case '\n':
			if i != 0 {
				seq[section] = seq[section] + string(f[i:k])
				i = 0
				continue
			}
			section = string(f[j:k])
			j = 0
		default:
			if j == 0 {
				j = k + 1
			}
		}
	}
	for k1, v1 := range seq {
		fmt.Println(k1)
		fmt.Println(v1)
	}
}
