package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
 }

func main() {
	var myVar myStruct
	myVar.FirstName = "Wisdom"

	myVar2 := myStruct {
		FirstName: "Elue",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}