package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myValue struct {
	IntValue     int       `json:"intValue"`
	BoolValue    bool      `json:"boolValue"`
	ObjectValue  *myObject `json:"objectValue"`
	NullIntValue *int      `json:"nullIntValue"`
}

type myObject struct {
}

func main() {
	data := map[string]interface{}{
		"intValue":     123,
		"dateValue":    time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"nullIntValue": nil,
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	newData := &myValue{
		IntValue:     1234,
		BoolValue:    true,
		ObjectValue:  &myObject{},
		NullIntValue: nil,
	}

	jsonData, err := json.Marshal(newData)
	if err != nil {
		fmt.Printf("Could not marshal json: %s\n", err)
		return
	}
	fmt.Printf("json data: %s\n", string(jsonData))
	fmt.Printf("map data: %s\n", data)
	fmt.Printf("map data: %#v\n", newData)

	// Unmarshaller
	var jData *myValue
	err = json.Unmarshal([]byte(jsonData), &jData)
	if err != nil {
		fmt.Printf("Could not unmarshall json: %s\n", err)
		return
	}
	fmt.Printf("json struct: %#v\n", jData)
	fmt.Printf("dateValue: %#v\n", jData.NullIntValue)
	fmt.Printf("objectValue: %#v\n", jData.ObjectValue)

}
