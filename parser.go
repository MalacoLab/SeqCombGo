package main

import (
	"flag"
	"fmt"
	fas_parser "gocomb/src"
	"os"
	"strings"
	"text/template"
)

type dna struct {
	name    string
	min_dna map[string]string
	count   int
}

type charset struct {
	Name string
	From int
	To   int
}

type tmpl_data struct {
	Ntax    int
	Nchar   int
	Matrix  map[string]string
	Charset []charset
}

func main() {

	// 读取命令行，这里一定要是指针
	file_export := flag.String("o", "a.nex", "files name wait to out")
	flag.Parse()
	file_names := flag.Args() // []string{"foo", "bar"}
	fmt.Println("[ export here ]", *file_export)

	// 遍历文件
	sum_nex := make([]dna, 0, 5)
	for k, v := range file_names {
		i, j := fas_parser.Fas_parser(v)
		new_nex := dna{v, i, j}
		sum_nex = append(sum_nex, new_nex)
		fmt.Println("[ working ]", k+1, v)
	}

	// 整合若干文件的统计
	sum_charset := []charset{}
	for k, v := range sum_nex {
		n := v.name
		f := 1
		if k != 0 {
			f = sum_charset[k-1].To
		}
		t := f + v.count
		new_charset := charset{n, f, t}
		sum_charset = append(sum_charset, new_charset)
	}
	fmt.Println(sum_charset)

	// dna 的整合
	seq := sum_nex[0].min_dna
	ntax := 0 // 待修补
	nchar := sum_charset[len(sum_charset)-1].To
	for k, v := range sum_nex {
		if k == 0 {
			continue
		}
		last_count := sum_nex[k-1].count
		for k1, v1 := range v.min_dna {
			if _, ok := seq[k1]; ok {
				seq[k1] = seq[k1] + v1
			} else {
				seq[k1] = seq[k1] + strings.Repeat("?", last_count)
			}
		}
	}
	last_data := tmpl_data{ntax, nchar, seq, sum_charset}
	// fmt.Println(last_data)

	// 读取模板
	nex_tmpl, err := template.New("nex").Parse(fas_parser.Nex_tmpl)
	if err != nil {
		fmt.Println("[ tmpl err ]", err)
		return
	}

	// 覆盖创建要写入的 nex 文件
	new_file, err := os.OpenFile(*file_export, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("[ create or open file error ]", err)
		return
	}
	defer new_file.Close()

	// 写入 nex 模板
	err = nex_tmpl.Execute(new_file, last_data)
	if  err!= nil {
		fmt.Println("[ err at tmpl exec ]", err)
		return
	}
}
