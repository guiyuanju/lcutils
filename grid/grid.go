package grid

import (
	"fmt"
	"strconv"
	"strings"
)

func Print(g [][]int) {
	for _, r := range g {
		for _, v := range r {
			fmt.Printf("%6d", v)
		}
		fmt.Println()
	}
}

func Ints(s string) [][]int {
	var res [][]int
	s = s[1 : len(s)-1]
	if len(s) == 0 {
		return res
	}
	for p := range strings.SplitSeq(s, "],[") {
		p = strings.Trim(p, "[]")
		if len(p) == 0 {
			res = append(res, []int{})
			continue
		}
		ints := strings.Split(p, ",")
		var cur []int
		for _, i := range ints {
			n, err := strconv.ParseInt(i, 10, 0)
			if err != nil {
				panic(err)
			}
			cur = append(cur, int(n))
		}
		res = append(res, cur)
	}
	return res
}

func Strings(s string) [][]string {
	var res [][]string
	s = s[1 : len(s)-1]
	if len(s) == 0 {
		return res
	}
	for p := range strings.SplitSeq(s, "],[") {
		p = strings.Trim(p, "[]")
		if len(p) == 0 {
			res = append(res, []string{})
			continue
		}
		ints := strings.Split(p, ",")
		var cur []string
		for _, i := range ints {
			cur = append(cur, strings.Trim(i, "\""))
		}
		res = append(res, cur)
	}
	return res
}

func Bytes(s string) [][]byte {
	var res [][]byte
	s = s[1 : len(s)-1]
	if len(s) == 0 {
		return res
	}
	for p := range strings.SplitSeq(s, "],[") {
		p = strings.Trim(p, "[]")
		if len(p) == 0 {
			res = append(res, []byte{})
			continue
		}
		ints := strings.Split(p, ",")
		var cur []byte
		for _, i := range ints {
			cur = append(cur, strings.Trim(i, "\"")[0])
		}
		res = append(res, cur)
	}
	return res
}
