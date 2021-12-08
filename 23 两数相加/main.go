package main

import (
	"fmt"
	"io"
)

func main() {
	str := "NESW"
	N := 0
	for {
		_, err1 := fmt.Scan(&N)
		if err1 == io.EOF {
			break
		}
		s := ""
		_, err2 := fmt.Scan(&s)
		if err2 == io.EOF {
			break
		}
		index := 0
		for i := 0; i < N; i++ {
			if s[i] == 'L' {
				index -= 1
			}else {
				index += 1
			}
		}
		for index < 0{
			index += 4
		}
		fmt.Println(string(str[index]))
	}
}
