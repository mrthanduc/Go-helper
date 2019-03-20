package helper

import (
	"reflect"
	"testing"
)

func TestIntersectionString (t *testing.T){
	slice1 := []string{"foo", "bar", "hello", "bao"}
	slice2 := []string{"foo", "bar", "nghia", "bao"}
	
	result := Intersection2(slice1,slice2)
	expected := []string{"foo", "bar","bao"}

	if !reflect.DeepEqual(result, expected) {
		t.Error (" Bi sai roi")
	}
}

func TestIntersectionInt (t *testing.T){
	slice1 := []int{1,2,4,8,5}
	slice2 := []int{1,2,3}
	
	result := Intersection1(slice1,slice2)
	expected := []int{1,2}

	if !reflect.DeepEqual(result, expected) {
		t.Error (" Bi sai roi")
	}
}