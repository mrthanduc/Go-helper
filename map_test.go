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

func TestMapInt3(t *testing.T) {
	v := []int{1, 2, 3}
	expected := []int{2, 4, 6}
	actual := Map(v, func(elem int) int {
		return elem * 2
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Value should be expected!")
	}
}

func TestMapString(t *testing.T) {
	v := []string{"a", "b", "c"}
	expected := []string{"aa", "bb", "cc"}
	actual := Map(v, func(elem string) string {
		return elem + elem
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Value should be expected!")
	}
}

func TestMapString_2(t *testing.T) {
	v := []string{"a", "bb", "ccc"}
	expected := []int{1, 2, 3}
	actual := Map(v, func(elem string) int {
		return len(elem)
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Value should be expected!")
	}
}