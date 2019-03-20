package helper

import (
	"testing"
)

func TestContaintInt(t *testing.T){
	v := []int{1,2,3}
	result := Contains(v, 3)
	expected := true
	if result != expected {
		t.Error (" Bi sai roi")
	}
}

func TestContaintInt1(t *testing.T){
	v := []int{1,2,3}
	result := Contains(v, 4)
	expected := false
	if result != expected {
		t.Error (" Bi sai roi")
	}
}


func TestContaintString(t *testing.T){
	v := []string{"Nghia","Phu","Hieu"}
	result := Contains(v, "Hieu")
	expected := true
	if result != expected {
		t.Error (" Bi sai roi")
	}
}

func TestContaintString2(t *testing.T){
	v := []string{"Nghia","Phu","Hieu"}
	result := Contains(v, "Hai")
	expected := false
	if result != expected {
		t.Error (" Bi sai roi")
	}
}

func TestContaintFloat642(t *testing.T){
	v := []float64{1.23,1.54,5.67}
	result := Contains(v, 5.67)
	expected := true
	if result != expected {
		t.Error (" Bi sai roi")
	}
}