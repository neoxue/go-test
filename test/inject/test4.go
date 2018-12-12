package main

import (
	"fmt"
	"reflect"
)

type Injector struct {
	mappers map[reflect.Type]reflect.Value // 根据类型map实际的值
}

func (inj *Injector) SetMap(value interface{}) {
	inj.mappers[reflect.TypeOf(value)] = reflect.ValueOf(value)
}

func (inj *Injector) Get(t reflect.Type) reflect.Value {
	return inj.mappers[t]
}

func (inj *Injector) Invoke(i interface{}) []reflect.Value {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Func {
		panic("Should invoke a function!")
	}
	inValues := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		inValues[k] = inj.Get(t.In(k))
	}
	ret := reflect.ValueOf(i).Call(inValues)
	return ret
}

func New() *Injector {
	return &Injector{make(map[reflect.Type]reflect.Value)}
}

func Host(name string, i interface{}) { // 让注入的方法不受限制
	inj.Invoke(i) // 利用注入器中的环境调用f
}

func Dependency(a int, b string) {
	fmt.Println("Dependency: ", a, b)
}

type Camera struct { // 定义一个单反相机
	Name string
	L    Lens
}

type Lens struct {
	LensType int // 0: 普通镜头，1: 广角镜头，2: 长焦镜头
}

func (c Camera) Capture() { // 使用不同的镜头拍照
	m := make(map[int]string, 3)
	//m[0], m[1], m[2] = "normal lens", "wide-angle lens", "telephoto lens"
	m = map[int]string{0: "normal lens", 1: "wide-angle lenss", 2: "telephoto lens"}

	value := inj.Get(reflect.TypeOf(c.L))
	var index int
	if value.Kind() != reflect.Invalid {
		index = value.Interface().(Lens).LensType
	} else {
		index = 0
	}
	fmt.Println("capture with", m[index])
}

var inj *Injector // 持有一个全局的注入器，复用之前实现的注入器。
func main() {
	inj = New()

	c := Camera{"ZddCamera", Lens{0}} // 组装一个普通镜头的单反相机, 通常组装一个单反需要很长的时间，假设你有工匠情结，并且DIY成功的概率是100%，努力工作一个月就可以组装完一台。
	c.Capture()                       // 拍照
	// 如果想继续测试广角镜头和长焦镜头的拍照效果，那么是不是要这样？
	// d := Camera{"ZddCamera", Lens{1}} // 你又努力工作了一个月
	// e := Camera{"ZddCamera", Lens{2}} // 你又努力工作了一个月
	// 如果Boss告诉你，最近又进口了十几款镜头，你是不是感觉要死的心都有了呢。

	// 正常人只需要组装新的镜头
	wLens := Lens{0} // 努力工作十天
	tLens := Lens{1} // 努力工作十天

	inj.SetMap(wLens) // 复用之前的单反，只更换镜头
	c.Capture()       // 拍照

	inj.SetMap(tLens) // 复用之前的单反，只更换镜头
	c.Capture()       // 拍照

	// 是不是感觉棒棒嗒
	//: capture with normal lens
	//: capture with wide-angle lens
	//: capture with normal lens
}
