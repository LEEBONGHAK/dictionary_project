package main

import (
	"fmt"

	"github.com/LEEBONGHAK/dictionary_project/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary)

	definition, err := dictionary.Search("first")
	//definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	fmt.Println()

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}

	definition, err = dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dictionary)
}
