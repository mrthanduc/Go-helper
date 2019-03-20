package helper

import (
	// "fmt"
	"reflect"
)

// IntersectionString la lay giao cua 2 chuoi
func IntersectionString(a, b []string) (inter []string){
	// result := []string{}
	//https://blog.golang.org/go-maps-in-action
	//The make function allocates and initializes a hash map data structure and returns a map value that points to it. 
	m := make(map[string]bool)
    for _, item := range a {
        m[item] = true
    }
    for _, item := range b {
        if _, ok := m[item]; ok {
            inter = append(inter, item)
        }
	}
	result := inter
	return result
}

// IntersectionInt la lay giao cua 2 mang so
func IntersectionInt(a, b []int) (inter []int){
	m := make(map[int]bool)

	// duyet qua cac gia tri trong a
    for _, item := range a {
		m[item] = true
    }
    for _, item := range b {
        if _, ok := m[item]; ok {
            inter = append(inter, item)
        }
	}
	result := inter
	return result
}

// IntersectionInt2 la lay giao cua 2 mang so
func IntersectionInt2(a, b []int) (inter []int){
	for i := 0; i < len(a); i++ {
		item := a[i];
		    if Contains(b, item) {
				inter = append(inter, item)		      
		    }
		}
	result := inter
	return result
}


// Intersection la lay giao cua 2 ...
func Intersection(a, b interface{}) interface{}{

	v := reflect.ValueOf(a)
	u := reflect.ValueOf(b)
	var inter []interface{}
	for i := 0; i < v.Len(); i++ {
		// var item []interface{}
		item := v.Index(i).Interface()
		in := reflect.ValueOf(item)
		    if Contains(u, in) {
				inter = append(inter, in)	      
		    }
		}
	
	result := inter
	return result
}

// Intersection1 la lay giao cua 2 ...
func Intersection1(a, b interface{}) interface{}{

	v := reflect.ValueOf(a)
	vType := v.Type()
	u := reflect.ValueOf(b)
	zType := reflect.SliceOf(vType.Elem())
	inter := reflect.MakeSlice(zType, 0, 0)
	for i := 0; i < v.Len(); i++ {
		// var item []interface{}
		item := v.Index(i).Interface()
		// in := reflect.ValueOf(item)
		    if Contains(u, item) {
				inter =  reflect.Append(inter, v.Index(i))	      
		    }
		}
	
	result := inter.Interface()
	return result
}

// Intersection2 la lay giao cua 2 ...
func Intersection2(a, b interface{}) interface{}{

	v := reflect.ValueOf(a)
	vType := v.Type()
	u := reflect.ValueOf(b)	
	// uType := u.Type()
	zType := reflect.SliceOf(vType.Elem())
	inter := reflect.MakeSlice(zType, 0, 0)

	m := map[interface{}]struct{}{}

	for i := 0; i < v.Len(); i++ {
		v1 := v.Index(i).Interface()
		m[v1] = struct{}{}
	}

	for i := 0; i < u.Len(); i++ {
		v1 := u.Index(i).Interface()
		_, ok := m[v1]
		if ok {
			inter = reflect.Append(inter, u.Index(i))
		}
	}

	result := inter.Interface()

	return result
}