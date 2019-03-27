package helper

import "reflect"

type hamInt3 func(int) int

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

//Map3 ham anh xa
func Map3(arr interface{}, predicate interface{}) interface{} {
	// var result []interface{}
	v := reflect.ValueOf(arr)
	funcValue := reflect.ValueOf(predicate)
	result := make([]interface{}, v.Len())

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		in := []reflect.Value{item}
		out := funcValue.Call(in)
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

//Map bai sua cua thay
func Map(list interface{}, iterateeFunc interface{}) interface{} {
	listValue := reflect.ValueOf(list)
	listType := listValue.Type()
	listKind := listType.Kind()  // check kiểu dữ liệu của list

	iterateeFuncValue := reflect.ValueOf(iterateeFunc)
	typeOfResult := reflect.SliceOf(iterateeFuncValue.Type().Out(0))
	result := reflect.MakeSlice(typeOfResult, 0, 0)


	// Xác định xem list có phải là Slice hay Array không 
	if listKind == reflect.Slice || listKind == reflect.Array {
		for i := 0; i < listValue.Len(); i++ {
			elem := listValue.Index(i)
			in := []reflect.Value{elem}
			out := iterateeFuncValue.Call(in)[0]
			result = reflect.Append(result, out)
		}

	}
	return result.Interface()
}