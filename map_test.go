package helper

import (
	"reflect"
	"fmt"
	"testing"
)

func TestMapInt(t *testing.T){
	v:= []int{1,2,3}
	
	result := MapInt(v,func(elem int) int {return elem*4})
	fmt.Println(result)
	// expected := []int{3,6,9}
	expected := []int{4,8,12}
	if !reflect.DeepEqual(result, expected) {
		t.Error (" Bi sai roi")
	}
}

func TestMapInt2(t *testing.T){
	v:= []int{1,2,3}
	
	result := Map2(v,func(val interface{}) interface{} {
		return val.(int) * 4
	})
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
	
	expected := []int{4,8,12}
	if !reflect.DeepEqual(result, expected) {
		t.Error (" Bi sai roi")
	}
}
