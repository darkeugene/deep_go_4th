package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// **📗 В домашнем задании нужно реализовать сериализацию объекта структуры данных в** `.properties` **формат с использованием тегов структур и рефлексии.**
//При сериализации данных в некоторых случаях могут возникнуть проблемы с пустыми значениями. Например, если у вас есть структура, в которой некоторые поля могут быть не заполнены, и вы сериализуете, то в результате получится объект с пустыми полями. Если это не является ожидаемым поведением, то в нашей реализации можно будет использовать тег `omitempty`, чтобы пропустить пустые поля при сериализации.
//Стуктура для сериализации в `.properties` формат (*поподробнее с* `.properties` *форматом можно ознакомиться [здесь](https://ru.wikipedia.org/wiki/Properties):
// 📌 Для выполнения домашнего задания подготовлен шаблон кода и основные тесты, которую помогут проверить корректность реализации конвертации. Шаблона доступен [по ссылке.](https://github.com/Balun-courses/deep_go/blob/master/homework/generics_and_reflection/homework_test.go)
//Тесты в домашнем задании не всегда покрывают все случаи, а только основные. Будет круто, если вы добавите еще свои тест-кейсы, в случае необходимости!
//**Задание со звездочкой**
//__Выполнять необязательно__, но если вы хотите, можете попробовать реализовать обощенную функцию сериализации, которая сможет работать не только со структорой `Person`.

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

type PersonBig struct {
	Name    string  `properties:"name"`
	Address string  `properties:"address,omitempty"`
	Age     int     `properties:"age"`
	Married bool    `properties:"married"`
	Balance float64 `properties:"balance"`

	ExtraField string `properties:"-"`
	someField  int    `properties:"some_field"`
}

func Serialize(el any) string {
	t := reflect.TypeOf(el)
	v := reflect.ValueOf(el)
	results := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		value, present := field.Tag.Lookup("properties")
		if !present || value == "-" {
			continue
		}

		fieldValue := v.Field(i)
		var stringValue string

		switch fieldValue.Interface().(type) {
		case int, int8, int16, int32, int64:
			stringValue = strconv.FormatInt(fieldValue.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			stringValue = strconv.FormatUint(fieldValue.Uint(), 10)
		case float32, float64:
			stringValue = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
		case bool:
			stringValue = strconv.FormatBool(fieldValue.Bool())
		default:
			stringValue = fieldValue.String()
		}

		tags := strings.Split(value, ",")

		if len(tags) > 1 && tags[1] == "omitempty" && stringValue == "" {
			continue
		}

		results = append(results, fmt.Sprintf("%s=%s", tags[0], stringValue))
	}

	return strings.Join(results, "\n")
}

func Deserialize[T any](data string, el *T) {
	dataMap := stringIntoMap(data)

	t := reflect.TypeOf(*el)
	v := reflect.ValueOf(el).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		value, present := field.Tag.Lookup("properties")
		if !present || value == "-" {
			continue
		}

		tags := strings.Split(value, ",")

		dataValue, ok := dataMap[tags[0]]
		if !ok {
			if len(tags) > 1 && tags[1] == "omitempty" {
				continue
			}

			panic(fmt.Sprintf("missing required field %s", tags[0]))
		}

		fieldValue := v.Field(i)

		switch fieldValue.Interface().(type) {
		case int, int8, int16, int32, int64:
			intValue, err := strconv.ParseInt(dataValue, 10, 64)
			if err != nil {
				panic(err)
			}

			fieldValue.SetInt(intValue)
		case uint, uint8, uint16, uint32, uint64:
			uintValue, err := strconv.ParseUint(dataValue, 10, 64)
			if err != nil {
				panic(err)
			}

			fieldValue.SetUint(uintValue)
		case float32, float64:
			floatValue, err := strconv.ParseFloat(dataValue, 64)
			if err != nil {
				panic(err)
			}

			fieldValue.SetFloat(floatValue)
		case bool:
			boolValue, err := strconv.ParseBool(dataValue)
			if err != nil {
				panic(err)
			}

			fieldValue.SetBool(boolValue)
		default:
			fieldValue.SetString(dataValue)
		}
	}
}

func stringIntoMap(data string) map[string]string {
	rows := strings.Split(data, "\n")
	result := make(map[string]string, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		values := strings.Split(row, "=")

		if len(values) == 2 {
			result[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
		} else {
			result[strings.TrimSpace(values[0])] = ""
		}
	}

	return result
}
