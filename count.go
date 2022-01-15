package main

import (
	"fmt"
	"regexp"
)

type charset struct {
	Name string
	From int
	To   int
}

// 遍历文件得到基本数据
func fas_sum() []dna {
	sum := []dna{}
	for i, f := range file_input {
		sum = append(sum, fas_parser(f))
		fmt.Println("[ working A ]", i+1, f)
	}
	return sum
}

// 整合若干文件的统计
func fas_count(sum_nex []dna) []charset {
	fas_charset := []charset{}
	for k, v := range sum_nex {
		n := fas_name(v.name)
		f := 1
		if k != 0 {
			f = fas_charset[k-1].To + 1
		}
		t := f + v.count - 1
		fmt.Println("[ working B ]", n, f, t)
		new_charset := charset{n, f, t}
		fas_charset = append(fas_charset, new_charset)
	}
	fmt.Println(fas_charset)
	return fas_charset
}

func fas_name(old_name string) string {
	//needed to import string
	compileRegex := regexp.MustCompile(`(w+.`)
	matchArr := compileRegex.FindStringSubmatch(old_name)
	//needed to use the string get from the old string
	new_name := matchArr[len(matchArr)-1]
	return new_name
}
