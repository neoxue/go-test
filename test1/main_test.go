package main

import "testing"

func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}
}

func TestQuz(t *testing.T) {
	result := Qux("bar")
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}

	result = Qux("qux")
	if result != "INVALID" {
		t.Errorf("expecting INVALID, got %s", result)
	}
}

//func TestFoo(t *testing.T) {
//	t.Error("intentional error 1")
//}
//
////func TestBar(t *testing.T) {
////	result := Bar()
////	if result != "bar" {
////		t.Errorf("expecting bar, got %s", result)
////	}
////}
//
//func TestQuz(t *testing.T) {
//	t.Error("intentional error 2")
//}
