package main

import (
	"strings"
)

type tmpl_data struct {
	Ntax    int
	Nchar   int
	Matrix  map[string]string
	Charset []charset
}

func main() {

	dna_flag()

	sum_nex := fas_sum()

	sum_charset := fas_count(sum_nex)

	sum_dna, ntax, nchar := fas_mix(sum_nex, sum_charset)

	matrix := make(map[string]string, ntax)
	for k := range sum_dna {
		matrix[k] = strings.Join(sum_dna[k], "")
	}

	// 准备发射到模板的数据
	last_data := tmpl_data{ntax, nchar, matrix, sum_charset}
	// fmt.Println(last_data)
	do_impl(last_data)
}
