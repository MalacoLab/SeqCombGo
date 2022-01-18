package main

import (
	"fmt"
	"os"
	"text/template"
)

const f string = `#NEXUS
BEGIN DATA;
	DIMENSIONS NTAX={{ .Ntax }} NCHAR={{ .Nchar }};
	FORMAT DATATYPE=DNA GAP=- MISSING=?;
MATRIX
{{- range $k, $v := .Matrix }}
'{{ $k }}' {{ $v }}
{{- end }}
;
END;
BEGIN SETS;
{{- range $_, $i := .Charset }}
	CHARSET {{ $i.Name }} = {{ $i.From }}-{{ $i.To }};
{{- end }}
END;`

func do_impl(last_data tmpl_data) {
	// 读取模板
	// read the template
	nex_tmpl, err := template.New("nex").Parse(f)
	if err != nil {
		fmt.Println("[ tmpl err ]", err)
		return
	}

	// 覆盖创建要写入的 nex 文件
	// create the output nex file
	new_file, err := os.OpenFile(file_output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("[ create or open file error ]", err)
		return
	}
	defer new_file.Close()

	// 写入 nex 模板
	// write the nex data, last data is the inputed data
	err = nex_tmpl.Execute(new_file, last_data)
	if err != nil {
		fmt.Println("[ err at tmpl exec ]", err)
		return
	}
}
