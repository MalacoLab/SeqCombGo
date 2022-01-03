package main

import (
	"flag"
	"fmt"
	"gocomb/src"
	"io/ioutil"
	"os"
	"text/template"
)

type Person struct {
	Name string
	DNA  string
}

func main() {

	file_flag := flag.String("o", "a.nex", "files name wait to out")
	// 这里一定要是指针样子
	flag.Parse()
	file_names := flag.Args() // []string{"foo", "bar"}
	// file_out := ""
	fmt.Println("输出在这里", *file_flag)

	nex_tmpl, err := template.New("nex").Parse(nex_tmpl.Nex_tmpl)
	if err != nil {
		panic("tmpl err")
	}

	new_file, err := os.OpenFile("a.nex", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
        fmt.Println("open file error :", err)
        return
    }
	defer new_file.Close()


	for _, v := range file_names {
		new_nex := read(v)
		err := nex_tmpl.Execute(new_file, new_nex)
		if err != nil {
			fmt.Println("err at tmpl exec", err)
		}

	}
}

func read(file_name string) map[string]string {
	f, err := ioutil.ReadFile("./" + file_name)
	if err != nil {
		fmt.Println(err)
		return nil
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
	// for k1, v1 := range seq {
	// 	fmt.Println(k1)
	// 	fmt.Println(v1)
	// }
	return seq
}
