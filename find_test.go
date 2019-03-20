package helper

import "testing"

func TestFindInt (t *testing.T){
	v := []int{1, 2, 3, 4, 5, 6}

	result := Find(v, func(elem int, in int) bool {
		return elem%2 == 0
	})
	expected := 2
	if result != expected {
		t.Error (" Bi sai roi")
	}
}

func TestFindInt2 (t *testing.T){
	v := []int{1, 2, 3, 4, 5, 6}

	result := Find(v, func(elem int) bool {
		return elem%2 == 0
	})
	expected := 2
	if result != expected {
		t.Error (" Bi sai roi")
	}
}


func TestFindString (t *testing.T){
	v := []string{"nghia","phu","bao","hoa"}
	result := FindString2(v, func(elem string) bool {
		return elem == "bao"
	})
	expected := "bao"
	if result != expected {
		t.Error (" Bi sai roi")
	}
}
