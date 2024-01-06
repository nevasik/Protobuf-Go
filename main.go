package main

import (
	"ProtobufApp/example/protos/poplaukhin"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	// создание нового объекта Person
	person := &poplaukhin.Person{
		Name:   "Dima",
		Id:     1,
		Phones: []string{"8-927-947-83-20", "8-977-168-95-75"},
	}

	// преобразуем его в binary format
	data, err := proto.Marshal(person)

	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	// Преобразуем binary format обратно в объект Person
	newPerson := &poplaukhin.Person{}
	err = proto.Unmarshal(data, newPerson)

	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
		return
	}

	fmt.Println(newPerson)
}
