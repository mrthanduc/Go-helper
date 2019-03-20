package helper

import "reflect"

type hamreduce func(int,int) int

// func(value, memo) interface
type reducef func(interface{}, interface{}) interface{}

//ReduceInt gom gia tri cua list them hamreduce ra 1
func ReduceInt(list []int, memo int, ham hamreduce) int{
	// result := 0
	for _, result:= range list{
		memo = ham(result, memo)
	}
	return memo
}

//Reduce (slice, starting value, func)
func Reduce(in interface{}, memo interface{}, fn reducef) interface{} {
	val := reflect.ValueOf(in)

	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}

	return memo
}