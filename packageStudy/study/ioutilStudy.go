package study

import (
	"fmt"
	"io/ioutil"
	"strings"
)
func IoutilMain(){
	content, err := ioutil.ReadFile("./demo.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file content: %s,\n 类型：%T\n", content,content)

	s:=string(content)
	fmt.Printf("内容:\n%v\n 类型：%T,长度l=%d\n", s,s,len(s))
	fmt.Printf("ft:",strings.Contains(s, "Co"))
}
