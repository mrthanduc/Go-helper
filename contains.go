package helper

import (
	"reflect"
)

// ContainsInt tim gia tri value trong mang list
func ContainsInt(list []int, value int) bool {
	result := false
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			result = true
		}		
	}
	return result
}

// ContainsString tim chuoi value trong mang list
func ContainsString(list []string, value string) bool {
	result := false
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			result = true
		}		
	}
	return result
}

// Contains tim gia tri trong mang
func Contains(a interface{}, value interface{}) bool {
	result := false
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			result = true
		}
	}
	return result
}