package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func printAllValues(anything interface{}) {

	currentValue := reflect.ValueOf(anything)

	switch currentValue.Kind() {

	case reflect.Map:
		for _, mapKey := range currentValue.MapKeys() {
			printAllValues(currentValue.MapIndex(mapKey).Interface())
		}

	case reflect.Slice, reflect.Array:
		for index := 0; index < currentValue.Len(); index++ {
			printAllValues(currentValue.Index(index).Interface())
		}

	default:
		fmt.Printf(
			"Found value of type %T with data %v\n",
			currentValue.Interface(),
			currentValue.Interface(),
		)
	}
}

func main() {

	responseFromAPI := `
	{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			},
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			}
		],
		"sponsers" : {
			"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text" : ["one","two"],
		"int_text" : [1,3,4]
	}`

	var finalResult map[string]interface{}

	err := json.Unmarshal([]byte(responseFromAPI), &finalResult)
	if err != nil {
		panic(err)
	}

	printAllValues(finalResult)
}
