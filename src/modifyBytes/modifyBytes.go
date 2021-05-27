package main

import "fmt"

func modifyBytes() {
	b := []byte("가나다")
	b[2]++ // slice라 수정 가능, string은 수정 불가능
	fmt.Println(string(b))
}

func main() {
	modifyBytes()
}
