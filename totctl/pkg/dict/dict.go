package dict

import (
	"fmt"

	"github.com/jx2lee/learngo-nomadcoder/totctl/pkg/dict/mydict"
)

func DictExam() {
	//dictionary := mydict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"
	//fmt.Println(dictionary)
	//definition, err := dictionary.Search("second")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(definition)
	//}

	//dictionary := mydict.Dictionary{}
	//word := "hello"
	//definition := "Greeting"
	//err := dictionary.Add(word, definition)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//hello, _ := dictionary.Search(word)
	//fmt.Println("found: ", word, "\ndefinition: ", hello)
	//err2 := dictionary.Add(word, definition)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}

	//dictionary := mydict.Dictionary{}
	//word := "hello"
	//definition := "Greeting"
	//dictionary.Add(word, definition)
	//err := dictionary.Update(word, "Second")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//test, _ := dictionary.Search(word)
	//fmt.Println(test)

	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	dictionary.Add(word, definition)
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	}
}
