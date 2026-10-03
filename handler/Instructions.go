package handler

import "fmt"

type Instructions struct {
	path []string
}

func (i Instructions) callComands() {
	if len(i.path) == 0 {
		fmt.Println("You don't especified any comand")
		return
	}

	firstComand := i.path[0]

	comandFn, ok := allComands[firstComand]

	if !ok || comandFn == nil {
		fmt.Println("This comand don't exists")
		return
	}

	err := comandFn.Do()

	if err != nil {
		fmt.Println("Exists an error in the comand:")
		fmt.Println(err.Error())
	}
}
