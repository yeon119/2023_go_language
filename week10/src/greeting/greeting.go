package greeting

import "fmt"

func Hello() { //함수의 이름 첫 글자를 대문자로 해야 외부 파일에서 참조할수 있다
	fmt.Println("안녕하세요")
}

func Hi() {
	fmt.Println("안녕")
}
