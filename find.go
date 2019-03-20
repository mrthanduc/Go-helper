package helper

import "reflect"

type hamInt func(int, int) bool
type hamInt2 func(int) bool
type hamString func(string, int) bool
type hamString2 func(string) bool
type ham func(interface{}, int) bool

// FindInt tim so int thoa man ham
func FindInt(list []int, ham1 hamInt ) int {
	result := 0
	for index, result := range list {
		if ham1(result, index) {
			return result			
		}
	}
	return result	
}
// FindInt2 tim so int thoa man ham
func FindInt2(list []int, ham1 hamInt2 ) int {
	result := 0
	for _,result := range list {
		if ham1(result) {
			return result			
		}
	}
	return result	
}

// FindString tim so int thoa man ham
func FindString(list []string, ham1 hamString ) string {
	result := ""
	for index, result := range list {
		if ham1(result, index) {
			return result			
		}
	}
	return result
	
}

// FindString2 tim so int thoa man ham
func FindString2(list []string, ham1 hamString2 ) string {
	result := ""
	for _, result := range list {
		if ham1(result) {
			return result			
		}
	}
	return result
	
}

// Find iterates over the slice and apply ConditionIterator to every item. Returns first item that meet ConditionIterator or nil otherwise.
func Find(arr interface{}, predicate interface{}) interface{} {

	iteratorValue := reflect.ValueOf(arr)
	funcValue := reflect.ValueOf(predicate)

	for i := 0; i < iteratorValue.Len(); i++ {
		item := iteratorValue.Index(i)
		in := []reflect.Value{item}
		out := funcValue.Call(in)
		result := out[0].Bool()

		if result == true {
			return item.Interface()
		}
	}
	return nil
}

