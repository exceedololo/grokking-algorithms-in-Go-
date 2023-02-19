package main

import (
	"container/list"
	"fmt"
)

var graph = map[string]map[string]bool{
	"you": {
		"alice":  true,
		"bob":    true,
		"claire": true,
	},
	"bob": {
		"anuj":  true,
		"peggy": true,
	},
	"alice": {
		"peggy": true,
	},
	"claire": {
		"thom":  true,
		"jonny": true,
	},
	"anuj":  {},
	"peggy": {},
	"thom":  {},
	"jonny": {},
}

func personIsSeller(name string) bool {
	return name[len(name)-1] == 'm'
}

func search(name string) bool {
	searchQueue := list.New()
	for neighbor := range graph[name] {
		searchQueue.PushBack(neighbor)
	}
	searched := make(map[string]bool)
	searched[name] = true
	for searchQueue.Len() > 0 {
		person := searchQueue.Remove(searchQueue.Front()).(string)  //Таким образом, person := searchQueue.Remove(searchQueue.Front()).(string) удаляет первый элемент из списка searchQueue, 
		преобразует его тип к string и присваивает его переменной person. В итоге мы получаем строку, содержащую следующий узел для поиска.
		if !searched[person] {
			if personIsSeller(person) {
				fmt.Println(person + " is a mango seller!")
				return true
			} else {
				for neighbor := range graph[person] {
					searchQueue.PushBack(neighbor)
				}
				searched[person] = true
			}
		}
	}
	return false
}

func main() {
	search("you")
}
