package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializationPerson(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestSerializationBigPerson(t *testing.T) {
	tests := map[string]struct {
		person PersonBig
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false\nbalance=0",
		},
		"test case with fields": {
			person: PersonBig{
				Name:    "John Doe",
				Age:     18,
				Married: true,
				Address: "Some Street",
				Balance: 100000000000000000000.23,

				ExtraField: "Married",
				someField:  23,
			},
			result: "name=John Doe\naddress=Some Street\nage=18\nmarried=true\nbalance=100000000000000000000",
		},
		"test case with omitempty field": {
			person: PersonBig{
				Name:    "John Doe",
				Age:     18,
				Married: true,

				Balance: 100000000000000000000.23,

				ExtraField: "Married",
				someField:  23,
			},
			result: "name=John Doe\nage=18\nmarried=true\nbalance=100000000000000000000",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestDeserializationPersonBig(t *testing.T) {
	tests := map[string]struct {
		person PersonBig
		data   string
	}{
		"test case with empty fields": {
			data: "name=\nage=0\nmarried=false\nbalance=0",
		},
		"test case with fields": {
			person: PersonBig{
				Name:    "John Doe",
				Age:     18,
				Married: true,
				Address: "Some Street",
				Balance: 100000000000000000000,
			},
			data: "name=John Doe\naddress=Some Street\nage=18\nmarried=true\nbalance=100000000000000000000",
		},
		"test case with omitempty field": {
			person: PersonBig{
				Name:    "John Doe",
				Age:     18,
				Married: true,

				Balance: 100000000000000000000,
			},
			data: "name=John Doe\nage=18\nmarried=true\nbalance=100000000000000000000",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var person PersonBig
			Deserialize(test.data, &person)
			assert.Equal(t, test.person, person)
		})
	}
}

func TestDeserializationPerson(t *testing.T) {
	tests := map[string]struct {
		person Person
		data   string
	}{
		"test case with empty fields": {
			data: "name=\nage=0\nmarried=false\nbalance=0",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     18,
				Married: true,
				Address: "Some Street",
			},
			data: "name=John Doe\naddress=Some Street\nage=18\nmarried=true\nbalance=100000000000000000000",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     18,
				Married: true,
			},
			data: "name=John Doe\nage=18\nmarried=true\nbalance=100000000000000000000",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var person Person
			Deserialize(test.data, &person)
			assert.Equal(t, test.person, person)
		})
	}
}
