// All material is licensed under the Apache License Version 2.0, January 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to inspect a structs fields and display the field
// name, type and value.
package main

import (
	"fmt"
	"reflect"
)

// user represents a basic user in the system.
type user struct {
	name     string
	age      int
	building float32
	secure   bool
	roles    []string
}

func main() {

	// Create a value of the conrete type user.
	u := user{
		name:     "Cindy",
		age:      27,
		building: 321.45,
		secure:   true,
		roles:    []string{"admin", "developer"},
	}

	// Display the value we are passing.
	display(&u)
}

func display(v interface{}){
	userStrcut := reflect.ValueOf(v)
	if userStrcut.Kind() != reflect.Ptr {
		return
	}
	userStrcutValue := userStrcut.Elem()

	switch userStrcutValue.Kind(){
		case reflect.Struct:
		displayStruct(userStrcutValue)
	}
}

func displayStruct(v reflect.Value){
	for i:=0;i<v.NumField();i++{
		fld := v.Type().Field(i)
		fmt.Printf("Name: %s\t Kind: %s",fld.Name,fld.Type.Kind())
		fmt.Printf("\tValue: ")

	}
}

func displayStrcutValue(v reflect.Value){
	switch v.Kind(){
	 case reflect.String:
	 	fmt.Printf("%s",v.String())
	case reflect.Int:
		fmt.Printf("%v",v.Int())
	case reflect.Float32:
		fmt.Printf("%v",v.Float())
	case reflect.Bool:
		fmt.Printf("%v",v.Bool())
	case reflect.Slice:
		for i:=0;i<v.Len();i++{
			displayStrcutValue(v.Index(i))
			fmt.Printf(" ")
		}
	}
}