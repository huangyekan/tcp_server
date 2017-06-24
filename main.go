package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	
}

func (t *Test) Test1(arg string)  string{
	return arg
}

func main() {
	t := new(Test)
	v := reflect.ValueOf(t).MethodByName("Test1").Call()
	fmt.Println(v[0])
}