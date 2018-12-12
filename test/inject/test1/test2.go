package test1

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

func Host(name string, f func(a int, b string)) int {

	//s := 5.4
	//ss := reflect.ValueOf(s)
	//fmt.Println(ss.Interface().(float64))
	//fmt.Println(ss.String())
	//fmt.Println(ss)
	//fmt.Println("settability of v:", ss.CanSet())

	fmt.Println("Enter Host:", name)

	inj.Invoke(f) // 利用注入器中的环境调用f
	// 这种使用方法，看起来就像把自定义的方法f里的执行语句放在Host中执行一样自然
	// 语句从f里穿透到Host方法中，这就是注入一词的由来。

	fmt.Println("Exit Host:", name)
	return 1
}

func Dependency(a int, b string) {
	fmt.Println("Dependency: ", a, b)
}

var inj *Injector // 全局的注入器，保存注入环境

func main() {
	// 创建注入器
	inj = &Injector{make(map[reflect.Type]reflect.Value)}
	inj.SetMap(3030)
	inj.SetMap("zdd")

	d := Dependency
	Host("zddhub", d)

	inj.SetMap(8080)
	inj.SetMap("www.zddhub.com")
	Host("website", d)
}
