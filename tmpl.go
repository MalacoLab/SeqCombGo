package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func do_impl(last_data tmpl_data) {

	f, err := ioutil.ReadFile("nex.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取模板
	nex_tmpl, err := template.New("nex").Parse(string(f))
	if err != nil {
		fmt.Println("[ tmpl err ]", err)
		return
	}

	// 覆盖创建要写入的 nex 文件
	new_file, err := os.OpenFile(file_output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("[ create or open file error ]", err)
		return
	}
	defer new_file.Close()

	// 写入 nex 模板
	err = nex_tmpl.Execute(new_file, last_data)
	if err != nil {
		fmt.Println("[ err at tmpl exec ]", err)
		return
	}
}
