package main

type MyReader struct {

}

func main () {
	reader := MyReader{}
	reader.Validate(MyReader{})
}
