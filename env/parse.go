package env

import (
	"fmt"
	"reflect"
)

func Parse(s interface{}) (err error){
	if s == nil {
		return fmt.Errorf("invalid argument value (nil)")
	}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s).Kind()
	if t != reflect.Ptr {
		return fmt.Errorf("argument must be pointer to structure")
	}
	if v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("argument must be pointer to structure")
	}

	return
}