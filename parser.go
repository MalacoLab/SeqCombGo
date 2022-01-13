package main

import (
	"fmt"
	"io/ioutil"
)

type dna struct {
	name    string
	count   int
	min_dna map[string]string
}

// 读取fas文件
func fas_parser(file_name string) dna {

	// 读取文件
	f, err := ioutil.ReadFile("./" + file_name)
	if err != nil {
		fmt.Println(err)
		return dna{"", 0, nil}
	}

	count := 0
	i := 0  // acgt行计数
	j := -1 // 标题行计数
	seq := make(map[string]string)
	section := ""

	for k, v := range f {
		switch v {
		case '>':
			j = k
			count++
		case '\n':
			if j != -1 {
				section = string(f[j:k])
				i = k + 1
				j = -1
				continue
			}
			seq[section] = seq[section] + string(f[i:k])
			i = k + 1
		}
	}
	// for k1, v1 := range seq {
	// 	fmt.Println(k1)
	// 	fmt.Println(v1)
	// }
	// fmt.Println(count)
	return dna{file_name, count, seq}
}
