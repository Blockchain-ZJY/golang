package main

import (
	"fmt"
	"reflect"
)

type MyStruct1 struct {
	F1 int
	F2 string
}

func isMyStruct(v interface{}) string {
	value := reflect.ValueOf(v)
	//如果是指针，则先取到值
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	// 上面三行基本是所有使用反射时必备的三句
	switch value.Kind() {
	case reflect.Int:
		return "是整数！！"
	case reflect.Struct:
		if value.Type() == reflect.TypeOf(MyStruct1{}) {
			return "是MyStruct1"
		}
	}

	return "啥也不是"
}

func main() {
	myVar1 := MyStruct1{F1: 1, F2: "test"}
	myVar2 := 123

	fmt.Printf("myVar1 is MyStruct1: %v\n", isMyStruct(myVar1))  // 输出: myVar1 is MyStruct1: 是MyStruct1
	fmt.Printf("myVar2 is MyStruct1: %v\n", isMyStruct(myVar2))  // 输出: 是整数！！myVar2 is MyStruct1: 是整数！！
	fmt.Printf("myVar2 is MyStruct1: %v\n", isMyStruct(&myVar1)) // 输出:  myVar1 is MyStruct1: 是MyStruct1
	fmt.Printf("myVar2 is MyStruct1: %v\n", isMyStruct(&myVar2)) // 输出: 是整数！！myVar2 is MyStruct1: 是整数！！
}
