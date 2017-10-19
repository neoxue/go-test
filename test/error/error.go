package main

import (
	"github.com/pkg/errors"
	"fmt"
	"reflect"
)

func main () {
	_, err := test();
	fmt.Println(err.Error())
	_, err2 := testNegative()
	fmt.Println(err2.Error())


	_, err3 := testComplex()
	fmt.Println(err3)
	fmt.Println(errors.Cause(err3))
	//fmt.Printf("%+v", err3)
}

func test() (string, error)  {
	err := errors.New("test error");
	return "", err
}

func testNegative() (string, error) {
	a := NegativeSqrtError(-9)
	//fmt.Println(reflect.Type(a))
	fmt.Println(reflect.TypeOf(a))
	return "", a
}

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number: %g", float64(f))
}

type SyntaxError struct {
	msg string
	offset int64
}

func (e *SyntaxError) Error() string {
	return e.msg
}




func testComplex() (string, error) {
	//err := errors.Errorf("error: %s", "foo")
	//cause := errors.New("by cause")
	_, cause := testComplex2()
	err3 := errors.Wrap(cause, "oh noes")
	return "", err3
}

func testComplex2() (string, error) {
	cause := errors.New("by cause1")
	err3 := errors.Wrap(cause, "oh noes1")
	return "", err3
}