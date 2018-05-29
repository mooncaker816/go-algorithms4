package base

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadInts 从文件中读取所有整数，一行一个整数
func ReadInts(path string) []int {
	var ret []int
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("can not open test input %s : %v\n", path, err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		l, err := readline(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("can not read file: %v\n", err)
		}
		num, err := strconv.Atoi(strings.TrimSpace(l))
		if err != nil {
			log.Fatalf("%s is not an interger.\n", l)
		}
		ret = append(ret, num)
	}
	return ret
}

func readline(r *bufio.Reader) (s string, err error) {
	var line []byte
	var prefix bool
	for {
		line, prefix, err = r.ReadLine()
		s = s + string(line)
		if !prefix {
			break
		}
	}
	return s, err
}
