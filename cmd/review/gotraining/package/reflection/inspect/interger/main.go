package main

import (
	"reflect"
	"strconv"
)

func main(){
	var number int

}

func decodeInt(v interface{},number *int)error{
	rv := reflect.ValueOf(v)

	val := reflect.ValueOf(number).Elem()

	switch getKind(rv){
	case reflect.Int:
		val.SetInt(rv.Int())
		return nil
	case reflect.Uint:
		val.SetUint(rv.Uint())
		return nil
	case reflect.Float32:
		val.SetFloat(rv.Float())
		return nil
	case reflect.String:
		vv,err := strconv.ParseInt(rv.String(),0,val.Type().Bits())
		if err != nil{
			return err
		}
		val.SetInt(vv)
		return nil
	}
}

func getKind(val reflect.Value) reflect.Kind{
	// Capture the value's Kind.
	kind := val.Kind()

	// Check each condition until a case is true.
	switch {

	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int

	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint

	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32

	default:
		return kind
	}
}