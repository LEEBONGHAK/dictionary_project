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
}
