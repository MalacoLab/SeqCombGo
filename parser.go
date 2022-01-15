package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type dna struct {
	name  string
	count int
	min   map[string]string
}

// 读取fas文件
// read fas format files
func fas_parser(file_name string) dna {

	// 读取文件
	f, err := ioutil.ReadFile("./" + file_name)
	if err != nil {
		fmt.Println(err)
		return dna{"", 0, nil}
	}

	count := 0
	//sequence lines amount
	i := 0  // acgt行计数
	//samples amount
	j := -1 // 标题行计数
	seq := make(map[string]string)
	indid := ""

	for k, v := range f {
		switch v {
		case '>':
			j = k + 1
		case '\n':
			if j != -1 {
				indid = string(f[j:k])
				i = k + 1
				j = -1
				continue
			}
			seq[indid] = seq[indid] + strings.ToLower(string(f[i:k]))
			i = k + 1
		}
	}
	count = len(seq[indid])
	// for k1, v1 := range seq {
	// 	fmt.Println(k1)
	// 	fmt.Println(v1)
	// }
	// fmt.Println(count)
	return dna{file_name, count, seq}
}
