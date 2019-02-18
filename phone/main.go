package main

import "bytes"

/*
pkg: github.com/moficodes/gophercises/phone
BenchmarkNormalize-8   	 1000000	      1209 ns/op	    1024 B/op	      16 allocs/op
*/
func normalize(phone string) string {
	var buf bytes.Buffer
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

/*
pkg: github.com/moficodes/gophercises/phone
BenchmarkNormalize-8   	   30000	     41238 ns/op	  305712 B/op	     198 allocs/op
*/

// func normalize(phone string) string {
// 	re := regexp.MustCompile("\\D")
// 	return re.ReplaceAllString(phone, "")
// }

func main() {

}
