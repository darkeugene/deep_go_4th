package main

import (
	"fmt"
)

func main() {
	p1 := PersonBig{
		Name:    "John Doe",
		Age:     18,
		Married: true,
		Address: "Some Street",
		Balance: 100000000000000000000.23,

		ExtraField: "Married",
		someField:  23,
	}

	result := Serialize(p1)

	fmt.Println(result)

	var (
		p2, p3 PersonBig
		p4     Person
	)

	Deserialize(result, &p2)

	fmt.Printf("%#v\n", p2)

	Deserialize("name=John\nage=33\nmarried=false\nbalance=100000", &p3)

	fmt.Printf("%#v\n", p3)

	Deserialize(result, &p4)

	fmt.Printf("%#v\n", p4)
}
