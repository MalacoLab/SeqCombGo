package main

import "strings"

func fas_mix(sum_nex []dna, sum_charset []charset) (map[string][]string, int, int) {
	// dna 的整合
	ntax := 0
	nchar := sum_charset[len(sum_charset)-1].To
	sum_dna := make(map[string][]string)
	for _, v := range sum_nex {
		for k1 := range v.min {
			_, has := sum_dna[k1]
			if !has {
				sum_dna[k1] = make([]string, len(sum_charset))
				ntax++
			}
		}
	}
	for k, v := range sum_nex {
		for _, v1 := range v.min {
			for k2 := range sum_dna {
				if _, ok := v.min[k2]; ok {
					sum_dna[k2][k] = v1
				} else {
					sum_dna[k2][k] = strings.Repeat("?", v.count)
					// 之前就没写错吗
				}
			}
		}
	}
	// fmt.Println(sum_dna)
	return sum_dna, ntax, nchar
}
