package main

import "fmt"

func methoda(name string, f func(a int, b string)) {
	fmt.Println("Enter MethodA:", name)
	f(3030, "zdd") // 给f注入参数
	fmt.Println("Exit MethodA:", name)
}

func methodb(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	d := methodb
	methoda("zddhub", d)
}
