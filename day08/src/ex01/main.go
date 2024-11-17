package main

import (
	"fmt"
	"log"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	v := reflect.ValueOf(plant)
	t := reflect.TypeOf(plant)

	if v.Kind() != reflect.Struct {
		log.Println("provided value is not a struct")
		return
	}

	var output string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag
		fieldName := field.Name
		tagString := ""
		if tag != "" {
			tagString = tag.Get("color_scheme")
			if tagString == "" {
				tagString = tag.Get("unit")
			}
			if tagString != "" {
				fieldName += fmt.Sprintf("(unit=%s)", tagString)
			}
		}

		if i > 0 {
			output += ", "
		}
		output += fmt.Sprintf("%s:%v", fieldName, value)
	}

	fmt.Println(output)
}

func main() {
	plant := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	}
	describePlant(plant)

	plant2 := UnknownPlant{
		FlowerType: "aster",
		LeafType:   "heart-leaf",
		Color:      134,
	}
	describePlant(plant2)
}
