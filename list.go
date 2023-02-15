package main

import (
	"container/list"
	"fmt"
)

func mai() {
	var intList list.List
	intList.PushBack(9)
	intList.PushBack(27)
	intList.PushBack(39)
	intList.PushBack(44)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
