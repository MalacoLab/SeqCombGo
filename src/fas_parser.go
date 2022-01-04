package fas_parser

import(
	"io/ioutil"
	"fmt"
)

func Fas_parser(file_name string) (map[string]string, int) {
	f, err := ioutil.ReadFile("./" + file_name)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	// fmt.Println(f)
	count := 0
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
				if len(seq) < 2 && j == 0 {
					count += k - i
				}
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
	// fmt.Println(count)
	return seq, count
}
