package helper

import (
	"fmt"
	// "reflect"
	"testing"
)

func TestReduceInt(t *testing.T){
	v:= []int{1,2,3}
	result := ReduceInt(v, 0, func(val int, memo int) int {
		return memo + val
	})
	fmt.Println(result)
	expected := 7
	if result != expected {
		t.Error (" Bi sai roi")
	}
}
func TestReduce(t *testing.T){
	v:= []int{1,2,3}
	result := Reduce(v, 0, func(val interface{}, memo interface{}) interface{} {
		return memo.(int) + val.(int)*2
	})
	fmt.Println(result)
	expected := 12
	if result != expected {
		t.Error (" Bi sai roi")
	}
}

// func TestReduce1(t *testing.T){
// 	v:= []string{"a","b","c"}
// 	result := Reduce(v, "", func(val interface{}, memo interface{}) interface{} {
// 		return memo + val
// 	})
// 	fmt.Println(result)
// 	expected := 12
// 	if result != expected {
// 		t.Error (" Bi sai roi")
// 	}
// }

// func TestReduce2(t *testing.T){
// 	v:= [][]int{{0, 1}, {2, 3}, {4, 5}}
// 	result := Reduce(v, 0, func(val interface{}, memo interface{}) interface{} {
// 		return append(val.([]int)) // dang bi sai o cai ham nay
// 	})
// 	fmt.Println(result)
// 	expected := []int{0,1,2,3,4,5}
// 	if !reflect.DeepEqual(result, expected) {
// 		t.Error (" Bi sai roi")
// 	}
// }