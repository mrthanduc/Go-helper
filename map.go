package helper

import "reflect"

type hamInt3 func(int) int
type ham3 func(int) int

//function types
type mapf func(interface{}) interface{}

//MapInt la ham anh xa 
func MapInt(list []int, ham hamInt3) []int {
	result := make([]int, 0)
	for _,data := range list {
		result = append(result,ham(data))		
	}
	return result
}

//Map ham anh xa
func Map(arr interface{}, predicate interface{}) interface{} {
	// var result []interface{}
	v := reflect.ValueOf(arr)
	funcValue := reflect.ValueOf(predicate)
	result := make([]interface{}, v.Len())

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		in := []reflect.Value{item}
		out := funcValue.Call(in)
		// result := out[i]
		result := out
		return result	
	}
	return result
}

//Map2 (slice, func)
func Map2 (in interface{}, fn mapf) interface{} {
	v := reflect.ValueOf(in)
	result := make([]interface{}, v.Len())

	for i := 0; i < v.Len(); i++ {
		result[i] = fn(v.Index(i).Interface())
	}

	return result
}