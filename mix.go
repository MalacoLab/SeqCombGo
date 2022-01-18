package main

import "strings"

// dna 的整合
// comb of dna
func fas_mix(sum_nex []dna, sum_charset []charset) (map[string][]string, int, int) {
	
	ntax := 0
	nchar := sum_charset[len(sum_charset)-1].To

	sum_dna := make(map[string][]string)

	for _, v := range sum_nex {  // v is sequence of a sample in a file
		for k1 := range v.min {  // v.min is sum of intaxname and indsq
			_, has := sum_dna[k1] // k1 indataxname
			if !has { // has is true, not have is false
				sum_dna[k1] = make([]string, len(sum_charset))
				ntax++
			}
		}
	}
	for k, v := range sum_nex {
		for _, v1 := range v.min {
			for k2 := range sum_dna {
				if _, ok := v.min[k2]; ok {
					sum_dna[k2][k] = v1 //include is fine
				} else {
					sum_dna[k2][k] = strings.Repeat("?", v.count) //not include, repeat string ?
				}
			}
		}
	}
	// fmt.Println(sum_dna)
	return sum_dna, ntax, nchar
}
