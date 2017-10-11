package main

func Bar() string {
	return "bar"
}

func Foo() string {
	return "foo"
}

func Qux(v string) string {
	if v == "foo" {
		return Foo()
	}

	if v == "bar" {
		return Bar()
	}

	return "INVALID"
}

type Foo struct {
	Name    string
	Ports   []int
	Enabled bool
}

type foo struct {
	quz Qux
}

func main() {
	foo := Foo{Name: "gopher", Ports: []int{80, 443}, Enabled: true}
}
